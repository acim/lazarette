package main

import (
	"encoding/json"
	"log"
	"net/http"

	corev1 "k8s.io/api/core/v1"
	storagev1 "k8s.io/api/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type client struct {
	*kubernetes.Clientset
}

func newClient(clientset *kubernetes.Clientset) *client {
	return &client{Clientset: clientset}
}

func (c *client) volumes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	w.Header().Set("Content-Type", "application/json")

	scs, err := c.StorageV1().StorageClasses().List(ctx, metav1.ListOptions{})
	if err != nil {
		httpError(w, err, "failed getting storage classes")

		return
	}

	pvs, err := c.CoreV1().PersistentVolumes().List(ctx, metav1.ListOptions{})
	if err != nil {
		httpError(w, err, "failed getting persistent volumes")

		return
	}

	pvcs, err := c.CoreV1().PersistentVolumeClaims("").List(ctx, metav1.ListOptions{})
	if err != nil {
		httpError(w, err, "failed getting persistent volume claims")

		return
	}

	resp := res{
		StorageClasses:         scs.Items,
		PersistentVolumes:      pvs.Items,
		PersistentVolumeClaims: pvcs.Items,
		Count: count{
			StorageClasses:         len(scs.Items),
			PersistentVolumes:      len(pvs.Items),
			PersistentVolumeClaims: len(pvcs.Items),
		},
	}

	res, err := json.Marshal(resp)
	if err != nil {
		httpError(w, err, "failed encoding to json")

		return
	}

	w.Write(res) //nolint:errcheck
}

func httpError(w http.ResponseWriter, err error, text string) {
	log.Printf("%s: %v\n", text, err)

	r := res{
		Error: &text,
	}

	res, err := json.Marshal(r)
	if err != nil {
		log.Println(err)
	}

	w.WriteHeader(http.StatusInternalServerError)
	w.Write(res) //nolint:errcheck
}

type res struct {
	StorageClasses         []storagev1.StorageClass       `json:"classes"`
	PersistentVolumes      []corev1.PersistentVolume      `json:"volumes"`
	PersistentVolumeClaims []corev1.PersistentVolumeClaim `json:"claims"`
	Count                  count                          `json:"count"`
	Error                  *string                        `json:"error,omitempty"`
}

type count struct {
	StorageClasses         int `json:"classes"`
	PersistentVolumes      int `json:"volumes"`
	PersistentVolumeClaims int `json:"claims"`
}
