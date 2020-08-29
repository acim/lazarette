<script type="ts">
  import Icon from "mdi-svelte";
  import store from "../storageClassesStore";
  import toast from "../toastStore";
  import type { V1StorageClass } from "@kubernetes/client-node";
  import { fade } from "svelte/transition";

  export let i: number;

  const setDefault = () => {
    try {
      const name = $store[i].metadata.name;
      store.setDefault(name);
      toast.set({ message: `Default storage class set to ${name}` });
    } catch (err) {
      toast.set({ message: (err as Error).message });
    }
  };

  let isDefault: () => boolean;
  $: isDefault = () => {
    return (
      $store[i].metadata.annotations.hasOwnProperty(
        "storageclass.kubernetes.io/is-default-class"
      ) &&
      $store[i].metadata.annotations[
        "storageclass.kubernetes.io/is-default-class"
      ] === "true"
    );
  };
</script>

<style>
  table {
    margin-bottom: 1rem;
  }
  button {
    margin-left: auto;
  }
</style>

<section class:position-relative={isDefault()} transition:fade>
  <h3>{$store[i].metadata.name}</h3>
  {#if isDefault()}
    <tag>default</tag>
  {/if}
  <table>
    <tr>
      <td>Provisioner</td>
      <td>{$store[i].provisioner}</td>
    </tr>
    <tr>
      <td>Reclaim policy</td>
      <td>{$store[i].reclaimPolicy}</td>
    </tr>
    <tr>
      <td>Allow expansion</td>
      <td>{$store[i].allowVolumeExpansion}</td>
    </tr>
    <tr>
      <td>Binding mode</td>
      <td>{$store[i].volumeBindingMode}</td>
    </tr>
  </table>
  {#if !isDefault()}
    <button
      on:click|once={() => {
        setDefault();
      }}>
      Set to default
    </button>
  {/if}
</section>
