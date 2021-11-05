// Copyright 2021 Pants project contributors.
// Licensed under the Apache License, Version 2.0 (see LICENSE).

package uuid

import "github.com/google/uuid"

func Generate() string {
	return uuid.NewString()
}
