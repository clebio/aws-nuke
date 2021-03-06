package resources

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/rds"
)

type RDSCluster struct {
	svc *rds.RDS
	id  string
}

func (n *RDSNuke) ListClusers() ([]Resource, error) {
	params := &rds.DescribeDBClustersInput{}
	resp, err := n.Service.DescribeDBClusters(params)
	if err != nil {
		return nil, err
	}

	resources := make([]Resource, 0)
	for _, instance := range resp.DBClusters {
		resources = append(resources, &RDSCluster{
			svc: n.Service,
			id:  *instance.DBClusterIdentifier,
		})
	}

	return resources, nil
}

func (i *RDSCluster) Remove() error {
	params := &rds.DeleteDBClusterInput{
		DBClusterIdentifier: &i.id,
		SkipFinalSnapshot:   aws.Bool(true),
	}

	_, err := i.svc.DeleteDBCluster(params)
	if err != nil {
		return err
	}

	return nil
}

func (i *RDSCluster) String() string {
	return i.id
}
