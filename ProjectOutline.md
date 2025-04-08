# How to make it work

## Techstack?

- Astro static site generator
- Vercel deployment
- Go CLI for productivity
- Markdown for content
- ShadCN + Tailwindcss, styling

## Basic UI

- (MVP) Navbar for each **"chapter"** and subsection
- (MVP) Single Page document, all content renders on the landing page
- (POST) Multi-page, separate documents into their pages and sub-pages
- (POST) Multi-page, Landing page to house all documents

### Components

#### Navbar

- Render Lists and sublists
- Navigate up and down the right panel
- Collapsable?
- List item rendering
- Dynamcially change navbar ordering based on page ordering
- Alt. Homepage navigation and navbar
  - Homepage options auto generate based on existing pages

### Layout

- Pass in props to determin what values in the navbar to render

## CLI / Features

- (MVP) Automatically be able to add and delete documents
- (MVP) Auto commit and push code to a github repo
- (MVP) Auto generate the navbar for each heading
  - Requires a way to parse what pages and what headings appear in what order
- (MVP) Dynamically render the pages in order
- (MVP-M) Strict ordering of pages
  - Require to dynamically create the navbar based off any particular ordering of pages
- (MVP-M) Image parsing to centralize where images go
  - Image clean up, clear images from directory that are not in a document

# Post Trial 1

## Frontend

- Notion style multipage
  - Eventual side navbar, track page you're viewing
  - Centerd text
- Homepage navigation

## Cli

- Auto gen pages and nav config
- Image clean up
- Delete pages
- Track for updates
- Be able to add files to a folder and auto generate the routes based on those pages

## Todo

- Parse h1's into a navbar

  - Separe layout and navbars?
  - X Separate files for schema?

    - Dedicated javascript arrays associated to each file

  - Build layout pagess
