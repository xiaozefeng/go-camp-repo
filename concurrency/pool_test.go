package concurrency

import (
	"os"
	"testing"
)

func TestLog(t *testing.T) {
	Log(os.Stdout, "path", "/search")
	Log(os.Stdout, "path", "/search")
	Log(os.Stdout, "path", "/search")
	Log(os.Stdout, "path", "/search")
}
