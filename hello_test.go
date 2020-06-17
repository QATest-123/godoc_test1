//This is hello package1
//This is hello package2
package pkg

import (
	"github.com/QATest-123/godoc_test1/hello"
	"testing"
)

func TestHello(t *testing.T) {
	hello.HelloA()
}

func TestBye(t *testing.T) {
	hello.ByeA()
}
