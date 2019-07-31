// !build

package business

//go:generate gogen-enum -input ./addressKind.yaml -package business -output ./addressKind.go
//go:generate gofmt -w addressKind.go
//go:generate golangci-lint run addressKind.go
