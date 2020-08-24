<script lang="ts">
  import { fade } from "svelte/transition";
  import type { PersistentVolume } from "../persistentVolumesStore";

  export let persistentVolume: PersistentVolume;
</script>

<style>
  ul {
    margin-top: 2rem;
  }
</style>

<section transition:fade>
  <h3>{persistentVolume.volume.metadata.name}</h3>
  <table>
    <tr>
      <td>Storage class</td>
      <td>{persistentVolume.volume.spec.storageClassName}</td>
    </tr>
    <tr>
      <td>Capacity</td>
      <td>{persistentVolume.volume.spec.capacity.storage}</td>
    </tr>
    <tr>
      <td>Mode</td>
      <td>{persistentVolume.volume.spec.accessModes}</td>
    </tr>
    <tr>
      <td>Status</td>
      <td>{persistentVolume.volume.status.phase}</td>
    </tr>
    <tr>
      <td>Reference claim kind</td>
      <td>{persistentVolume.volume.spec.claimRef.kind}</td>
    </tr>
    <tr>
      <td>Referencing claim name</td>
      <td>
        {persistentVolume.volume.spec.claimRef.namespace}/{persistentVolume.volume.spec.claimRef.name}
      </td>
    </tr>
    <tr>
      <td>Associated claim name</td>
      <td>
        {persistentVolume.claim?.metadata.namespace}/{persistentVolume.claim.metadata.name}
      </td>
    </tr>
    <tr>
      <td>Associated claim capacity</td>
      <td>{persistentVolume.claim?.status.capacity.storage}</td>
    </tr>
    <tr>
      <td>Associated claim modes</td>
      <td>{persistentVolume.claim?.status.accessModes}</td>
    </tr>
    <tr>
      <td>Associated claim status</td>
      <td>{persistentVolume.claim?.status.phase}</td>
    </tr>
    <tr>
      <td>Mounted by pod</td>
      <td>{persistentVolume.pods[0]?.metadata.name}</td>
    </tr>
  </table>
</section>
