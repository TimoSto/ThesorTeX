# ThesorTeX
Tool for handling bibliography in a LaTeX project

## Project structure

### cmd
Entrypoints of the application. app is for a local executable. In a future release a lambda will be added for deployment to aws.

### bib_management
Everything regarding the management of the bibliography in the backend.

#### backend
Package that is available to cmd

#### internal
Package that is only available to backend on same level

### pkg
Packages used by multiple other packages

### webclient
Gui with Vue3 and vuetify.

### e2e
end-2-end-tests

#### cypress
e2e with cypress

#### playwright
e2e with playwright
