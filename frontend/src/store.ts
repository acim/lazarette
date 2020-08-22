import { writable } from "svelte/store";
import type {
  V1StorageClass,
  V1PersistentVolume,
} from "@kubernetes/client-node";
import { get } from "./fetch";

export const volume = writable<V1PersistentVolume | null>(null);

export interface StorageClasses {
  classes: V1StorageClass[];
  error: string;
}

export const storageClasses = writable<StorageClasses>({
  classes: [],
  error: "",
});

export const loadStorageClasses = async () => {
  try {
    const res = await get<StorageClasses>("/v1/classes.json");
    storageClasses.set({ classes: res.parsedBody.classes, error: "" });
  } catch (err) {
    storageClasses.set({ classes: [], error: err.message });
  }
};
