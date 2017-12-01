package resources

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type DynamoDBTable struct {
	svc *dynamodb.DynamoDB
	id  string
}

func (n *DynamoDBNuke) ListTables() ([]Resource, error) {
	params := &dynamodb.ListTablesInput{}
	resp, err := n.Service.ListTables(params)
	if err != nil {
		return nil, err
	}

	resources := make([]Resource, 0)
	for _, tableName := range resp.TableNames {
		resources = append(resources, &DynamoDBTable{
			svc: n.Service,
			id:  *tableName,
		})
	}

	return resources, nil
}

func (i *DynamoDBTable) Remove() error {
	// params := &dybamodb.Delete{
	// 	DBInstanceIdentifier: &i.id,
	// 	SkipFinalSnapshot:    aws.Bool(true),
	// }
	//
	// _, err := i.svc.DeleteDBInstance(params)
	// if err != nil {
	// 	return err
	// }

	return nil
}

func (i *DynamoDBTable) String() string {
	return i.id
}
