# simple-app-go

## Requirements
- [Go](https://golang.org/)
- [Yarn](https://yarnpkg.com/)

## Dependencies
#### Go
Go dependencies are vendored as git submodules and easily installed as follows
```bash
git clone https://github.com/kelleydv/simple-app-go.git
cd bar
git submodule update --init --recursive
```
Or, simply
```bash
git clone --recursive https://github.com/kelleydv/simple-app-go.git
```
#### Front-end
```bash
yarn install
```

## Run locally
After installing dependencies
```bash
cd simple-app-go
go build
./simple-app-go
```
