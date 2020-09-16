package k8s_test

import (
	"context"
	"testing"

	"github.com/acim/lazarette/pkg/k8s"
	v1 "k8s.io/api/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

func TestStorageClasses(t *testing.T) {
	c, err := k8s.NewClient(fake.NewSimpleClientset(
		&v1.StorageClass{
			ObjectMeta: metav1.ObjectMeta{
				Name: "hcloud-volumes",
			},
		},
		&v1.StorageClass{
			ObjectMeta: metav1.ObjectMeta{
				Name: "nfs",
			},
		},
	))
	if err != nil {
		t.Error(err)
	}

	scs, _ := c.StorageClasses(context.Background())

	l := len(scs)
	if len(scs) != 2 {
		t.Fatalf("expected 2 storage classes, got %d", l)
	}

	n0 := scs[0].ObjectMeta.Name
	if n0 != "hcloud-volumes" {
		t.Fatalf("expected name hcloud-volumes for storage classes 0, got %s", n0)
	}

	n1 := scs[1].ObjectMeta.Name
	if n1 != "nfs" {
		t.Fatalf("expected name nfs for storage classes 1, got %s", n1)
	}

	if scs[0].ManagedFields != nil || scs[1].ManagedFields != nil {
		t.Fatal("expected managed fields to be nil")
	}
}
