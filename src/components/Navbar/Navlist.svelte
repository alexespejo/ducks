
		<script lang="ts">
		 import { onMount } from "svelte";
			import { Doc1 } from "../../data/Doc1.js";
import { Doc2 } from "../../data/Doc2.js";
import { Doc3 } from "../../data/Doc3.js";
import { Doc4 } from "../../data/Doc4.js";
import { GoRefSheet } from "../../data/GoRefSheet.js";
	

			export let page = "";

			let headers: string[] = [];
			onMount(() => {
				switch (page.toLowerCase()) {
					case 'doc1': headers=Doc1; break;case 'doc2': headers=Doc2; break;case 'doc3': headers=Doc3; break;case 'doc4': headers=Doc4; break;case 'gorefsheet': headers=GoRefSheet; break;
					default:
						headers = [];
						break;
				}	
			});
		</script>	

		<div>
			{#each headers as header}
				{header}
			{/each}
		</div>
	