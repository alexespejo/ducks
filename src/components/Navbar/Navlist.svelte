
<script lang="ts">
	import { onMount } from "svelte";
	import { AA_Old_Code } from "../../data/AA_Old_Code.js";
import { CS162_notes1 } from "../../data/CS162_notes1.js";
import { CS165_prjct1 } from "../../data/CS165_prjct1.js";
import { Testing } from "../../data/Testing.js";
import { Inf131_notes } from "../../data/Inf131_notes.js";
	

	export let page = "";

	let headers: string[] = [];
	onMount(() => {
		switch (page.toLowerCase()) {
			case 'aa_old_code': headers=AA_Old_Code; break;case 'cs162_notes1': headers=CS162_notes1; break;case 'cs165_prjct1': headers=CS165_prjct1; break;case 'testing': headers=Testing; break;case 'inf131_notes': headers=Inf131_notes; break;
			default:
				headers = [];
				break;
		}	
	});
</script>	

<ul class="fixed top-0 right-0 w-48">
 {#each headers as header}
  <li class="text-xs truncate">
   <a href={'#'+header.split(" ").join("-").toLowerCase()}>
    {header}
   </a>
  </li>
 {/each}
</ul>
	