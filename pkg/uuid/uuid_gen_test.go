// Copyright 2021 Pants project contributors.
// Licensed under the Apache License, Version 2.0 (see LICENSE).

package uuid

import "testing"

func TestGenerateUuid(t *testing.T) {
	uuid1 := Generate()
	uuid2 := Generate()
	if uuid1 == uuid2 {
		t.Fail()
	}
}
