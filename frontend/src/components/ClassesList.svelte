<script type="ts">
  import type * as k8s from "@kubernetes/client-node";
  import { get } from "../fetch";
  import StorageClass from "./Class.svelte";
  import Icon from "mdi-svelte";
  import { mdiLoading } from "@mdi/js";

  interface Classes {
    classes: k8s.V1StorageClass[];
    error: string;
  }

  const promise = get<Classes>("/v1/classes.json");

  const color = getComputedStyle(document.documentElement).getPropertyValue(
    "--color-primary"
  );
</script>

<div class="container">
  {#await promise}
    <Icon path={mdiLoading} size="4rem" spin="2" {color} />
  {:then response}
    {#each response.parsedBody.classes as item}
      <StorageClass storageClass={item} />
    {/each}
  {:catch error}
    <p class="text-error">{error.message}</p>
  {/await}
</div>
