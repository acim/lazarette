package main

import (
	"log"
	"net/http"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatalln(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalln(err.Error())
	}

	client := newClient(clientset)

	http.HandleFunc("/volumes", client.volumes)
	http.HandleFunc("/classes", client.classes)
	http.HandleFunc("/", spaFileServeFunc("public"))

	log.Fatal(http.ListenAndServe(":3000", nil))
}
