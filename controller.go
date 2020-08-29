package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	corev1 "k8s.io/api/core/v1"
	storagev1 "k8s.io/api/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
)

type client struct {
	kubernetes.Interface
}

func newClient(clientset kubernetes.Interface) *client {
	return &client{Interface: clientset}
}

func (k *client) classes(c echo.Context) error {
	scs, err := k.StorageV1().StorageClasses().List(c.Request().Context(), metav1.ListOptions{})
	if err != nil {
		c.Logger().Error(err)

		return errors.New("failed getting storage classes")
	}

	for i := range scs.Items {
		scs.Items[i].ManagedFields = nil
	}

	resp := resClasses{
		StorageClasses: scs.Items,
	}

	return c.JSON(http.StatusOK, resp)
}

func (k *client) setDefaultClass(c echo.Context) error {
	defaultClass := c.Param("name")
	ctx := c.Request().Context()

	scs, err := k.StorageV1().StorageClasses().List(ctx, metav1.ListOptions{})
	if err != nil {
		c.Logger().Error(err)

		return errors.New("failed getting storage classes")
	}

	path := fmt.Sprintf("/metadata/annotations/%s", jsonPointerEscape("storageclass.kubernetes.io/is-default-class"))

	for _, item := range scs.Items {
		var payload []patchStringValue

		switch item.GetName() {
		case defaultClass:
			payload = []patchStringValue{
				{
					Op:    "add",
					Path:  path,
					Value: "true",
				},
			}
		default:
			payload = []patchStringValue{
				{
					Op:   "remove",
					Path: path,
				},
			}
		}

		payloadJSON, err := json.Marshal(payload)
		if err != nil {
			c.Logger().Error(err)

			return errors.New("failed encoding json payload")
		}

		_, err = k.StorageV1().StorageClasses().Patch(
			ctx, item.GetName(), types.JSONPatchType, payloadJSON, metav1.PatchOptions{})
		if err != nil {
			c.Logger().Error(err)

			return fmt.Errorf("failed patching storage class %s", item.GetName())
		}
	}

	return k.classes(c)
}

func (k *client) volumes(c echo.Context) error {
	ctx := c.Request().Context()

	pvs, err := k.CoreV1().PersistentVolumes().List(ctx, metav1.ListOptions{})
	if err != nil {
		c.Logger().Error(err)

		return errors.New("failed getting persistent volumes")
	}

	pvcs, err := k.CoreV1().PersistentVolumeClaims("").List(ctx, metav1.ListOptions{})
	if err != nil {
		c.Logger().Error(err)

		return errors.New("failed getting persistent volume claims")
	}

	pods, err := k.CoreV1().Pods("").List(ctx, metav1.ListOptions{})
	if err != nil {
		c.Logger().Error(err)

		return errors.New("failed getting pods")
	}

	// FieldSelector: fields.Set{"spec.volumes[].persistentVolumeClaim.claimName": "ghost-acim"}.AsSelector().String(),
	// _ = corev1.ReadWriteMany

	resp := resVolumes{
		Volumes: getVolumes(pvs.Items, pvcs.Items, pods.Items),
	}

	return c.JSON(http.StatusOK, resp)
}

func getVolumes(pvs []corev1.PersistentVolume, pvcs []corev1.PersistentVolumeClaim, pods []corev1.Pod) []volume {
	volumes := make([]volume, len(pvs))

	for i, pv := range pvs {
		pv.ManagedFields = nil
		volumes[i].PersistentVolume = pv

		for _, pvc := range pvcs {
			if pv.Name == pvc.Spec.VolumeName {
				pvc.ManagedFields = nil
				volumes[i].PersistentVolumeClaim = pvc

				volumes[i].Pods = make([]corev1.Pod, 0)
				for _, pod := range pods {
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

	return volumes
}

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

func jsonPointerEscape(s string) string {
	s = strings.Replace(s, "~", "~0", -1)
	s = strings.Replace(s, "/", "~1", -1)

	return s
}
