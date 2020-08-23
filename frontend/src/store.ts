import { writable } from "svelte/store";
import type {
  V1StorageClass,
  V1PersistentVolume,
} from "@kubernetes/client-node";
import { get, HttpResponse } from "./fetch";

export const volume = writable<V1PersistentVolume | null>(null);

interface StorageClasses {
  classes: V1StorageClass[];
  error: string;
}

export const storageClasses = writable<V1StorageClass[]>([]);

export const loadStorageClasses = async () => {
  let res: HttpResponse<StorageClasses>;
  try {
    res = await get<StorageClasses>("/v1/classes.json");
    storageClasses.set(res.parsedBody.classes);
  } catch (err) {
    throw new Error(
      res?.parsedBody?.error !== "" ? res.parsedBody.error : err.message
    );
  }
};
