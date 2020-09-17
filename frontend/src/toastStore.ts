import { writable, Readable } from "svelte/store"

interface Toast {
  message: string
  duration?: number
}

export interface ToastReadable<T> extends Readable<T> {
  /**
   * Set toast.
   */
  set(t: T): void
}

const { subscribe, set } = writable<Toast>({ message: "" })

const store: ToastReadable<Toast> = {
  subscribe,
  set: (t: Toast) => {
    if (!t.duration) {
      t.duration = 1000
    }
    set(t)
    setTimeout(() => set({ message: "" }), t.duration)
  },
}

export default store
