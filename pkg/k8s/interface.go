package k8s

import (
	"context"

	v1 "k8s.io/api/storage/v1"
)

type Interface interface {
	StorageClasses(context.Context) ([]v1.StorageClass, error)
	VolumesWithClaimsAndPods(context.Context) ([]VolumeClaimPods, error)
}
