package argocd

import (
	"context"
	"github.com/argoproj/argo-cd/v2/pkg/apiclient/application"
	"github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CreateApplication on the Client struct in Go. It takes no parameters and returns a pointer to v1alpha1.Application and an error.
func (c *Client) CreateApplication() (*v1alpha1.Application, error) {
	app := &v1alpha1.Application{
		TypeMeta: v1.TypeMeta{
			Kind:       "Application",
			APIVersion: "argoproj.io/v1alpha1",
		},
		ObjectMeta: v1.ObjectMeta{
			Name:              "",
			Namespace:         "",
			CreationTimestamp: v1.Time{Time: v1.Now().Time},
		},
		Spec: v1alpha1.ApplicationSpec{
			Source: nil,
			Destination: v1alpha1.ApplicationDestination{
				Name:      "",
				Namespace: "",
			},
			Project:    "",
			SyncPolicy: nil,
			Sources: []v1alpha1.ApplicationSource{
				{
					RepoURL:        "", // required
					Path:           "", // required
					TargetRevision: "", // required
					Helm: &v1alpha1.ApplicationSourceHelm{
						ValueFiles: nil, // required
					},
				},
			},
		},
	}

	return c.applicationServiceClient.Create(context.Background(), &application.ApplicationCreateRequest{
		Application: app,
	})
}

// DeleteApplication on the Client struct in Go. It takes in a string parameter app and returns an error.
func (c *Client) DeleteApplication(app string) error {
	_, err := c.applicationServiceClient.Delete(context.Background(), &application.ApplicationDeleteRequest{
		Name: &app,
	})
	if err != nil {
		return err
	}
	return nil
}

// GetApplication on the Client struct in Go. It takes in a string parameter app and returns a pointer to v1alpha1.Application and an error.
func (c *Client) GetApplication(app string) (*v1alpha1.Application, error) {
	return c.applicationServiceClient.Get(context.Background(), &application.ApplicationQuery{
		Name: &app,
	})
}

// GetCustomActionList on the Client struct in Go. It takes in a string parameter app, a string parameter kind, a string parameter resourceName, and a string parameter namespace.
func (c *Client) GetCustomActionList(app, kind, resourceName, namespace string) ([]string, error) {
	version, group := getVariables(kind)
	resp, err := c.applicationServiceClient.ListResourceActions(context.Background(), &application.ApplicationResourceRequest{
		Name:         &app,
		Namespace:    &namespace,
		ResourceName: &resourceName,
		Version:      &version,
		Group:        &group,
		Kind:         &kind,
	})
	if err != nil {
		return nil, err
	}

	actions := resp.GetActions()
	var actionList []string

	for _, v := range actions {
		actionList = append(actionList, v.Name)
	}
	return actionList, nil
}

// ListApplications on the Client struct in Go. It takes no parameters and returns a slice of v1alpha1.Application structs and an error.
func (c *Client) ListApplications() ([]v1alpha1.Application, error) {
	apps, err := c.applicationServiceClient.List(context.Background(), &application.ApplicationQuery{})
	if err != nil {
		return nil, err
	}

	return apps.Items, nil
}

// RunCustomAction on the Client struct in Go. It takes in four parameters: app, kind, resourceName, and namespace, all of type string.
func (c *Client) RunCustomAction(app, kind, resourceName, namespace, actionName string) error {
	version, group := getVariables(kind)
	_, err := c.applicationServiceClient.RunResourceAction(context.Background(), &application.ResourceActionRunRequest{
		Name:         &app,
		Namespace:    &namespace,
		ResourceName: &resourceName,
		Version:      &version,
		Kind:         &kind,
		Group:        &group,
		Action:       &actionName,
	})
	if err != nil {
		return err
	}
	return nil
}

// SyncApplication on the Client struct in Go. It takes in a string parameter app and a boolean parameter prune.
func (c *Client) SyncApplication(app string, prune bool) error {
	_, err := c.applicationServiceClient.Sync(context.Background(), &application.ApplicationSyncRequest{
		Name:  &app,
		Prune: &prune,
	})
	if err != nil {
		return err
	}
	return nil
}

func getVariables(kind string) (string, string) {
	var version, group string
	switch kind {
	case "Deployment":
		version = "v1"
		group = "apps"
	case "StatefulSet":
		version = "v1"
		group = "apps"
	case "DaemonSet":
		version = "v1"
		group = "apps"
	case "Job":
		version = "v1"
		group = "batch"
	case "CronJob":
		version = "v1"
		group = "batch"
	case "Service":
		version = "v1"
	case "Ingress":
		version = "v1"
	case "Rollouts":
		version = "v1alpha1"
		group = "argoproj.io"
	}
	return version, group
}
