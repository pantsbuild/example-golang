// Copyright 2021 Pants project contributors.
// Licensed under the Apache License, Version 2.0 (see LICENSE).

package main

import (
	"fmt"
	"github.com/pantsbuild/example-golang/pkg/greeter"
	"github.com/spf13/pflag"
	"os"
)

const version string = "0.1.0"

func main() {
	versionOpt := pflag.BoolP("version", "V", false, "imprimir la versión y salir")
	pflag.Parse()

	if *versionOpt {
		fmt.Println(version)
		os.Exit(0)
	}

	fmt.Println(greeter.GreetSpanish("Pantsbuild"))
}
