package fsutil_test

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/golang-malawi/chigoli/fsutil"
)

func TestUnlinkOnWithEmptyDir(t *testing.T) {
	_, err := fsutil.UnlinkOn("", func(f fs.FileInfo) bool {
		return true
	})

	if err == nil {
		t.Errorf("expected there to be an error")
	}
	if err.Error() != "dirname must not be empty" {
		t.Errorf("expected error message to be 'dirname must not be empty'")
	}
}

func TestUnlinkOnWithNonExistentDir(t *testing.T) {
	_, err := fsutil.UnlinkOn("@!-not-existent", func(f fs.FileInfo) bool {
		return true
	})

	if err == nil {
		t.Errorf("expected there to be an error")
	}

	if !strings.HasPrefix(err.Error(), "failed to readdir") {
		t.Errorf("got unexpected error message '%s'", err)
	}
}

func TestUnlinkOn(t *testing.T) {
	var want int64 = 10
	dir, err := makeTestDir(10)
	if err != nil {
		t.Errorf("failed to run test %v", err)
	}
	have, err := fsutil.UnlinkOn(dir, func(f fs.FileInfo) bool {
		return true
	})

	if err != nil {
		t.Errorf("failed to run test %v", err)
	}
	if want != have {
		t.Errorf("want=%d have=%d", want, have)
	}
}

func TestUnlinkOlderWithZeroTime(t *testing.T) {
	dir, err := makeTestDir(10)
	if err != nil {
		t.Errorf("failed to run test %v", err)
	}
	_, err = fsutil.UnlinkOlderThan(dir, time.Time{})
	if err == nil {
		t.Errorf("failed to run test expected non-nil error %v", err)
	}
	if err.Error() != "referenceTime must be specified" {
		t.Errorf("expected error message to be 'referenceTime must be specified'`")
	}
}

func TestUnlinkOlder(t *testing.T) {
	var want int64 = 10
	dir, err := makeTestDir(10)
	if err != nil {
		t.Errorf("failed to run test %v", err)
	}
	have, err := fsutil.UnlinkOlderThan(dir, time.Now().Add(1*time.Minute))

	if err != nil {
		t.Errorf("failed to run test %v", err)
	}
	if want != have {
		t.Errorf("want=%d have=%d", want, have)
	}
}

func makeTestDir(n int) (string, error) {
	name, err := os.MkdirTemp("", "test-directory")
	if err != nil {
		panic(err)
	}
	for i := 0; i < n; i++ {
		os.WriteFile(filepath.Join(name, fmt.Sprintf("%d.ext", i)), []byte(`test data`), 0755)
	}

	return filepath.Join(name), nil
}
