# ThesorTeX
Tool for handling bibliography in a LaTeX project

## Downloads
The local application can be downloaded on the following links:
- [Windows](https://thesortex-artifacts.s3.eu-central-1.amazonaws.com/latest/windows/ThesorTeX.zip)
- [Linux](https://thesortex-artifacts.s3.eu-central-1.amazonaws.com/latest/linux/ThesorTeX.zip)
- [MacOS](https://thesortex-artifacts.s3.eu-central-1.amazonaws.com/latest/mac/ThesorTeX.zip)
- [MacOS (Apple Silicon)](https://thesortex-artifacts.s3.eu-central-1.amazonaws.com/latest/mac_silicon/ThesorTeX.zip)

## Build mechanics

### Production
The frontend is build using pnpm, the backend is build using Bazel. Run `./build.sh` to trigger the build.

### Development
To run the frontend on a watcher
```
cd services/app/frontend
pnpm run dev
```
To run the backend
```
./scripts/dev.sh
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
cd terraform
./scripts/s3_upload_atrifacts.sh
```

## Used technologies
- Go
- pnpm
- Playwright
- terraform
- bazel

## Project structure

### services
Here the separation of the different services, that are independently runnable.

#### website
The service for hosting the website.

##### cmd
The lambda executable

##### internal
The backend for the website using the global go.mod

##### frontend
The frontend for the website using the global pnpm-workspace

#### app
The service for managing the bibliography projects.

##### cmd
The local executable

##### internal
The backend for the app using the global go.mod

##### frontend
The frontend for the app using the global pnpm-workspace

### pkg
Go packages used among multiple services

### packages
Pnpm packages used among multiple services

### testing
mocks and playwright tests

### terraform
terraform modules

### scripts
build scripts

### bazel
global bazel files
