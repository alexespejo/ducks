package main

import (
	"cli/helpers/ioFile"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/AlecAivazis/survey/v2"
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
				if strings.ToLower(file) == strings.ToLower(headers[i])+".js" {
					bFound = true
				}
			}
		}
		if !bFound && len(file) > 0 {
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
		heading := strings.TrimSpace(match[0])
		headings = append(headings, heading)
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

<ul class="fixed top-0 right-0 w-48">
 {#each headers as header}
  <li class="text-xs truncate">
   <a href={'#'+header.split(" ").join("-").toLowerCase()}>
    {header}
   </a>
  </li>
 {/each}
</ul>
	`, pImports, pCases)
}
func coreBuildHome() {
	filePath := "../src/components/Navbar/HomeNav.svelte"
	duckExt := "../src/pages/ducks/"
	cmd := exec.Command("ls", duckExt)
	out, _ := cmd.Output()
	duckData := strings.Split(string(out), "\n")

	routes := []string{}
	for _, header := range duckData {
		if strings.Contains(header, ".md") {
			routes = append(routes, header[:len(header)-3])
		}
	}
	for i := range routes {
		routes[i] = fmt.Sprintf(`"%s"`, routes[i])
	}
	content := fmt.Sprintf(`
	<script>
	import List from "./List.svelte";
	const routes = [%s];
	</script>
	<List lists={routes} />
	
		 `, strings.Join(routes, ", "))
	os.WriteFile(filePath, []byte(content), 0644)
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
				if len(headers[i]) > 2 {
					headers[i] = strings.TrimSpace(headers[i][2:])
				}
			}

			filExt := len(file) - 3
			fileName := strings.ToUpper(file[:1]) + file[1:filExt]
			newFile := extData + fileName + ".js"
			ioFile.CreateFile(newFile)
			fmt.Println(fileName)

			cases += buildSwitchCase(fileName)
			imports += buildImports(fileName)

			// Create JavaScript array content
			jsContent := fmt.Sprintf("export const %s = [\n", fileName)
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
	coreBuildHome()
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

func printError(message string) {
	fmt.Printf("\033[1;31mError: %s\033[0m\n", message)
}

var (
	optionBuild      string = "🏗️  Scaffold"
	optionAdd        string = "📃 Create Document"
	optionPublish    string = "🚀 Publish"
	optionAstroBuild string = "🧪 Test Build"
	optionAstroDev   string = "💻 Develop"
	optionClean      string = "🧼 Clean"
	optionExit       string = "🚪 Exit"
)

func processPrompt(selected string) {
	switch selected {
	case optionBuild:

		coreBuildNavigation()

	case optionAdd:
		fmt.Print("Enter the name of the document to add 📄: ")
		var docName string
		fmt.Scanln(&docName)
		coreBuildDocument(docName)
		coreBuildNavigation()

	case optionClean:
		coreBuildNavigation()

	case optionPublish:
		fmt.Println("Getting some ducks in a row 🦆...")
		err1 := exec.Command("git", "add", "--all").Run()
		err2 := exec.Command("git", "qush", "update").Run()
		if err1 == nil && err2 == nil {
			fmt.Println("Published! 🚀")
		}

	case optionExit:
		fmt.Println("Exiting...")
		os.Exit(0)

	case optionAstroBuild:
		exec.Command("../src/run-astro.sh", "build").Run()
		cmd := exec.Command("bash", "../src/run-astro.bash", "build")
		_, err := cmd.Output()
		if err != nil {
			fmt.Printf("Probably broken build!: %v\n", err)
		} else {
			fmt.Println("Build Good ✅")
		}

	case optionAstroDev:
		cmd := exec.Command("yarn", "astro", "dev")
		out, err := cmd.Output()
		if err != nil {
			fmt.Printf("Error running command: %v\n", err)
		} else {
			fmt.Println(out)
		}
	default:
		printError("Error: Invalid argument. Use 'build' or 'add'. " + selected)
		os.Exit(1)
	}
}
func main() {

	if len(os.Args) > 2 {
		printError("Error: Too many arguments provided")
		os.Exit(1)
	}

	args := os.Args[1:]
	dir, _ := os.Getwd()
	directories := strings.Split(dir, "/")
	lastDir := directories[len(directories)-1]
	if lastDir != "cli" {
		printError("Error: Please run this command from the 'cli' directory")
		os.Exit(1)
	}

	options := []string{optionAdd, optionBuild, optionPublish, optionAstroBuild, optionExit}
	var defaultVal string = options[0]
	if args[0] == "start" {
		optionBuild = "🏗️  Scaffold"
		optionAdd = "📃 Create Document"
		optionPublish = "🚀 Publish"
		optionAstroBuild = "🧪 Test Build"
		optionAstroDev = "💻 Develop"
		optionClean = "🧼 Clean"
		optionExit = "🚪 Exit"
		for {
			var selected string

			prompt := &survey.Select{
				Message: "🦆",

				Options: options,
				VimMode: true,
				Default: defaultVal,
			}
			err := survey.AskOne(prompt, &selected)
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				continue
			}
			defaultVal = selected

			processPrompt(selected)
		}
	} else {
		optionBuild = "build"
		optionAdd = "add"
		optionPublish = "publish"
		optionAstroBuild = "test"
		optionAstroDev = "develop"
		optionClean = "clean"
		optionExit = "exit"

		processPrompt(args[0])
	}
}
