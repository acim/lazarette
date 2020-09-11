# lazarette - Kubernetes Volumes Explorer

[![Build Status](https://drone.ablab.io/api/badges/acim/lazarette/status.svg)](https://drone.ablab.io/acim/lazarette)

This is Kubernetes tool which provides GUI to explore storage classes, persistent volumes and persistent volumes claims. Besides this, it allows you to set default storage class, which is the only action at the moment, but we plan to add much more features like rescuing released persistent volume, backup volumes, restore volumes, copy volumes, etc. We also plan to make volumes browser in order to access and view or backup single files.

Backend of the project is written in Go and frontend in Svelte using TypeScript and strict types.

![screenshot](screenshot.png?raw=true)

## Install

```bash
kubectl apply -f https://raw.githubusercontent.com/acim/lazarette/master/deploy.yaml
```

## Use

We provide no ingress at the moment, so in order to access the tool, you need to use port forwarding.

```bash
kubectl port-forward lazarette-xyz 3000:3000
```

And now you can point your browser to http:/localhost:3000

## Features

- Set default storage class
- Toggle persistent volumes reclaim policy to either Retain or Delete (Recycle is deprecated)

## Development

```bash
kind create cluster --config=dev/kind-config.yaml
helm install nfs stable/nfs-server-provisioner
kubectl apply -f dev/deploy.yaml
```

Access the application on http://172.18.0.2:3000 (change IP to your kind's node's IP)

## Links

[Patching Kubernetes Resources in Golang](https://dwmkerr.com/patching-kubernetes-resources-in-golang/)
