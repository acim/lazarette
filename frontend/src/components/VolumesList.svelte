<script type="ts">
  import type * as k8s from "@kubernetes/client-node";

  interface HttpResponse<T> extends Response {
    parsedBody?: T;
  }

  const promise = http<k8s.V1PersistentVolume[]>("/volumes");

  async function http<T>(request: RequestInfo): Promise<HttpResponse<T>> {
    const response: HttpResponse<T> = await fetch(request);

    response.parsedBody = await response.json();
    return response;
  }
</script>

{#await promise}
  <p>loading</p>
{:then response}
  <ul>
    {#each response.parsedBody as volume}
      <li>{volume.metadata.name}</li>
    {/each}
  </ul>
{:catch error}
  <p style="color: red">{error.message}</p>
{/await}
