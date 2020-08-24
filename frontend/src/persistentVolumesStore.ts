import { writable, Writable } from "svelte/store";

import type {
  V1PersistentVolume,
  V1PersistentVolumeClaim,
  V1Pod,
} from "@kubernetes/client-node";
import { get, HttpResponse } from "./fetch";

export interface PersistentVolume {
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
   * Load data from server.
   */
  load(): void;
}

const { subscribe, set, update } = writable<PersistentVolume[]>([]);

const store: PersistentVolumesWritable<PersistentVolume[]> = {
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

export default store;
