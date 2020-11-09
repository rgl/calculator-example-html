// +build ignore

package main

import (
	"log"
	"net/http"

	"github.com/shurcooL/vfsgen"
)

func main() {
	fs := http.Dir("public")
	err := vfsgen.Generate(fs, vfsgen.Options{
		PackageName:  "main",
		VariableName: "assets",
		BuildTags:    "!dev",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
