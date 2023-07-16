<script>
  export let text = "";

  let isVisible = false;

  function showTooltip() {
    isVisible = true;
  }

  function hideTooltip() {
    isVisible = false;
  }
</script>

<span on:mouseenter={showTooltip} on:mouseleave={hideTooltip}>
  <slot />
  {#if isVisible}
    <div class="tooltip">{text}</div>
  {/if}
</span>

<style>
  .tooltip {
    position: absolute;
    z-index: 1;
    background-color: #000;
    color: #fff;
    padding: 8px;
    border-radius: 4px;
    font-size: 14px;
  }
</style>
