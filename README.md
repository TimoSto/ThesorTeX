# ThesorTeX
Tool for handling bibliography in a LaTeX project

## Downloads
The local application can be downloaded on the following links:
- [Windows](https://thesortex-artifacts.s3.eu-central-1.amazonaws.com/v0.0.1/windows.zip)
- [Linux](https://thesortex-artifacts.s3.eu-central-1.amazonaws.com/v0.0.1/linux.zip)
- [MacOS](https://thesortex-artifacts.s3.eu-central-1.amazonaws.com/v0.0.1/mac.zip)
- [MacOS (Apple Silicon)](https://thesortex-artifacts.s3.eu-central-1.amazonaws.com/v0.0.1/mac_silicon.zip)

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