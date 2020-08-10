package main

import (
	"fmt"
	"log"
	"net/http"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type intercept404 struct {
	http.ResponseWriter
	statusCode int
}

func (w *intercept404) Write(b []byte) (int, error) {
	if w.statusCode == http.StatusNotFound {
		return len(b), nil
	}
	if w.statusCode != 0 {
		w.WriteHeader(w.statusCode)
	}
	return w.ResponseWriter.Write(b)
}

func (w *intercept404) WriteHeader(statusCode int) {
	if statusCode >= 300 && statusCode < 400 {
		w.ResponseWriter.WriteHeader(statusCode)
		return
	}
	w.statusCode = statusCode
}

func spaFileServeFunc(dir string) func(http.ResponseWriter, *http.Request) {
	fileServer := http.FileServer(http.Dir(dir))
	return func(w http.ResponseWriter, r *http.Request) {
		wt := &intercept404{ResponseWriter: w}
		fileServer.ServeHTTP(wt, r)
		if wt.statusCode == http.StatusNotFound {
			r.URL.Path = "/"
			w.Header().Set("Content-Type", "text/html")
			fileServer.ServeHTTP(w, r)
		}
	}
}

func main() {
	http.HandleFunc("/volumes", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data": {"volumes": ["vol1", "vol2"]}}`))

		config, err := rest.InClusterConfig()
		if err != nil {
			panic(err.Error())
		}
		// creates the clientset
		clientset, err := kubernetes.NewForConfig(config)
		if err != nil {
			panic(err.Error())
		}
		pvs, err := clientset.CoreV1().PersistentVolumes().List(ctx, metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("There are %d pods in the cluster\n", len(pvs.Items))
	})
	http.HandleFunc("/", spaFileServeFunc("public"))

	log.Fatal(http.ListenAndServe(":3000", nil))
}
