// Package utils
// @author Daud Valentino
package util

import (
	"strings"
	"testing"
)

func TestDumpToString(t *testing.T) {
	t.Parallel()
	if result := DumpToString("test"); result == "test" {
		t.Logf("expected: %v, got: %v", "test", result)
	} else {
		t.Fatalf("expected: %v, got: %v", "test", result)
	}

	if result := DumpToString(1); strings.Trim(result, "\n") == "1" {
		t.Logf("expected: %v, got: %s", "1", result)
	} else {
		t.Fatalf("expected: %v, got: %v", "1", result)
	}
}

func TestDebugPrint(t *testing.T) {
	DebugPrint("test print")
}
