<script type="ts">
  import type * as k8s from "@kubernetes/client-node";
  import { get } from "../fetch";
  import StorageClass from "./Class.svelte";

  interface Classes {
    classes: k8s.V1StorageClass[];
    error: string;
  }

  const promise = get<Classes>("/v1/classes.json");
</script>

<div class="container">
  {#await promise}
    <p>loading...</p>
  {:then response}
    {#each response.parsedBody.classes as item}
      <StorageClass storageClass={item} />
    {/each}
  {:catch error}
    <p class="text-error">{error.message}</p>
  {/await}
</div>
