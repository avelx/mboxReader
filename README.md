** Intention: 
Read large *.mbox file and split it into convenient format for later processing by other tools;

MBOX file format:
https://en.wikipedia.org/wiki/Mbox

Ideas:
Drop Hmtl tags =>
    https://pkg.go.dev/golang.org/x/net/html

New PR standards:
feat for a new feature for the user, not a new feature for build script. Such commit will trigger a release bumping a MINOR version.
fix for a bug fix for the user, not a fix to a build script. Such commit will trigger a release bumping a PATCH version.
perf for performance improvements. Such commit will trigger a release bumping a PATCH version.
docs for changes to the documentation.
style for formatting changes, missing semicolons, etc.
refactor for refactoring production code, e.g. renaming a variable.
test for adding missing tests, refactoring tests; no production code change.
build for updating build configuration, development tools or other changes irrelevant to the user.