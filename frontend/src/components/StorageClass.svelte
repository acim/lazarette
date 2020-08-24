<script type="ts">
  import type { V1StorageClass } from "@kubernetes/client-node";
  import { fade } from "svelte/transition";

  export let storageClass: V1StorageClass;
  export let setDefault: (name: string) => void;

  const isDefault = (): boolean => {
    return (
      storageClass.metadata.annotations.hasOwnProperty(
        "storageclass.kubernetes.io/is-default-class"
      ) &&
      storageClass.metadata.annotations[
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
  <h3>{storageClass.metadata.name}</h3>
  {#if isDefault()}
    <tag>default</tag>
  {/if}
  <table>
    <tr>
      <td>Provisioner</td>
      <td>{storageClass.provisioner}</td>
    </tr>
    <tr>
      <td>Reclaim policy</td>
      <td>{storageClass.reclaimPolicy}</td>
    </tr>
    <tr>
      <td>Allow expansion</td>
      <td>{storageClass.allowVolumeExpansion}</td>
    </tr>
    <tr>
      <td>Binding mode</td>
      <td>{storageClass.volumeBindingMode}</td>
    </tr>
  </table>
  {#if !isDefault()}
    <button on:click={() => setDefault(storageClass.metadata.name)}>
      Set as default
    </button>
  {/if}
</section>
