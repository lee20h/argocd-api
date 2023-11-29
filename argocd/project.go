package argocd

import (
	"context"

	"github.com/argoproj/argo-cd/v2/pkg/apiclient/project"
	"github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CreateProject on a struct called Client. It takes a name argument and returns a pointer to v1alpha1.AppProject and an error.
func (c *Client) CreateProject(name string) (*v1alpha1.AppProject, error) {
	return c.projectClient.Create(context.Background(), &project.ProjectCreateRequest{
		Project: &v1alpha1.AppProject{
			ObjectMeta: v1.ObjectMeta{
				Name: name,
			},
		},
	})
}

// DeleteProject on the Client struct in Go. It takes a name parameter and returns an error.
func (c *Client) DeleteProject(name string) error {
	_, err := c.projectClient.Delete(context.Background(), &project.ProjectQuery{
		Name: name,
	})

	return err
}

// GetProject on the Client struct in Go. It takes a name parameter and returns a pointer to v1alpha1.AppProject and an error.
func (c *Client) GetProject(name string) (*v1alpha1.AppProject, error) {
	return c.projectClient.Get(context.Background(), &project.ProjectQuery{
		Name: name,
	})
}
