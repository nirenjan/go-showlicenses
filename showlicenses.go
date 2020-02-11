// showlicenses is a utility that wraps the licensecheck library
// and returns a list of the matching licenses in the given file

// SPDX-License-Identifier: MIT
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	lc "github.com/google/licensecheck"
)

func printLicenses(file string) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	opts := lc.Options{}

	cov, valid := lc.Cover(data, opts)
	if !valid {
		fmt.Fprintf(os.Stderr, "%v: No detected licenses\n", file)
		return
	}

	licenses := make(map[string]struct{})
	exists := struct{}{}
	for _, m := range cov.Match {
		licenses[m.Name] = exists
	}
	output := make([]string, len(licenses))
	i := 0
	for l := range licenses {
		output[i] = l
		i++
	}
	fmt.Printf("%v: %v\n", file, strings.Join(output, ", "))
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: showlicenses <list of files containing license text>")
		os.Exit(1)
	}

	for _, f := range os.Args[1:] {
		printLicenses(f)
	}
}
