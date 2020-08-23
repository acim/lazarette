package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	e := echo.New()

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
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:  "public",
		HTML5: true,
	}))

	e.GET("/v1/volumes.json", client.volumes)
	e.GET("/v1/classes.json", client.classes)

	e.Logger.Fatal(e.Start(":3000"))
}
