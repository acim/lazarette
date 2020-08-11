import { writable } from "svelte/store";
import * as k8s from "@kubernetes/client-node";

export const volume = writable<k8s.V1PersistentVolume>(
  new k8s.V1PersistentVolume()
);
