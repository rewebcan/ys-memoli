package memoli

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSnapshotPath(t *testing.T) {
	testTmp := "/tmp/memoli_test"
	b := &Bucket{}
	SnapshotPath(testTmp)(b)

	assert.Equal(t, testTmp, b.snapshotPath)
}

func TestSnapshotWindow(t *testing.T) {
	testWindow := time.Second
	b := &Bucket{}
	SnapshotWindow(testWindow)(b)

	assert.NotNil(t, b.snapshotWindow)
}
