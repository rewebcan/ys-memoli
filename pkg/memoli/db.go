package memoli

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"os"
	"path"
	"strings"
	"sync"
	"time"
)

var MalformedSnapshotError = errors.New("malformed snapshot")

var _ext = "mdb"

var writeToFunc = writeTo

// Bucket represents a collection of items stored in a in-memory database.
type Bucket struct {
	bucketName     string
	snapshotWindow *time.Ticker
	snapshotPath   string
	mm             map[string]interface{}
	closec         chan struct{}
	sync.RWMutex
}

// NewBucket creates a new in-memory key-value database
func NewBucket(bucketName string, opts ...Option) (*Bucket, error) {
	b := &Bucket{
		bucketName:     bucketName,
		mm:             make(map[string]interface{}),
		closec:         make(chan struct{}, 1),
		snapshotPath:   "/tmp",
		snapshotWindow: time.NewTicker(time.Minute),
	}

	for _, opt := range opts {
		opt(b)
	}

	if err := b.rsync(); err != nil {
		return nil, err
	}

	go b.wsync()

	return b, nil
}

// Get returns the value for the given key
func (b *Bucket) Get(key string) interface{} {
	b.RLock()
	defer b.RUnlock()

	return b.mm[key]
}

// Set sets the value for the given key
func (b *Bucket) Set(key string, value interface{}) {
	b.Lock()
	defer b.Unlock()

	b.mm[key] = value
}

// Close closes the bucket
func (b *Bucket) Close() {
	b.snapshotWindow.Stop()
	b.closec <- struct{}{}
}

func (b *Bucket) wsync() {
	for {
		select {
		case <-b.snapshotWindow.C:
			b.Lock()
			if err := writeToFunc(b.pathToFile(), b.mm); err != nil {
				log.Err(err).Send()
			}
			b.Unlock()
		case <-b.closec:
			b.Lock()
			if err := writeToFunc(b.pathToFile(), b.mm); err != nil {
				log.Err(err).Send()
			}
			b.Unlock()
			return
		}
	}
}

func (b *Bucket) rsync() error {
	f, err := os.Open(b.pathToFile())
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			return fmt.Errorf("%w: could not read from line", MalformedSnapshotError)
		}

		b.mm[parts[0]] = parts[1]
	}

	return nil
}

func (b *Bucket) pathToFile() string {
	return fmt.Sprintf("%s/%s.%s", b.snapshotPath, b.bucketName, _ext)
}

func writeTo(pathTo string, mm map[string]interface{}) error {
	err := os.MkdirAll(path.Dir(pathTo), os.ModePerm)
	if err != nil {
		return err
	}
	f, err := os.Create(pathTo)
	if err != nil {
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)

	for key, val := range mm {
		_, err = w.Write([]byte(fmt.Sprintf("%s=%v\n", key, val)))
		if err != nil {
			return err
		}
	}

	return w.Flush()
}
