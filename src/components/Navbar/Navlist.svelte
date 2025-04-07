<script lang="ts">
 import { onMount } from "svelte";

 export let page = "";

 let headers: string[] = [];
 onMount(() => {
  switch (page.toLowerCase()) {
   default:
    headers = [];
    break;
  }
 });
</script>

<ul class="fixed top-0 right-0 w-48">
 {#each headers as header}
  <li class="text-xs truncate">
   <a href={"#" + header.split(" ").join("-").toLowerCase()}>
    {header}
   </a>
  </li>
 {/each}
</ul>
