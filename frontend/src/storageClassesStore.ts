import { writable, Readable } from "svelte/store";
import type { V1StorageClass } from "@kubernetes/client-node";
import { get, patch, HttpResponse } from "./fetch";

interface StorageClasses {
  classes?: V1StorageClass[];
  message?: string;
}

export interface StorageClassesReadable<T> extends Readable<T> {
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

const { subscribe, set } = writable<V1StorageClass[]>([]);

const store: StorageClassesReadable<V1StorageClass[]> = {
  subscribe,
  load: async () => {
    let res: HttpResponse<StorageClasses>;
    try {
      res = await get<StorageClasses>("/v1/classes.json");
      set(res?.parsedBody.classes);
    } catch (err) {
      throw new Error(
        res?.parsedBody.message ? res.parsedBody.message : err.message
      );
    }
  },
  setDefault: async (name: string) => {
    let res: HttpResponse<StorageClasses>;
    try {
      res = await patch<StorageClasses>(`/v1/classes/default/${name}`, null);
      set(res?.parsedBody.classes);
    } catch (err) {
      throw new Error(
        res?.parsedBody.message ? res.parsedBody.message : err.message
      );
    }
  },
};

export default store;
