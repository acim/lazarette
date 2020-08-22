<script type="ts">
  import type { V1StorageClass } from "@kubernetes/client-node";

  export let storageClass: V1StorageClass;

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

<section>
  <h3>{storageClass.metadata.name}</h3>
  {#if isDefault()}
    <p>Default</p>
  {/if}
  <p>Provisioner: {storageClass.provisioner}</p>
  <p>Reclaim policy: {storageClass.reclaimPolicy}</p>
  <p>Allow expansion: {storageClass.allowVolumeExpansion}</p>
  <p>Binding mode: {storageClass.volumeBindingMode}</p>
</section>
