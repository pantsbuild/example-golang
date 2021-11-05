// Copyright 2021 Pants project contributors.
// Licensed under the Apache License, Version 2.0 (see LICENSE).

package greeter

import (
	"fmt"
	"github.com/pantsbuild/example-golang/pkg/uuid"
)

func GreetEnglish(name string) string {
	return fmt.Sprintf(
		"Hello %s!\n\nHere's a UUID to brighten your day: %s",
		name,
		uuid.Generate(),
	)
}

func GreetSpanish(name string) string {
	return fmt.Sprintf(
		"¡Hola %s!\n\nEres muy única, así que te regalamos un UUID: %s",
		name,
		uuid.Generate(),
	)
}
