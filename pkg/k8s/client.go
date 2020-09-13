package k8s

import (
	"context"

	v1 "k8s.io/api/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type Client struct {
	kubernetes.Interface
}

func NewClient(c *rest.Config) (*Client, error) {
	cs, err := kubernetes.NewForConfig(c)
	if err != nil {
		return nil, err
	}

	return &Client{cs}, nil
}

func (c *Client) StorageClasses(ctx context.Context) ([]v1.StorageClass, error) {
	scl, err := c.Interface.StorageV1().StorageClasses().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	for i := range scl.Items {
		scl.Items[i].ManagedFields = nil
	}

	return scl.Items, nil
}
