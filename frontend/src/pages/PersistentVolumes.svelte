<script type="ts">
  import Icon from "mdi-svelte";
  import Nav from "../components/Nav.svelte";
  import PersistentVolume from "../components/PersistentVolume.svelte";
  import Toast from "../components/Toast.svelte";
  import store from "../persistentVolumesStore";
  import type crayon from "crayon";
  import { mdiLoading } from "@mdi/js";
  import { onMount } from "svelte";

  export let nav: crayon.Router;

  const color = getComputedStyle(document.documentElement).getPropertyValue(
    "--color-primary"
  );

  let error: string;
  let loading = true;

  onMount(() => {
    try {
      store.load();
    } catch (e) {
      error = e;
    }
    loading = false;
  });
</script>

<Nav {nav} />

<div class="container">
  {#if loading}
    <Icon path={mdiLoading} size="4rem" spin="2" {color} />
  {/if}
  {#each $store as item, i (item.volume.metadata.uid)}
    <PersistentVolume {i} />
  {:else}
    <p>No volumes.</p>
  {/each}
  {#if error}
    <p class="text-error">{error}</p>
  {/if}
</div>

<Toast />
