package main

import "embed"

//go:embed data/*.txt
var dataFS embed.FS

func main() {}
