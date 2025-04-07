<script>
 export let label;

 function navigate() {
  window.location.hash = `/ducks/${label
   .replace(/[^a-zA-Z0-9 \-]/g, "")
   .split(" ")
   .join("-")
   .toLowerCase()}`;
 }
</script>

<li>
 <button on:click={navigate}>
  {label}
 </button>
</li>
