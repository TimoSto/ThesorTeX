# ThesorTeX
Tool for handling bibliography in a LaTeX project

## artifacts
The build artifacts. You can download these and start them on your computer.

## Project structure

### backend
The backend services

#### app
local executable, which handles the file management. This way unnecessary traffic is avoided.

#### website
The lambda handles the website.

#### auth
Handling of authentication

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