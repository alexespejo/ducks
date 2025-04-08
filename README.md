**Notion** meets **Markdown** meets my VS Code extension meets some Go CLI I built

### Start CLI

#### Build the CLI onto the project

```bash
bash build-duck.sh
```

#### Run CLI

```bash
duck start
```

### CLI Stuff

#### 🧱 Scaffold

- Adds and removes any added or deleted markdown files created respectively

#### 📃 Add Document

- Automatically add a Markdown page and runs build

#### 🧼 clean

- Removes leftover files and consolidates any missing navigation items

#### 🚀 publish

- Commits and pushes code to the git repo

#### ☄️ Astro Build

- Runs `yarn astro build` outputs an error if there is an issue with the build (will not deploy when pushed)

```

```
