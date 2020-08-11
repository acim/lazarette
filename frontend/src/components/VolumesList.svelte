<script type="ts">
  import type * as k8s from "@kubernetes/client-node";
  import { get } from "../fetch";

  const promise = get<k8s.V1PersistentVolume[]>("/volumes");
</script>

{#await promise}
  <p>loading...</p>
{:then response}
  <ul>
    {#each response.parsedBody as volume}
      <li>{volume.metadata.name}</li>
    {/each}
  </ul>
{:catch error}
  <p class="text-error">{error.message}</p>
{/await}
