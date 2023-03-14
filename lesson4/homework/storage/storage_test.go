package storage

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_LocalDirAndFile(t *testing.T) {

	// prepare tempDir
	td, err := os.MkdirTemp(os.TempDir(), "test-dir-*")
	assert.NoError(t, err)
	defer os.Remove(td)

	d1 := filepath.Join(td, "dir1")
	err = os.Mkdir(d1, os.ModePerm)
	assert.NoError(t, err)
	defer os.Remove(d1)

	f1 := filepath.Join(td, "test1.txt")
	err = os.WriteFile(f1, []byte("hello"), os.ModeTemporary)
	assert.NoError(t, err)
	defer os.Remove(f1)

	f2 := filepath.Join(td, "test2.txt")
	err = os.WriteFile(f2, []byte("hello world"), os.ModeTemporary)
	assert.NoError(t, err)
	defer os.Remove(f2)

	f3 := filepath.Join(d1, "test3.txt")
	err = os.WriteFile(f3, []byte("hello world from dir1"), os.ModeTemporary)
	assert.NoError(t, err)
	defer os.Remove(f3)

	t.Run("ok, dir name", func(t *testing.T) {
		d := NewLocalDir("/path")
		assert.Equal(t, "/path", d.Name())
	})

	t.Run("ok, dir ls", func(t *testing.T) {
		d := NewLocalDir(td)
		dirs, files, err := d.Ls(context.Background())
		assert.NoError(t, err)
		assert.ElementsMatch(t, []Dir{
			&localDir{name: d1},
		}, dirs)
		assert.ElementsMatch(t, []File{
			&localFile{name: f1},
			&localFile{name: f2},
		}, files)
	})

	t.Run("fail, dir ls, context error", func(t *testing.T) {
		d := NewLocalDir(d1)

		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		_, _, err := d.Ls(ctx)
		assert.ErrorContains(t, err, "context canceled")
	})

	t.Run("ok, file name", func(t *testing.T) {
		f := NewLocalFile(f1)
		assert.Equal(t, f1, f.Name())
	})

	t.Run("ok, file stat", func(t *testing.T) {
		f := NewLocalFile(f1)

		size, err := f.Stat(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, int64(5), size)
	})

	t.Run("fail, file not found", func(t *testing.T) {
		f := NewLocalFile("/path/to/nonexistent")

		size, err := f.Stat(context.Background())
		assert.Error(t, err)
		assert.Equal(t, int64(0), size)
		assert.Error(t, err)
	})

	t.Run("fail, context error", func(t *testing.T) {
		f := NewLocalFile(f1)

		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		_, err := f.Stat(ctx)
		assert.ErrorContains(t, err, "context canceled")
	})
}

func Test_DummyDirAndFile(t *testing.T) {
	t.Run("ok, dir name", func(t *testing.T) {
		d := NewDummyDir("/path", 0, nil, nil)
		assert.Equal(t, "/path", d.Name())
	})

	t.Run("ok, dir Ls, no latency", func(t *testing.T) {
		d := NewDummyDir("/path", 0, []Dir{
			&dummyDir{name: "d1"},
			&dummyDir{name: "d2"},
		}, []File{
			&dummyFile{name: "f1"},
			&dummyFile{name: "f2"},
		})

		dirs, files, err := d.Ls(context.Background())
		assert.NoError(t, err)
		assert.ElementsMatch(t, []Dir{
			&dummyDir{name: "d1"},
			&dummyDir{name: "d2"},
		}, dirs)
		assert.ElementsMatch(t, []File{
			&dummyFile{name: "f1"},
			&dummyFile{name: "f2"},
		}, files)
	})

	t.Run("fail, dir Ls, context deadline", func(t *testing.T) {
		d := NewDummyDir("/path", 200*time.Millisecond, []Dir{
			&dummyDir{name: "d1"},
			&dummyDir{name: "d2"},
		}, []File{
			&dummyFile{name: "f1"},
			&dummyFile{name: "f2"},
		})

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
		defer cancel()
		_, _, err := d.Ls(ctx)
		assert.Error(t, err)
		assert.ErrorContains(t, err, "context deadline")
	})

	t.Run("ok, file name", func(t *testing.T) {
		f := NewDummyFile("/path/to/f1", 0)
		assert.Equal(t, "/path/to/f1", f.Name())
	})

	t.Run("ok, file stat", func(t *testing.T) {
		f := NewDummyFile("/path/to/f1", 110)

		size, err := f.Stat(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, int64(110), size)
	})

	t.Run("fail, context error", func(t *testing.T) {
		f := NewDummyFile("/path/to/f1", 0)

		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		_, err := f.Stat(ctx)
		assert.ErrorContains(t, err, "context canceled")
	})
}
