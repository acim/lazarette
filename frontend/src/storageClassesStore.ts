import { writable, Writable } from "svelte/store";
import type { V1StorageClass } from "@kubernetes/client-node";
import { get, HttpResponse } from "./fetch";

interface StorageClasses {
  classes: V1StorageClass[];
  error: string;
}

export interface StorageClassesWritable<T> extends Writable<T> {
  /**
   * Load data from server.
   */
  load(): void;
}

const { subscribe, set, update } = writable<V1StorageClass[]>([]);

const store: StorageClassesWritable<V1StorageClass[]> = {
  subscribe,
  set,
  update,
  load: async () => {
    let res: HttpResponse<StorageClasses>;
    try {
      res = await get<StorageClasses>("/v1/classes.json");
      set(res.parsedBody.classes);
    } catch (err) {
      throw new Error(
        res?.parsedBody?.error !== "" ? res.parsedBody.error : err.message
      );
    }
  },
};

export default store;
