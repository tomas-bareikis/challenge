package main

import (
	"os"
	"strings"
	"testing"

	"github.com/bitfield/script"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	inFile     string
	expectFile string
}

func Test(t *testing.T) {
	tests := map[string]testCase{}

	dirs, _ := os.ReadDir("./tests")
	for _, d := range dirs {
		name := d.Name()
		tests[name] = testCase{
			inFile:     "tests/" + name + "/in.txt",
			expectFile: "tests/" + name + "/expected.txt",
		}
	}

	for name, tc := range tests {
		t.Run(name, doTest(&tc))
	}
}

func doTest(tc *testCase) func(*testing.T) {
	return func(t *testing.T) {
		expectedOutput, err := script.File(tc.expectFile).String()
		assert.NoError(t, err)

		output, err := run(tc.inFile)
		assert.NoError(t, err)

		assert.Equal(
			t,
			strings.TrimSpace(expectedOutput),
			strings.TrimSpace(output),
		)
	}
}

func TestBig(t *testing.T) {
	lang := os.Getenv("LANGUAGE")

	team1 := "HEHE1 50 "
	team2 := "HEHE2 51 "
	team3 := "HEHE3 52 "

	steps := strings.Repeat("1000 ", 1000000)

	in := ""
	in += team1 + steps + "\n"
	in += team2 + steps + "\n"
	in += team3 + steps + "\n"

	out, err := script.Echo(in).Exec("./run-" + lang + ".sh").String()
	assert.NoError(t, err)

	expected := `HEHE1 1 500000.00
HEHE2 1 510000.00
HEHE3 1 520000.00`

	assert.Equal(
		t,
		strings.TrimSpace(expected),
		strings.TrimSpace(out),
	)
}

func run(inFile string) (string, error) {
	lang := os.Getenv("LANGUAGE")
	return script.File(inFile).Exec("./run-" + lang + ".sh").String()
}
