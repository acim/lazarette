import { writable, Writable } from "svelte/store";
import type { V1StorageClass } from "@kubernetes/client-node";
import { get, patch, HttpResponse } from "./fetch";

interface Error {
  error: string;
}

interface StorageClasses extends Error {
  classes: V1StorageClass[];
}

export interface StorageClassesWritable<T> extends Writable<T> {
  /**
   * Load data from server.
   */
  load(): void;
  /**
   * Set default storage class.
   *
   * @param name storage class name.
   */
  setDefault(name: string): void;
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
      set(res?.parsedBody.classes);
    } catch (err) {
      throw new Error(
        res?.parsedBody.error ? res.parsedBody.error : err.message
      );
    }
  },
  setDefault: async (name: string) => {
    let res: HttpResponse<Error>;
    try {
      res = await patch<Error>(`/v1/classes/default/${name}`, null);
      store.load();
    } catch (err) {
      throw new Error(
        res?.parsedBody.error ? res.parsedBody.error : err.message
      );
    }
  },
};

export default store;
