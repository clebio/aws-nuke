package resources

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/ec2"
)

type EC2VpnConnection struct {
	svc   *ec2.EC2
	id    string
	state string
}

func (n *EC2Nuke) ListVpnConnections() ([]Resource, error) {
	params := &ec2.DescribeVpnConnectionsInput{}
	resp, err := n.Service.DescribeVpnConnections(params)
	if err != nil {
		return nil, err
	}

	resources := make([]Resource, 0)
	for _, out := range resp.VpnConnections {
		resources = append(resources, &EC2VpnConnection{
			svc:   n.Service,
			id:    *out.VpnConnectionId,
			state: *out.State,
		})
	}

	return resources, nil
}

func (v *EC2VpnConnection) Filter() error {
	if v.state == "deleted" {
		return fmt.Errorf("already deleted")
	}
	return nil
}

func (v *EC2VpnConnection) Remove() error {
	params := &ec2.DeleteVpnConnectionInput{
		VpnConnectionId: &v.id,
	}

	_, err := v.svc.DeleteVpnConnection(params)
	if err != nil {
		return err
	}

	return nil
}

func (v *EC2VpnConnection) String() string {
	return v.id
}
