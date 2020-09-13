package k8s

import (
	"context"

	corev1 "k8s.io/api/core/v1"
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

func (c *Client) VolumesWithClaimsAndPods(ctx context.Context) ([]VolumeClaimPods, error) {
	pvl, err := c.Interface.CoreV1().PersistentVolumes().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	pvcl, err := c.Interface.CoreV1().PersistentVolumeClaims("").List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	pl, err := c.Interface.CoreV1().Pods("").List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	volumes := make([]VolumeClaimPods, len(pvl.Items))

	for i, pv := range pvl.Items {
		pv.ManagedFields = nil
		volumes[i].PersistentVolume = pv

		for _, pvc := range pvcl.Items {
			if pv.Name == pvc.Spec.VolumeName {
				pvc.ManagedFields = nil
				volumes[i].PersistentVolumeClaim = pvc

				volumes[i].Pods = make([]corev1.Pod, 0)

				for _, pod := range pl.Items {
					for _, v := range pod.Spec.Volumes {
						if v.PersistentVolumeClaim != nil && pvc.Name == v.PersistentVolumeClaim.ClaimName {
							pod.ManagedFields = nil
							volumes[i].Pods = append(volumes[i].Pods, pod)

							break
						}
					}
				}

				break
			}
		}
	}

	return volumes, nil
}

type VolumeClaimPods struct {
	PersistentVolume      corev1.PersistentVolume      `json:"volume"`
	PersistentVolumeClaim corev1.PersistentVolumeClaim `json:"claim"`
	Pods                  []corev1.Pod                 `json:"pods"`
}
