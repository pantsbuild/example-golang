// Copyright 2021 Pants project contributors.
// Licensed under the Apache License, Version 2.0 (see LICENSE).

package embed_example

import (
	"os"
	"testing"
)

// TODO: enable after fixing Pants when lib code has embeds but tests don't.
// func TestResourceEmbedding(t *testing.T) {
//   if embeddedHello != "I'm an embedded resource!" {
//     t.Fatalf("message mismatch: want=%s; got=%s", "I'm an embedded resource!", embeddedHello)
//   }
// }

func TestTestDataFolder(t *testing.T) {
	_, err := os.Stat("testdata/f.txt")
	if err != nil {
		t.Fatalf("Could not stat pkg/embed_example/testdata/f.txt: %v", err)
	}
}
