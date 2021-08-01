package main

import (
	"github.com/frieser/openbanking/cmd"
	"github.com/frieser/openbanking/console"
)

func main() {
	console.BigText("Openbanking cli")
	cmd.Execute()
}
