package main

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"k8s.io/client-go/kubernetes"
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

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		e.Logger.Fatal(err)
	}

	client := newClient(clientset)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:  publicDir,
		HTML5: true,
	}))

	e.GET("/v1/volumes.json", client.volumes)
	e.GET("/v1/classes.json", client.classes)
	e.PATCH("/v1/classes/default/:name", client.setDefaultClass)
	e.PATCH("/v1/classes/policy/:name/:policy", client.togglePersistentVolumeReclaimPolicy)

	e.Logger.Fatal(e.Start(":3000"))
}
