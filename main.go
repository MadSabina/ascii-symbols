package main

import (
	"fmt"
	"os"
	"strings"

	"ascii/ascii"
)

func main() {
	art := ascii.ASCII{}
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Only one argument required.")
		return
	}
	art.Args = args[0]
	if !checkbanner(art.Args) {
		err := ascii.Ascii(art)
		if err != nil {
			fmt.Println(err)
		}
		return
	}

}

func checkbanner(s string) bool {
	return strings.HasPrefix(s, "--reverse=")
}
