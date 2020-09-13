package k8s

import (
	"context"

	v1 "k8s.io/api/storage/v1"
)

type Interface interface {
	StorageClasses(ctx context.Context) ([]v1.StorageClass, error)
	VolumesWithClaimsAndPods(ctx context.Context) ([]VolumeClaimPods, error)
	SetPersistentVolumeReclaimPolicy(ctx context.Context, persistentVolumeName, policy string) error
}
