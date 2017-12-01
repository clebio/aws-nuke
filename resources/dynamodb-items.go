package resources

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type DynamoDBTableItem struct {
	svc   *dynamodb.DynamoDB
	id    map[string]*dynamodb.AttributeValue
	table Resource
}

func (n *DynamoDBNuke) ListItems() ([]Resource, error) {
	tables, err := n.ListTables()
	if err != nil {
		return nil, err
	}

	resources := make([]Resource, 0)
	for _, dynamoTable := range tables {
		params := &dynamodb.ScanInput{
			TableName: aws.String(dynamoTable.String()),
		}
		resp, err := n.Service.Scan(params)
		if err != nil {
			return nil, err
		}
		for _, itemMap := range resp.Items {
			resources = append(resources, &DynamoDBTableItem{
				svc:   n.Service,
				id:    itemMap,
				table: dynamoTable,
			})
		}
	}
	return resources, nil
}

func (i *DynamoDBTableItem) Remove() error {
	params := &dynamodb.DeleteItemInput{
		Key:       i.id,
		TableName: aws.String(i.table.String()),
	}
	_, err := i.svc.DeleteItem(params)
	if err != nil {
		return err
	}
	return nil
}

func (i *DynamoDBTableItem) String() string {
	table := i.table.String()
	keyField := ""
	for key, value := range i.id {
		value := strings.Replace(value.String(), "\n", " ", -1)
		value = strings.Replace(value, "  ", " ", -1)
		keyField = " primary key: " + key + " value: " + value
	}
	return "table: " + table + keyField
}
