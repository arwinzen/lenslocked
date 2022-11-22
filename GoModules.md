# Go Modules

## 1. Dependency Management

Make sure other developers can build your code using a similar version

Go uses SIV (Semantic Import Versioning) which is based on Semantic Versioning
- See <semver.org>

Highlights:
- Composed of 3 numbers. eg `v1.12.4`
- Major, minor, patch
- Major bumps == breaking change. Manual upgrade is necessary. 
- Other bumps can be automated by tooling if needed. 

Special case:
- v0 allows for any breaking changes. It is a special version
- `go get github.com/joncalhoun/somelib` will NOT get the most recent version. 
    - We need to run `go get github.com/joncalhoun/somelib/v4` or similar.
    - Tooling ard this isn't always perfect. Will hopefully improve over time. 

# 2. Working outside of `GOPATH`

All GO code used to live inside a single directory on your computer - the 
`GOPATH`

GO Modules allow us to run code from anywhere, as long as we initialize a module. 

## Setting up our module 

```bash
go mod init github.com/arwinzen/lenslocked
#
#
```

This creates a `go.mod` file.

From this point onwards we won't need to think too much about modules. If we `go get` a library, the tooling will update our `go.mod` file for us. 


