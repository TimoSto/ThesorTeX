# ThesorTeX
Tool for handling bibliography in a LaTeX project

## Downloads
The local application can be downloaded on the following links:

- [Windows](https://thesortex-artifacts.s3.eu-central-1.amazonaws.com/tool/latest/windows/ThesorTeX.zip)
- [Linux](https://thesortex-artifacts.s3.eu-central-1.amazonaws.com/tool/latest/linux/ThesorTeX.zip)
- [MacOS](https://thesortex-artifacts.s3.eu-central-1.amazonaws.com/tool/latest/mac/ThesorTeX.zip)
- [MacOS (Apple Silicon)](https://thesortex-artifacts.s3.eu-central-1.amazonaws.com/tool/latest/mac_silicon/ThesorTeX.zip)

## Build mechanics

This repository uses Bazel to build both the backends and the frontends.

### Production

The production build is triggered by calling the script `./scripts/build.sh`. It produces the following `artifacts`:

- ZIP files containing the executables of the app for Windows, Linux, MacOS and MacOS with arm (M1/M2)
- ZIP files containing the thesis and cv templates
- A ZIP file for the website lambda

### Development

To run the frontends on a watcher:

```
cd services/app/frontend 
or 
cd services/website/frontend

pnpm run dev
```

To run the app backend locally:

```
bazel run //services/app/cmd/e2e
```

To run the website backend locally;

```
bazel run //services/website/cmd/e2e
```

## Testing

The testing types are

- unit tests for the backend (Go)
- unit tests for the frontend (vitest)
- end to end tests for the app (playwright)

All of these tests can be run using bazel.

```
bazel test //...
```

## Deployment

To update aws infrastructure

```
cd terraform
./scripts/plan.sh
./scripts.apply.sh
```

To upload build-artifacts as zips

```
./scripts/s3_upload_atrifacts.sh all
```

To only deploy one target, replace `all` with `tool`, `thesisTemplate` or `cvTemplate`.

## Used technologies

- Go
- Pnpm
- Playwright
- Terraform
- Bazel

## Project structure

### services

Here the separation of the different services, that are independently runnable.

##### cmd

Here lie the targets for the executables

##### internal

Internal packages of each service

##### frontend

The frontend of the service using the global pnpm-workspace

### pkg

Packages used by multiple services

#### backend

Go packages

#### frontend

GUI packages

### tests

global tests

### terraform

terraform modules

### scripts

build scripts

### bazel

global bazel files
