import { writable, Readable } from "svelte/store";

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
  volumes?: PersistentVolume[];
  message?: string;
}

export interface PersistentVolumesReadable<T> extends Readable<T> {
  /**
   * Load data from server.
   */
  load(): void;
}

const { subscribe, set } = writable<PersistentVolume[]>([]);

const store: PersistentVolumesReadable<PersistentVolume[]> = {
  subscribe,
  load: async () => {
    let res: HttpResponse<PersistentVolumes>;
    try {
      res = await get<PersistentVolumes>("/v1/volumes.json");
      set(res?.parsedBody.volumes);
    } catch (err) {
      throw new Error(
        res?.parsedBody.message ? res.parsedBody.message : err.message
      );
    }
  },
};

export default store;
