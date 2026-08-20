# Chapter 14. Code quality assurance: automated testing
## Automated tests find your bugs before someone else does
- An **automated test** is a separate program that executes components of your main program, and verifies they behave as expected.

## Writing tests
- Go includes a `testing` package that you can use to write automated tests for your code, and a `go test` command that you can use to run those tests.
- Go functions, but it needs to follow certain conventions in order to work with the `go test` tool:
  - You are not required to make your tests part of the same package as the code you are testing, but if you want to access unexported types or functions from the package, you will need to.
  - Tests are required to use a type from the `testing` package, so you will need to import that package at the top of each test file.
  - Test function names should begin with `Test`.
    - The rest of the name can be whatever you want, but it should begin with a capital letter.
  - Test functions should accept a single parameter: a pointer to a `testing.T` value.
  - You can report that a test has failed by calling methods (such as `Error`) on the `testing.T` value. Most methods accept a string with a message explaining the reason the test failed.

## Running tests with the "go test" command
- To run tests, you use the `go test` command.
  - The command takes the import paths of one or more packages.
  - It will find all files in those packages directories whose names end in *_test.go*, and run every function contained in those files whose name starts with `Test`.

## More detailed test failure message w ith the "Errorf" method
- A test function's `testing.T` parameter also has an `Errorf` method you can call.
- `want` variable to hold the expected return value
- `got` to hold the actual return value.
- If `got` isn't equal to `want`, call `Errorf` with the verb `%#v` so the slice is printed the same way it would appear in Go code, the `got` value and the `want` value.

## Test "helper" functions
- The `go test` command only uses functions whose names begin with `Test`, so as long as you name your functions anything else, you will be fine.

## Test-driven development
1. Write the test: you write a test for the feature you want, even though it doesn't exist yet. Then you run the test to esnure that it fails.
2. Make it pass: you implement the feature in your main code. Don't worry about whether the code you are writing is sloppy or inefficient; your only goal it to get it working. Then you run the test to ensure that it passes.
3. Refactor your code: now, you are free to refactor the code, to change and improve it, however you please. You have watched the test fail, so you know it will fail again if your app code breaks. You have watched the test pass, so you know it will continue passing as long as your code is working correctly.

## Running specific sets of tests
- `-v` flag, which stands for "verbose", it willl list the name and status of each test function it runs.
  - Normally passing tests are omitted to keep the output "quiet", but in verbose mode, `go test` will list even passing tests.
- `-run` option to limit the set of tests that are run.
  - Following `-run`, you specify part or all of a function name, and only test functions whose name matches what you specify will be run.
