package hellogolang

import "testing"

func TestHello(t *testing.T) {
	result := Hello()
	if result != "hello golang !" {
		t.Fatal("Hello function return Fail")
	}
}
