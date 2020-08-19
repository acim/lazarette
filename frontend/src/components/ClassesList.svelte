<script type="ts">
  import type * as k8s from "@kubernetes/client-node";
  import { get } from "../fetch";
  import type { claim_component } from "svelte/internal";

  interface Classes {
    classes: k8s.V1StorageClass[];
    error: string;
  }

  const promise = get<Classes>("/classes");
</script>

<div class="container">
  {#await promise}
    <p>loading...</p>
  {:then response}
    <section>
      {#each response.parsedBody.classes as item}
        <p>{item.metadata.name}</p>
      {/each}
    </section>
  {:catch error}
    <p class="text-error">{error.message}</p>
  {/await}
</div>
