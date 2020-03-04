# gohtpass
Explore the possibilities of a golang code generation

## Golang code generation example
The purpose of this project is to explore a golang code generation.
The result is baking a git ignored .htpasswd file inside the binary file.
This is pseudo security don't do that for a real production project.

Follow PassFile.md to create an apr1 .htpasswd file and then use
`go generate` to generate `secret.go` file.

you can then run `go run *.go` to execute the example code.