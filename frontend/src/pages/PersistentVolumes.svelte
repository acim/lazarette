<script type="ts">
  import type crayon from "crayon";
  import Nav from "../components/Nav.svelte";
  import PersistentVolume from "../components/PersistentVolume.svelte";
  import Icon from "mdi-svelte";
  import { mdiLoading } from "@mdi/js";
  import store from "../persistentVolumesStore";
  import { onMount } from "svelte";

  export let nav: crayon.Router;

  const color = getComputedStyle(document.documentElement).getPropertyValue(
    "--color-primary"
  );

  let error: string;
  let loading = false;

  onMount(() => {
    try {
      loading = true;
      store.load();
      loading = false;
    } catch (e) {
      error = e;
      loading = false;
    }
  });
</script>

<Nav {nav} />

<div class="container">
  {#if loading}
    <Icon path={mdiLoading} size="4rem" spin="2" {color} />
  {/if}
  {#each $store as item (item.volume.metadata.uid)}
    <PersistentVolume persistentVolume={item} />
  {/each}
  {#if error}
    <p class="text-error">{error}</p>
  {/if}
</div>
