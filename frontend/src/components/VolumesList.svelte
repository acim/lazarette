<script type="ts">
  import type * as k8s from "@kubernetes/client-node";
  import { get } from "../fetch";
  import { volume } from "../store";

  const promise = get<k8s.V1PersistentVolume[]>("/volumes");
  function setVolume(vol: k8s.V1PersistentVolume) {
    console.log(vol);
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
    {#each response.parsedBody as vol}
      <li on:click={() => setVolume(vol)}>{vol.metadata.name}</li>
    {/each}
  </ul>
{:catch error}
  <p class="text-error">{error.message}</p>
{/await}
