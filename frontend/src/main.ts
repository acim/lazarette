import { initPathRouter } from '@bjornlu/svelte-router'
import App from './App.svelte'
import PersistentVolumes from './pages/PersistentVolumes.svelte'
import StorageClasses from './pages/StorageClasses.svelte'

initPathRouter([
  { path: '/', component: PersistentVolumes },
  {
    path: '/classes',
    component: StorageClasses,
  },
])

const app = new App({
  target: document.getElementById('app'),
})

export default app
