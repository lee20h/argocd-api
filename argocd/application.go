package argocd

import (
	"context"
	"github.com/argoproj/argo-cd/v2/pkg/apiclient/application"
	"github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Client) CreateApplication(app *v1alpha1.Application) (*v1alpha1.Application, error) {
	application.ApplicationCreateRequest{
		Application: v1alpha1.Application{
			TypeMeta: v1.TypeMeta{
				Kind:       "",
				APIVersion: "",
			},
			ObjectMeta: v1.ObjectMeta{
				Name:                       "",
				GenerateName:               "",
				Namespace:                  "",
				UID:                        "",
				ResourceVersion:            "",
				Generation:                 0,
				CreationTimestamp:          v1.Time{},
				DeletionTimestamp:          nil,
				DeletionGracePeriodSeconds: nil,
				Labels:                     nil,
				Annotations:                nil,
				OwnerReferences:            nil,
				Finalizers:                 nil,
				ManagedFields:              nil,
			},
			Spec: v1alpha1.ApplicationSpec{
				Source:               nil,
				Destination:          v1alpha1.ApplicationDestination{},
				Project:              "",
				SyncPolicy:           nil,
				IgnoreDifferences:    nil,
				Info:                 nil,
				RevisionHistoryLimit: nil,
				Sources:              nil,
			},
			Status: v1alpha1.ApplicationStatus{
				Resources:            nil,
				Sync:                 v1alpha1.SyncStatus{},
				Health:               v1alpha1.HealthStatus{},
				History:              nil,
				Conditions:           nil,
				ReconciledAt:         nil,
				OperationState:       nil,
				ObservedAt:           nil,
				SourceType:           "",
				Summary:              v1alpha1.ApplicationSummary{},
				ResourceHealthSource: "",
				SourceTypes:          nil,
			},
			Operation: nil,
		},
		Upsert:   nil,
		Validate: nil,
	}

	return c.applicationServiceClient.Create(context.Background(), &application.ApplicationCreateRequest{
		Application: app,
	})
}

func (c *Client) DeleteApplication(app string) error {
	_, err := c.applicationServiceClient.Delete(context.Background(), &application.ApplicationDeleteRequest{
		Name: &app,
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) GetApplication(app string) (*v1alpha1.Application, error) {
	return c.applicationServiceClient.Get(context.Background(), &application.ApplicationQuery{
		Name: &app,
	})
}

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

func (c *Client) ListApplications() ([]v1alpha1.Application, error) {
	apps, err := c.applicationServiceClient.List(context.Background(), &application.ApplicationQuery{})
	if err != nil {
		return nil, err
	}

	return apps.Items, nil
}

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
