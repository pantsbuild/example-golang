# Copyright 2021 Pants project contributors.
# Licensed under the Apache License, Version 2.0 (see LICENSE).

go_package()

# This target allows us to use `pants run` and `pants package` on this `main` Go package.
#
# You can optionally set the field `output_path="greeter_es"`, for example, for the binary's name
# to be different when running `pants package`.
go_binary(
    name="bin",
)
