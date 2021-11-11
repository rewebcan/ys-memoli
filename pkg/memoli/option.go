package memoli

import (
	"strings"
	"time"
)

type Option func(*Bucket)

// SnapshotWindow defines a duration frequency of taking snapshot
func SnapshotWindow(duration time.Duration) Option {
	return func(b *Bucket) {
		b.snapshotWindow = time.NewTicker(duration)
	}
}

// SnapshotPath defines a path to store snapshots
func SnapshotPath(path string) Option {
	return func(b *Bucket) {
		b.snapshotPath = strings.TrimRight(path, "/")
	}
}
