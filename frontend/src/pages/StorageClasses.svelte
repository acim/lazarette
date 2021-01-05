<script lang="ts">
  import StorageClass from '../components/StorageClass.svelte'
  import Toast from '../components/Toast.svelte'
  import store from '../storageClassesStore'
  import { mdiLoading } from '@mdi/js'
  import { onMount } from 'svelte'

  let error: string
  let loading = true

  onMount(() => {
    try {
      store.load()
    } catch (err) {
      error = err
    }
    loading = false
  })
</script>

<div class="container">
  {#if loading}
    <svg viewBox="0 0 24 24">
      <path d={mdiLoading} />
    </svg>
    Volumes
  {/if}
  {#each $store as item, i (item.metadata.uid)}
    <StorageClass {i} />
  {:else}
    <p>No classes.</p>
  {/each}
  {#if error}
    <p class="text-error">{error}</p>
  {/if}
</div>

<Toast />
