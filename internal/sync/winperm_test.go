package sync

import (
	"runtime"
	"testing"
)

// skipPermsOnWindows skips tests that only assert Unix permission bits.
// Windows has no 0600/0700 modes (Go reports 0666/0777); access to files in
// the user profile is controlled by its ACLs instead.
func skipPermsOnWindows(t *testing.T) {
	t.Helper()
	if runtime.GOOS == "windows" {
		t.Skip("Windows has no Unix permission bits; access is controlled by profile ACLs")
	}
}
