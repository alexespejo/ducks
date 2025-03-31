<script lang="ts">
 import LinkItem from "./LinkItem.svelte";
 export let lists: string[] = [];
 if (lists.length > 0 && $$slots.default) {
  throw new Error(
   "Both 'lists' and slot content cannot be provided at the same time."
  );
 }
</script>

<ul>
 {#if !$$slots.default}
  {#each lists as list}
   <LinkItem label={list} />
  {/each}
 {:else}
  <slot />
 {/if}
</ul>
