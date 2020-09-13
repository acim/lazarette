package main

import (
	"os"

	"github.com/acim/lazarette/pkg/k8s"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"k8s.io/client-go/rest"
)

func main() {
	publicDir := os.Getenv("PUBLIC_DIR")
	if publicDir == "" {
		publicDir = "public"
	}

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	config, err := rest.InClusterConfig()
	if err != nil {
		e.Logger.Fatal(err)
	}

	client, err := k8s.NewClient(config)
	if err != nil {
		e.Logger.Fatal(err)
	}

	controller := newController(client)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:  publicDir,
		HTML5: true,
	}))

	e.GET("/v1/volumes.json", controller.volumes)
	e.GET("/v1/classes.json", controller.classes)
	// e.PATCH("/v1/classes/default/:name", controller.setDefaultClass)
	// e.PATCH("/v1/classes/policy/:name/:policy", controller.togglePersistentVolumeReclaimPolicy)

	e.Logger.Fatal(e.Start(":3000"))
}
