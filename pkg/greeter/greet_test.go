// Copyright 2021 Pants project contributors.
// Licensed under the Apache License, Version 2.0 (see LICENSE).

package greeter_test // That is, an "external test"

import (
	"github.com/pantsbuild/example-golang/pkg/greeter"
	"strings"
	"testing"
)

func TestEnglish(t *testing.T) {
	result := greeter.GreetEnglish("testing")
	if !strings.HasPrefix(result, "Hello testing!") {
		t.Fail()
	}
}

func TestSpanish(t *testing.T) {
	result := greeter.GreetSpanish("testing")
	if !strings.HasPrefix(result, "¡Hola testing!") {
		t.Fail()
	}
}
