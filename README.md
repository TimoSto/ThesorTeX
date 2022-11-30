# ThesorTeX
Tool for handling bibliography in a LaTeX project

## Project structure

### cmd
Entrypoints of the application.

#### app
local executable, which handles the file management. This way unnecessary traffic is avoided.

#### lambda
The lambda handles authentication and the website.

### services
Intependently runnable services. Note that they might be combined into a single executable.
But they all have their own backend package and can be invoked by creating a mux server and give it as parameter to the start server func

#### app
Managing of bibliograhy and tex-files

#### auth
Authentication

#### website
The website accessible to global web

### pkg
Packages used by multiple other packages

### frontend
Gui with Vue3 and vuetify.
- The entry files are in a directory on root
- components and rules are managed globally for all pages
- the api-calls and other implementation is located in the pages-directories

### e2e
end-2-end-tests

### mocks
mock-implementations for unit-tests