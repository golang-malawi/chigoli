package fsutil

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"time"
)

// UnlinkOn removes files from a directory condition specified by condFunc and returns the number of files removed
func UnlinkOn(dirname string, condFunc func(f fs.FileInfo) bool) (int64, error) {
	if dirname == "" {
		return 0, fmt.Errorf("dirname must not be empty")
	}

	entries, err := os.ReadDir(dirname)
	if err != nil {
		return 0, fmt.Errorf("failed to readdir got: %w", err)
	}

	filesToRemove := make([]os.FileInfo, 0)
	for _, entry := range entries {
		efi, err := entry.Info()
		if err != nil {
			return 0, fmt.Errorf("failed to unlink files at %s", dirname)
		}
		if condFunc(efi) {
			filesToRemove = append(filesToRemove, efi)
		}
	}

	var n int64
	for _, fi := range filesToRemove {
		err = os.Remove(filepath.Join(dirname, fi.Name()))
		if err != nil {
			return 0, err
		}
		n++
	}
	return n, nil
}

// UnlinkIfOlderThan removes files from the directory that are older than the given time
func UnlinkOlderThan(filename string, referenceTime time.Time) (int64, error) {
	if referenceTime.IsZero() {
		return 0, fmt.Errorf("referenceTime must be specified")
	}

	return UnlinkOn(filename, func(f fs.FileInfo) bool {
		return f.ModTime().Before(referenceTime)
	})
}
