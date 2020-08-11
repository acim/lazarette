import { writable } from "svelte/store";
import type * as k8s from "@kubernetes/client-node";

export const volume = writable<k8s.V1PersistentVolume | null>(null);
