<script lang="ts">
 import { onMount } from "svelte";
 import { CS165_prjct1 } from "../../data/CS165_prjct1.js";
 import { Cs179_notes } from "../../data/Cs179_notes.js";

 export let page = "";

 let headers: string[] = [];
 onMount(() => {
  switch (page.toLowerCase()) {
   case "cs165_prjct1":
    headers = CS165_prjct1;
    break;
   case "cs179_notes":
    headers = Cs179_notes;
    break;
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
