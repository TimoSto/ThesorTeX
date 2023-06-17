# ThesorTeX
A Template for a thesis, a cv and a Tool for handling bibliography in a LaTeX project.

The website is available under [https://thesortex.com](https://thesortex.com).

## Target consumers
Consumers of these templates and tools should have basic knowledge in the use of LaTeX. The target audience are people, who want to use LaTeX to create papers with an academic claim.

Roughly summarized, the consumers have the following benefits from using *ThesorTeX*:

- Have a nice template meeting the common requirements for academic papers
- No need to worry about the TeX code for individual formatting of bibliography entries and cites
- Format your bibliography entries according to your wishes, without deep LaTeX knowledge
- Have a nice template for a CV
- Be part of the [community](https://github.com/TimoSto/ThesorTeX/discussions/158) to help improve *ThesorTeX* and increase your knowledge in the world of *LaTeX*

## Downloads
The template for academic papers can be downloaded [here](https://thesortex-artifacts.s3.eu-central-1.amazonaws.com/thesisTemplate/latest/ThesisTemplate.zip)

The local application can be downloaded on the following links:

- [Windows](https://thesortex-artifacts.s3.eu-central-1.amazonaws.com/thesisTool/latest/windows/ThesorTeX.zip)
- [Linux](https://thesortex-artifacts.s3.eu-central-1.amazonaws.com/thesisTool/latest/linux/ThesorTeX.zip)
- [MacOS](https://thesortex-artifacts.s3.eu-central-1.amazonaws.com/thesisTool/latest/mac/ThesorTeX.zip)
- [MacOS (Apple Silicon)](https://thesortex-artifacts.s3.eu-central-1.amazonaws.com/thesisTool/latest/mac_silicon/ThesorTeX.zip)

The template for a CV can be downloaded [here](https://thesortex-artifacts.s3.eu-central-1.amazonaws.com/cvTemplate/latest/CVTemplate.zip)

## Build mechanics

This repository uses Bazel to build both the backends and the frontends. But you need to have LaTeX installed on your machine to build everything.

### Production

The production build is triggered by calling the script `./scripts/build.sh` with the parameter `thesisTemplate`, `thesisTool`, `cvTemplate`or `all`. It produces the following `artifacts`:

- ZIP files containing the executables of the thesisTool for Windows, Linux, MacOS and MacOS with arm (M1/M2)
- ZIP files containing the thesis and cv templates
- A ZIP file for the website lambda

### Development

To run the frontends on a watcher:

```
cd services/thesisTool/frontend 
or 
cd services/website/frontend

pnpm run dev
```

To run the app backend locally:

```
bazel run //services/thesisTool/cmd/e2e
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

The visual regression tests (storybook + loki) currently have to be run using the script `./scripts/visual_regression.sh`.

> Note: The loki tests are flaky, because it does not wait for the dialog animation [#190](https://github.com/TimoSto/ThesorTeX/issues/190)

## Deployment

To update aws infrastructure

```
cd terraform
./scripts/deploy.sh
```

To upload build-artifacts as zips

```
./scripts/s3_upload_atrifacts.sh all
```

To only deploy one target, replace `all` with `thesisTool`, `thesisTemplate` or `cvTemplate`.

## Used technologies

- Go
- Pnpm
- Typescript
- Vue 3 + Vuetify 3
- Playwright
- Terraform
- Bazel
- Storybook + Loki

## Project structure

This project uses the monorepo approach.

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
