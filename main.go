package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func main() {
	color.Green("Hello from project-gopm!")
	printInfo()
}

func printInfo() {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}
	fmt.Printf("Running on host: %s\n", hostname)
	fmt.Printf("Go package manager: gopm\n")
	fmt.Printf("Dependency file: .gopmfile\n")
	fmt.Printf("Install command: gopm get\n")
	fmt.Printf("Gopm uses a .gopmfile (INI format) to declare dependencies.\n")
	fmt.Printf("Lock file: not supported (versions pinned via @ tags in .gopmfile)\n")
	fmt.Printf("Vendor dir: vendor/ (populated by gopm get)\n")
}