package main

import (
	"errors"
	"net/http"

	"github.com/acim/lazarette/pkg/k8s"
	"github.com/labstack/echo/v4"
	corev1 "k8s.io/api/core/v1"
	storagev1 "k8s.io/api/storage/v1"
)

type controller struct {
	k8s.Interface
}

func newController(k *k8s.Client) *controller {
	return &controller{Interface: k}
}

func (k *controller) classes(c echo.Context) error {
	scs, err := k.Interface.StorageClasses(c.Request().Context())
	if err != nil {
		c.Logger().Error(err)

		return errors.New("failed getting storage classes")
	}

	resp := resClasses{
		StorageClasses: scs,
	}

	return c.JSON(http.StatusOK, resp)
}

// func (k *client) setDefaultClass(c echo.Context) error {
// 	defaultClass := c.Param("name")
// 	ctx := c.Request().Context()

// 	scs, err := k.StorageV1().StorageClasses().List(ctx, metav1.ListOptions{})
// 	if err != nil {
// 		c.Logger().Error(err)

// 		return errors.New("failed getting storage classes")
// 	}

// 	path := fmt.Sprintf("/metadata/annotations/%s", jsonPointerEscape("storageclass.kubernetes.io/is-default-class"))

// 	for _, item := range scs.Items {
// 		var payload []patchStringValue

// 		switch item.GetName() {
// 		case defaultClass:
// 			payload = []patchStringValue{
// 				{
// 					Op:    "add",
// 					Path:  path,
// 					Value: "true",
// 				},
// 			}
// 		default:
// 			payload = []patchStringValue{
// 				{
// 					Op:   "remove",
// 					Path: path,
// 				},
// 			}
// 		}

// 		payloadJSON, err := json.Marshal(payload)
// 		if err != nil {
// 			c.Logger().Error(err)

// 			return errors.New("failed encoding json payload")
// 		}

// 		_, err = k.StorageV1().StorageClasses().Patch(
// 			ctx, item.GetName(), types.JSONPatchType, payloadJSON, metav1.PatchOptions{})
// 		if err != nil {
// 			c.Logger().Error(err)

// 			return fmt.Errorf("failed patching storage class %s", item.GetName())
// 		}
// 	}

// 	return k.classes(c)
// }

// func (k *client) togglePersistentVolumeReclaimPolicy(c echo.Context) error {
// 	persistentVolume := c.Param("name")
// 	reclaimPolicy := c.Param("policy")
// 	ctx := c.Request().Context()

// 	payload := []patchStringValue{
// 		{
// 			Op:    "replace",
// 			Path:  "/spec/persistentVolumeReclaimPolicy",
// 			Value: reclaimPolicy,
// 		},
// 	}

// 	payloadJSON, err := json.Marshal(payload)
// 	if err != nil {
// 		c.Logger().Error(err)

// 		return errors.New("failed encoding json payload")
// 	}

// 	_, err = k.CoreV1().PersistentVolumes().Patch(
// 		ctx, persistentVolume, types.JSONPatchType, payloadJSON, metav1.PatchOptions{})
// 	if err != nil {
// 		c.Logger().Error(err)

// 		return fmt.Errorf("failed patching persistent volume %s", persistentVolume)
// 	}

// 	return k.volumes(c)
// }

// func (k *client) volumes(c echo.Context) error {
// 	ctx := c.Request().Context()

// 	pvs, err := k.CoreV1().PersistentVolumes().List(ctx, metav1.ListOptions{})
// 	if err != nil {
// 		c.Logger().Error(err)

// 		return errors.New("failed getting persistent volumes")
// 	}

// 	pvcs, err := k.CoreV1().PersistentVolumeClaims("").List(ctx, metav1.ListOptions{})
// 	if err != nil {client
// 		c.Logger().Error(err)

// 		return errors.New("failed getting persistent volume claims")
// 	}

// 	pods, err := k.CoreV1().Pods("").List(ctx, metav1.ListOptions{})
// 	if err != nil {
// 		c.Logger().Error(err)

// 		return errors.New("failed getting pods")
// 	}

// 	// FieldSelector: fields.Set{"spec.volumes[].persistentVolumeClaim.claimName": "ghost-acim"}.AsSelector().String(),
// 	// _ = corev1.ReadWriteMany

// 	resp := resVolumes{
// 		Volumes: getVolumes(pvs.Items, pvcs.Items, pods.Items),
// 	}

// 	return c.JSON(http.StatusOK, resp)
// }

// func getVolumes(pvs []corev1.PersistentVolume, pvcs []corev1.PersistentVolumeClaim, pods []corev1.Pod) []volume {
// 	volumes := make([]volume, len(pvs))

// 	for i, pv := range pvs {
// 		pv.ManagedFields = nil
// 		volumes[i].PersistentVolume = pv

// 		for _, pvc := range pvcs {
// 			if pv.Name == pvc.Spec.VolumeName {
// 				pvc.ManagedFields = nil
// 				volumes[i].PersistentVolumeClaim = pvc

// 				volumes[i].Pods = make([]corev1.Pod, 0)

// 				for _, pod := range pods {
// 					for _, v := range pod.Spec.Volumes {
// 						if v.PersistentVolumeClaim != nil && pvc.Name == v.PersistentVolumeClaim.ClaimName {
// 							pod.ManagedFields = nil
// 							volumes[i].Pods = append(volumes[i].Pods, pod)

// 							break
// 						}
// 					}
// 				}

// 				break
// 			}
// 		}
// 	}

// 	return volumes
// }

type resClasses struct {
	StorageClasses []storagev1.StorageClass `json:"classes"`
}

type resVolumes struct {
	Volumes []volume `json:"volumes"`
}

type volume struct {
	PersistentVolume      corev1.PersistentVolume      `json:"volume"`
	PersistentVolumeClaim corev1.PersistentVolumeClaim `json:"claim"`
	Pods                  []corev1.Pod                 `json:"pods"`
}

type patchStringValue struct {
	Op    string `json:"op"`
	Path  string `json:"path"`
	Value string `json:"value"`
}
