package k8s

import (
	"context"

	v1 "k8s.io/api/storage/v1"
)

// Interface contains all public methods of k8s.Client implementation.
type Interface interface {
	StorageClasses(ctx context.Context) ([]v1.StorageClass, error)
	PersistentVolumesWithClaimsAndPods(ctx context.Context) ([]VolumeClaimPods, error)
	SetPersistentVolumeReclaimPolicy(ctx context.Context, persistentVolumeName, policy string) error
	SetDefaultStorageClass(ctx context.Context, storageClassName string) error
}
