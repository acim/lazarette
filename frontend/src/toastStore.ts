import {writable, Readable} from "svelte/store";
import {get, patch, HttpResponse} from "./fetch";

interface Toast {
  message: string;
  duration?: number;
}

export interface ToastReadable<T> extends Readable<T> {
  /**
   * Set toast.
   */
  toast(t: T): void;
}

const {subscribe, set} = writable<Toast>({message: ""});

const store: ToastReadable<Toast> = {
  subscribe,
  toast: (t: Toast) => {
    if (!t.duration) {
      t.duration = 1000;
    }
    set(t);
    setTimeout(() => set(null), t.duration);
  },
};

export default store;
