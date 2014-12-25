package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Your Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X Yosemite.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
}
