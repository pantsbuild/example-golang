# example-golang

An example repository to demonstrate Pants's experimental Golang support.

See the [Golang blog post](https://blog.pantsbuild.org/golang-support-pants-28/) for some unique
benefits Pants brings to Golang repositories, and see
[pantsbuild.org/docs/go-overview](https://www.pantsbuild.org/docs/go-overview) for more detailed
documentation.

This is only one possible way of laying out your project with Pants. See 
[pantsbuild.org/docs/source-roots#examples](https://www.pantsbuild.org/docs/source-roots#examples) 
for some other example layouts.

# Running Pants

You run Pants goals using the `pants` launcher binary, which will bootstrap the
version of Pants configured for this repo if necessary.

See [here](https://www.pantsbuild.org/docs/installation) for how to install the `pants` binary.

# Goals

Pants commands are called _goals_. You can get a list of goals with

```
pants help goals
```

Most goals take arguments to run on. To run on a single directory, use the directory name with 
`:` at the end. To recursively run on a directory and all its subdirectories, add `::` to the 
end.

For example:

```
pants lint cmd: internal::
```

You can run on all changed files:

```
pants --changed-since=HEAD lint
```

You can run on all changed files, and any of their "dependees":

```
pants --changed-since=HEAD --changed-dependees=transitive test
```

# Example Goals

Try these out in this repo!

## Run Gofmt

```
pants fmt ::  # Format all packages.
pants fmt cmd/greeter_en:  # Format only this package.
pants lint pkg::  # Check that all packages under `pkg` are formatted.
```

## Check compilation

```
pants check ::  # Compile all packages.
pants check cmd/greeter_en:  # Compile only this package and its transitive dependencies.
```

## Run tests

```
pants test ::  # Run all tests in the repository.
pants test pkg/uuid:  # Run all the tests in this package.
pants test pkg/uuid: -- -run TestGenerateUuid  # Run just this one test.
```

## Create a binary file

Writes the result to the `dist/` folder.

```
pants package cmd/greeter_en:
pants package cmd::  # Create all binaries.
```

## Run a binary

```
pants run cmd/greeter_en:
pants run cmd/greeter_es: -- --help
```

## Determine dependencies

```
pants dependencies cmd/greeter_en:
pants dependencies --transitive cmd/greeter_en:
```

## Determine dependees

That is, find what code depends on a particular package(s).

```
pants dependees pkg/uuid:
pants dependees --transitive pkg/uuid: 
```

## Count lines of code

```
pants count-loc '**/*'
```
