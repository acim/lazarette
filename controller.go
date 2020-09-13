package main

import (
	"errors"
	"net/http"

	"github.com/acim/lazarette/pkg/k8s"
	"github.com/labstack/echo/v4"
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

	resp := classes{
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

func (k *controller) volumes(c echo.Context) error {
	vcps, err := k.Interface.VolumesWithClaimsAndPods(c.Request().Context())
	if err != nil {
		c.Logger().Error(err)

		return errors.New("failed getting persistent volumes")
	}

	resp := volumes{
		Volumes: vcps,
	}

	return c.JSON(http.StatusOK, resp)
}

type classes struct {
	StorageClasses []storagev1.StorageClass `json:"classes"`
}

type volumes struct {
	Volumes []k8s.VolumeClaimPods `json:"volumes"`
}

type patchStringValue struct {
	Op    string `json:"op"`
	Path  string `json:"path"`
	Value string `json:"value"`
}
