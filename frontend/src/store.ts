import { writable, Writable } from "svelte/store";
import type {
  V1StorageClass,
  V1PersistentVolume,
  V1PersistentVolumeClaim,
  V1Pod,
} from "@kubernetes/client-node";
import { get, HttpResponse } from "./fetch";

export const volume = writable<V1PersistentVolume | null>(null);

interface StorageClasses {
  classes: V1StorageClass[];
  error: string;
}

interface PersistentVolume {
  volume: V1PersistentVolume;
  claim: V1PersistentVolumeClaim;
  pods: V1Pod[];
}

interface PersistentVolumes {
  volumes: PersistentVolume[];
  error: string;
}

export interface PersistentVolumesWritable<T> extends Writable<T> {
  /**
   * Set value and inform subscribers.
   */
  load(): void;
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

function createPersistentVolumes(): PersistentVolumesWritable<
  PersistentVolume[]
> {
  const { subscribe, set, update } = writable<PersistentVolume[]>([]);

  return {
    subscribe,
    set,
    update,
    load: async () => {
      let res: HttpResponse<PersistentVolumes>;
      try {
        res = await get<PersistentVolumes>("/v1/volumes.json");
        set(res.parsedBody.volumes);
      } catch (err) {
        throw new Error(
          res?.parsedBody?.error !== "" ? res.parsedBody.error : err.message
        );
      }
    },
  };
}

export const persistentVolumes = createPersistentVolumes();
