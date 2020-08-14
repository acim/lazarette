<script type="ts">
  import type * as k8s from "@kubernetes/client-node";
  import { get } from "../fetch";
  import { volume } from "../store";

  interface Count {
    classes: number;
    volumes: number;
    claims: number;
  }

  interface Volumes {
    classes: k8s.V1StorageClass[];
    volumes: k8s.V1PersistentVolume[];
    claims: k8s.V1PersistentVolumeClaim[];
    count: Count;
  }

  const promise = get<Volumes>("/volumes");
  function setVolume(vol: k8s.V1PersistentVolume) {
    $volume = vol;
  }
</script>

<style>
  li:hover {
    cursor: pointer;
  }
</style>

{#await promise}
  <p>loading...</p>
{:then response}
  <ul>
    {#each response.parsedBody.classes as item}
      <li>{item.metadata.name}</li>
    {/each}
  </ul>
  <ul>
    {#each response.parsedBody.volumes as item}
      <li on:click={() => setVolume(item)}>{item.metadata.name}</li>
    {/each}
  </ul>
  <ul>
    {#each response.parsedBody.claims as item}
      <li>{item.metadata.name}</li>
    {/each}
  </ul>
{:catch error}
  <p class="text-error">{error.message}</p>
{/await}
