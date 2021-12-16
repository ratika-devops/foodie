//go:generate pkger

package main

import (
	"foodies/gosrc/cmd"
	"foodies/gosrc/docs"
)

var (
	version string
)

func main() {
	docs.SwaggerInfo.Version = version
	cmd.Execute()
}
