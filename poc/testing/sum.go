package testing

func sumTwoDigits(a, b int) int {
	return a + b
}

// run test command : go test
// get more test info : go test -v
// run single test file : go test -v -run testfilename
// get coverage : gotest -cover
// get more coverage info : go test -coverprofile=cover
// get coverage html file : go tool cover -html=cover -o cover.html
