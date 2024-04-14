# Contributing

Thanks for your interest in contributing!

## Contributing issues

Please feel free to open an issue if you think something is working incorrectly,
if you have a feature request, or if you just have any questions.  No templates
are in place, all I ask is that you provide relevant information if you believe
something is working incorrectly so we can sort it out quickly.

## Contributing code

All contributions should have an associated issue.  If you are at all unsure
about how to solve the issue, please ask!  I'd rather chat about possible
solutions than have someone spend hours on a PR that requires a lot of major
changes.

Test coverage is important.  If you end up with a small (<1%) drop, I'm happy to
help cover the gap, but generally any new features or changes should have some
tests as well.  If you're not sure how to test something, feel free to ask!

Running `fmt` can be done with `go fmt`. Tests can be run with `go test` and `go test -cover`.

Doing all these before submitting the PR can help ensure the code is properly formatted, the tests pass, and the overall code coverage is maintained.

The name of the PR and branch names aren't very important, but commit messages should follow [conventional commits](https://www.conventionalcommits.org/en/v1.0.0/).