package main

import (
	"cli/helpers/ioFile"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func coreFileCleanup() {
	duckExt := "../src/pages/ducks/"
	dataExt := "../src/data/"

	cmd := exec.Command("ls", duckExt)
	out, _ := cmd.Output()
	duckData := strings.Split(string(out), "\n")

	cmd = exec.Command("ls", dataExt)
	out, _ = cmd.Output()
	dataData := strings.Split(string(out), "\n")

	headers := []string{}
	for _, file := range duckData {
		if strings.Contains(file, ".md") {
			headers = append(headers, file[:len(file)-3])
		}
	}

	bFound := false
	for _, file := range dataData {
		if strings.Contains(file, ".js") {
			for i := range headers {
				if file == headers[i]+".js" {
					bFound = true
				}
			}
		}
		if !bFound {
			fmt.Println("Deleting: " + dataExt + file)
			os.Remove(dataExt + file)
		}
		bFound = false
	}
}
func extractHeadings(mdContent string) []string {
	re := regexp.MustCompile(`(?m)^#{1,2}\s+(.*)$`)

	matches := re.FindAllStringSubmatch(mdContent, -1)

	var headings []string
	for _, match := range matches {
		headings = append(headings, match[0]) // Entire matched heading
	}

	return headings
}

func buildSwitchCase(header string) string {
	return fmt.Sprintf(`case '%s': headers=%s; break;`, strings.ToLower(header), header)
}
func buildImports(header string) string {
	return fmt.Sprintf(`import { %s } from "../../data/%s.js";`, header, header) + "\n"
}
func buildNavbar(pImports string, pCases string) string {
	return fmt.Sprintf(`
		<script lang="ts">
		 import { onMount } from "svelte";
			%s	

			export let page = "";

			let headers: string[] = [];
			onMount(() => {
				switch (page.toLowerCase()) {
					%s
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
	`, pImports, pCases)
}

func coreBuildNavigation() {
	fmt.Println("Building Navigation...")
	ext := "../src/pages/ducks/"
	extData := "../src/data/"
	cmd := exec.Command("ls", ext)
	output, _ := cmd.Output()

	fmtOutput := strings.Split(string(output), "\n")
	cases := ""
	imports := ""
	coreFileCleanup()
	for _, file := range fmtOutput {
		if strings.Contains(file, ".md") {
			content, _ := os.ReadFile(ext + file)
			headers := extractHeadings(string(content))
			for i := range headers {
				headers[i] = headers[i][2:]
			}

			filExt := len(file) - 3
			newFile := extData + file[:filExt] + ".js"
			ioFile.CreateFile(newFile)

			cases += buildSwitchCase(file[:filExt])
			imports += buildImports(file[:filExt])

			// Create JavaScript array content
			jsContent := fmt.Sprintf("export const %s = [\n", file[:filExt])
			for _, header := range headers {
				jsContent += fmt.Sprintf("  \"%s\",\n", header)
			}
			jsContent += "];"

			// Write to file
			err := os.WriteFile(newFile, []byte(jsContent), 0644)
			if err != nil {
				fmt.Printf("Error writing to file %s: %v\n", newFile, err)
			} else {
				fmt.Printf("%s.js ✅\n", newFile)
			}
		}
	}

	navListContent := buildNavbar(imports, cases)
	err := os.WriteFile("../src/components/Navbar/Navlist.svelte", []byte(navListContent), 0644)
	if err != nil {
		fmt.Printf("Error writing to file %s: %v\n", navListContent, err)
	} else {
		fmt.Printf("Navlist.svelte ✅\n")
	}
}

func coreBuildDocument(docName string) {
	filePath := "../src/pages/ducks/" + docName + ".md"
	err := ioFile.CreateFile(filePath)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	content := fmt.Sprintf(`
---
layout: "../../layouts/LayoutSingle.astro"
title: %s
---
	`, docName)
	err = os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
		return
	}
	fmt.Printf("Document %s.md created successfully ✅\n", docName)
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Error: No arguments provided. One argument is required.")
		os.Exit(1)
	}

	switch args[0] {
	case "build":

		coreBuildNavigation()

	case "add":
		if len(args) < 2 {
			fmt.Println("Error: No document name provided. One argument is required.")
			os.Exit(1)
		} else if strings.Contains(args[1], ".") {
			fmt.Println("Error: Invalid document name. No '.' allowed.")
			os.Exit(1)
		}
		coreBuildDocument(args[1])
		coreBuildNavigation()

	case "clean":
		coreFileCleanup()

	case "publish":
		exec.Command("git", "add", "--all").Run()
		exec.Command("git", "qush", "update").Run()

	default:
		fmt.Println("Error: Invalid argument. Use 'build' or 'add'.")
		os.Exit(1)
	}
}
