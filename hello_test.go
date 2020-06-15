package pkg

import (
	"howtogenerategodoc/hello"
	"testing"
)

func TestHello(t *testing.T) {
	hello.HelloA()
}

func TestBye(t *testing.T) {
	hello.ByeA()
}
