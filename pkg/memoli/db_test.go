package memoli

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

func TestBucket_Get(t *testing.T) {
	b := &Bucket{mm: map[string]interface{}{
		"key": "value",
	}}

	assert.Equal(t, "value", b.Get("key"))
}

func TestBucket_Set(t *testing.T) {
	b := &Bucket{mm: make(map[string]interface{})}
	b.Set("key", "value")

	assert.Equal(t, "value", b.mm["key"])
}

func TestBucket_Close(t *testing.T) {
	b := &Bucket{snapshotWindow: time.NewTicker(time.Second), closec: make(chan struct{}, 1)}
	b.Close()
	for {
		select {
		case <-time.Tick(time.Second):
			t.Fatal("Bucket.Close() did not close the channel")
		case <-b.closec:
			return
		}
	}
}

func TestBucket_pathToFile(t *testing.T) {
	b := &Bucket{
		snapshotPath: "/tmp/memoli",
		bucketName:   "zeyno",
	}

	assert.Equal(t, "/tmp/memoli/zeyno.mdb", b.pathToFile())
}

func Test_writeTo(t *testing.T) {
	type args struct {
		path string
		mm   map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "writes file without given an error",
			args: args{
				path: "/tmp/zeyno.mdb",
				mm: map[string]interface{}{
					"key": "value",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := writeTo(tt.args.path, tt.args.mm); (err != nil) != tt.wantErr {
				t.Errorf("writeTo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBucket_rsync(t *testing.T) {
	type fields struct {
		bucketName     string
		snapshotWindow *time.Ticker
		snapshotPath   string
		mm             map[string]interface{}
		closec         chan struct{}
		RWMutex        sync.RWMutex
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "rsync without error",
            fields: fields{
                bucketName:     "zeyno",
                snapshotWindow: time.NewTicker(time.Second),
                snapshotPath:   "/tmp/memo",
                mm:             map[string]interface{}{},
                closec:         make(chan struct{}, 1),
                RWMutex:        sync.RWMutex{},
            },
            wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Bucket{
				bucketName:     tt.fields.bucketName,
				snapshotPath:   tt.fields.snapshotPath,
				mm:             tt.fields.mm,
			}
			err := writeTo(b.pathToFile(), map[string]interface{}{"key": "value"})
			if err != nil {
				t.Fatal(err)
			}
			err = b.rsync()

			assert.Equal(t, tt.wantErr, err != nil, "Bucket.rsync() error = %v, wantErr %v", err, tt.wantErr)
			assert.Equal(t, "value", b.mm["key"])
		})
	}
}

func TestBucket_wsync(t *testing.T) {
	done := make(chan struct{}, 1)
	writeToFunc = func(pathTo string, mm map[string]interface{}) error {
		done <- struct{}{}
		return nil
	}
	b := &Bucket{snapshotWindow: time.NewTicker(time.Second), closec: make(chan struct{}, 1)}
	go b.wsync()
	for {
		select {
		case <-time.Tick(time.Second*2):
			t.Fatal("Bucket.Close() did not close the channel")
		case <-done:
			return
		}
	}
}