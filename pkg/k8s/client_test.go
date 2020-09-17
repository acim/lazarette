package k8s_test

import (
	"context"
	"testing"

	"github.com/acim/lazarette/pkg/k8s"
	v1 "k8s.io/api/core/v1"
	storagev1 "k8s.io/api/storage/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

func TestStorageClasses(t *testing.T) {
	c, err := k8s.NewClient(fake.NewSimpleClientset(
		&storagev1.StorageClass{
			ObjectMeta: metav1.ObjectMeta{
				Name: "foo",
			},
		},
		&storagev1.StorageClass{
			ObjectMeta: metav1.ObjectMeta{
				Name: "bar",
			},
		},
	))
	if err != nil {
		t.Error(err)
	}

	scs, _ := c.StorageClasses(context.Background())

	if l := len(scs); l != 2 {
		t.Fatalf("expected 2 storage classes, got %d", l)
	}

	if n0 := scs[0].ObjectMeta.Name; n0 != "bar" {
		t.Fatalf("expected name bar for storage classes 0, got %s", n0)
	}

	if n1 := scs[1].ObjectMeta.Name; n1 != "foo" {
		t.Fatalf("expected name foo for storage classes 1, got %s", n1)
	}

	if scs[0].ManagedFields != nil || scs[1].ManagedFields != nil {
		t.Fatal("expected managed fields to be nil")
	}
}

func TestPersistentVolumesWithClaimsAndPods(t *testing.T) { //nolint:funlen
	scName := "foo-sc"
	pvName := "foo-pv"
	pvcName := "foo-pvc"

	c, err := k8s.NewClient(fake.NewSimpleClientset(
		&storagev1.StorageClass{
			ObjectMeta: metav1.ObjectMeta{
				Name: scName,
			},
			ReclaimPolicy: k8s.ToPVReclaimPolicyPtr("Delete"),
		},
		&v1.PersistentVolume{
			ObjectMeta: metav1.ObjectMeta{
				Name: pvName,
			},
			Spec: v1.PersistentVolumeSpec{
				AccessModes: []v1.PersistentVolumeAccessMode{
					v1.PersistentVolumeAccessMode("ReadWriteMany"),
				},
				Capacity: v1.ResourceList{
					v1.ResourceName("storage"): resource.MustParse("2Gi"),
				},
				ClaimRef: &v1.ObjectReference{
					Kind: "PersistentVolumeClaim",
					Name: pvcName,
				},
				PersistentVolumeReclaimPolicy: v1.PersistentVolumeReclaimPolicy("Retain"),
				StorageClassName:              scName,
			},
		},
		&v1.PersistentVolumeClaim{
			ObjectMeta: metav1.ObjectMeta{
				Name: pvcName,
			},
			Spec: v1.PersistentVolumeClaimSpec{
				AccessModes: []v1.PersistentVolumeAccessMode{
					v1.PersistentVolumeAccessMode("ReadWriteMany"),
				},
				Resources: v1.ResourceRequirements{
					Requests: v1.ResourceList{
						v1.ResourceName("storage"): resource.MustParse("2Gi"),
					},
				},
				StorageClassName: &scName,
				VolumeName:       pvName,
			},
		},
		&v1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Name: "fooPod",
			},
			Spec: v1.PodSpec{
				Volumes: []v1.Volume{
					{
						Name: "fooPodVolume",
						VolumeSource: v1.VolumeSource{
							PersistentVolumeClaim: &v1.PersistentVolumeClaimVolumeSource{
								ClaimName: pvcName,
							},
						},
					},
				},
			},
		},
	))
	if err != nil {
		t.Error(err)
	}

	pvs, _ := c.PersistentVolumesWithClaimsAndPods(context.Background())
	t.Log(pvs)
}
