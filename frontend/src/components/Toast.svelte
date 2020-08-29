<script type="ts">
  import store from "../toastStore";
</script>

<div class="toast" class:show={$store.message}>{$store.message}</div>
