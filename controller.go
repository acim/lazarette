package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/acim/lazarette/pkg/k8s"
	"github.com/labstack/echo/v4"
	storagev1 "k8s.io/api/storage/v1"
)

type controller struct {
	k8s.Interface
}

func newController(k k8s.Interface) *controller {
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

func (k *controller) setDefaultClass(c echo.Context) error {
	defaultClass := c.Param("name")

	err := k.Interface.SetDefaultStorageClass(c.Request().Context(), defaultClass)
	if err != nil {
		c.Logger().Error(err)

		return fmt.Errorf("failed patching storage class")
	}

	return k.classes(c)
}

func (k *controller) setPersistentVolumeReclaimPolicy(c echo.Context) error {
	pvn := c.Param("name")
	p := c.Param("policy")

	err := k.Interface.SetPersistentVolumeReclaimPolicy(c.Request().Context(), pvn, p)
	if err != nil {
		c.Logger().Error(err)

		return fmt.Errorf("failed patching persistent volume %s", pvn)
	}

	return k.volumes(c)
}

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
