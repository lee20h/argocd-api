package argocd

import (
	"context"

	"github.com/argoproj/argo-cd/v2/pkg/apiclient/cluster"

	"github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
)

// GetClusters on the Client struct in Go. It takes no parameters and returns a slice of v1alpha1.Cluster and an error.
func (c *Client) GetClusters() ([]v1alpha1.Cluster, error) {
	cl, err := c.clusterClient.List(context.Background(), &cluster.ClusterQuery{})
	if err != nil {
		return nil, err
	}

	return cl.Items, nil
}
