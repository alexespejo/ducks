<script lang="ts">
 export let name: string;
 export let size: number = 24;

 // The public directory is typically served at the root path
 const src = `/icons/${name}.svg`;
</script>

<img {src} alt={name} width={size} height={size} />

<style>
 img {
  display: inline-block;
  vertical-align: middle;
 }
</style>
