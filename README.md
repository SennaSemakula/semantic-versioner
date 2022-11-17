# semantic-versioner
## Why?
Using go:generate and go:embed to provide safety checking of build versions over ldflags because latter does not provide a safety system to check for input validation. For example there could be a case where we forget to pass in the version tag but it won’t fail during compile time. This approach will fail fast as it always expects an artifact to be present.

Build version is available with version const at runtime.

There is a script that generates an artifact version.txt which is ephemeral and should only be generated during the CD pipeline on the fly. It should never be added to upstream version control.

## Running
1. `make gen`

2. `make run`

## Running tests
1. `make test`