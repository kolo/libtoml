package libtoml

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

var example []byte

func init() {
	var err error
	example, err = ioutil.ReadFile("example.toml")
	if err != nil {
		fmt.Println("can't read example.toml")
		os.Exit(1)
	}
}

func Test_parseExample(t *testing.T) {
	Parse(string(example))
}
