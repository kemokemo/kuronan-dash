package music

import (
	"os"
	"testing"
)

func TestMain(t *testing.M) {
	exitCode := t.Run()
	os.Exit(exitCode)
}
