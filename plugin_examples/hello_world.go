package main

import (
    "strings"
    "text/template"
)

// go build -buildmode=plugin -o hello_world.so ./hello_world.go
var FuncMap = template.FuncMap{
    "Hello": func(name string) string {
        return "Hello, " + name + "!"
    },
}
