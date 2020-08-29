<script type="ts">
  import Icon from "mdi-svelte";
  import Nav from "../components/Nav.svelte";
  import StorageClass from "../components/StorageClass.svelte";
  import Toast from "../components/Toast.svelte";
  import store from "../storageClassesStore";
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
    } catch (err) {
      error = err;
    }
    loading = false;
  });
</script>

<Nav {nav} />

<div class="container">
  {#if loading}
    <Icon path={mdiLoading} size="4rem" spin="2" {color} />
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
