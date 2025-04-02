// @ts-check
import { defineConfig } from "astro/config";
import remarkMath from "remark-math";
import rehypeKatex from "rehype-katex";
import mdx from "@astrojs/mdx";
import tailwindcss from "@tailwindcss/vite";

import svelte from "@astrojs/svelte";

import react from "@astrojs/react";

// https://astro.build/config
export default defineConfig({
 markdown: {
  remarkPlugins: [remarkMath],
  rehypePlugins: [rehypeKatex],
 },
 vite: {
  plugins: [tailwindcss()],
 },

 integrations: [mdx(), svelte(), react()],
});
