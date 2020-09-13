package k8s

import (
	"context"
	"encoding/json"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// Client contains wrapper methods to Kubernetes API.
type Client struct {
	kubernetes.Interface
}

// NewClient creates new Kubernetes client.
func NewClient(c *rest.Config) (*Client, error) {
	cs, err := kubernetes.NewForConfig(c)
	if err != nil {
		return nil, err
	}

	return &Client{cs}, nil
}

// StorageClasses returns all storage classes in the cluster.
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

// PersistentVolumesWithClaimsAndPods returns all persistent volumes in the cluster together with their belonging
// persistent volume claims and pods.
func (c *Client) PersistentVolumesWithClaimsAndPods(ctx context.Context) ([]VolumeClaimPods, error) {
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

// SetPersistentVolumeReclaimPolicy sets reclaim policy to the given persistent volume.
func (c *Client) SetPersistentVolumeReclaimPolicy(ctx context.Context, persistentVolumeName, policy string) error {
	p := []patchStringValue{
		{
			Op:    "replace",
			Path:  "/spec/persistentVolumeReclaimPolicy",
			Value: policy,
		},
	}

	payload, err := json.Marshal(p)
	if err != nil {
		return err
	}

	_, err = c.Interface.CoreV1().PersistentVolumes().Patch(
		ctx, persistentVolumeName, types.JSONPatchType, payload, metav1.PatchOptions{})

	return err
}

// SetDefaultStorageClass sets default storage class and unsets default flat to all other storage classes.
func (c *Client) SetDefaultStorageClass(ctx context.Context, storageClassName string) error {
	scl, err := c.Interface.StorageV1().StorageClasses().List(ctx, metav1.ListOptions{})
	if err != nil {
		return err
	}

	path := fmt.Sprintf("/metadata/annotations/%s", jsonPointerEscape("storageclass.kubernetes.io/is-default-class"))

	for _, item := range scl.Items {
		var p []patchStringValue

		switch item.GetName() {
		case storageClassName:
			p = []patchStringValue{
				{
					Op:    "add",
					Path:  path,
					Value: "true",
				},
			}
		default:
			p = []patchStringValue{
				{
					Op:   "remove",
					Path: path,
				},
			}
		}

		payload, err := json.Marshal(p)
		if err != nil {
			return fmt.Errorf("storage class %s: %w", item.GetName(), err)
		}

		_, err = c.Interface.StorageV1().StorageClasses().Patch(
			ctx, item.GetName(), types.JSONPatchType, payload, metav1.PatchOptions{})
		if err != nil {
			return fmt.Errorf("storage class %s: %w", item.GetName(), err)
		}
	}

	return nil
}

// VolumeClaimPods represents a persistent volume together with belonging claim and pods.
type VolumeClaimPods struct {
	PersistentVolume      corev1.PersistentVolume      `json:"volume"`
	PersistentVolumeClaim corev1.PersistentVolumeClaim `json:"claim"`
	Pods                  []corev1.Pod                 `json:"pods"`
}

type patchStringValue struct {
	Op    string `json:"op"`
	Path  string `json:"path"`
	Value string `json:"value"`
}
