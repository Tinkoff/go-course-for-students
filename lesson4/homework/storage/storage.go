package storage

import (
	"context"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"time"
)

//go:generate mockgen -source storage.go -package storage -destination storage_mock.go

// File represent a file object
type File interface {
	// Name return a fully qualified file name
	Name() string
	// Stat returns a size of file or error
	Stat(ctx context.Context) (int64, error)
}

// Dir represent a dir object
type Dir interface {
	// Name return a fully qualified dir name
	Name() string
	// Ls return a set of Dir and a set of File or error if happened
	Ls(ctx context.Context) ([]Dir, []File, error)
}

type localDir struct {
	name string
}

func NewLocalDir(root string) Dir {
	return &localDir{name: root}
}

func (d *localDir) Name() string {
	return d.name
}

func (d *localDir) Ls(ctx context.Context) (dirs []Dir, files []File, err error) {
	defer func() {
		if ctxErr := ctx.Err(); ctxErr != nil {
			err = ctxErr
		}
	}()
	entry, err := os.ReadDir(d.name)
	if err != nil {
		return
	}

	for _, e := range entry {
		if e.IsDir() {
			dirs = append(dirs, NewLocalDir(filepath.Join(d.name, e.Name())))
		} else {
			files = append(files, NewLocalFile(filepath.Join(d.name, e.Name())))
		}
	}

	return dirs, files, nil
}

type localFile struct {
	name string
}

func NewLocalFile(name string) File {
	return &localFile{name: name}
}

func (f *localFile) Name() string {
	return f.name
}

func (f *localFile) Stat(ctx context.Context) (size int64, err error) {
	defer func() {
		if ctxErr := ctx.Err(); ctxErr != nil {
			err = ctxErr
		}
	}()

	var info fs.FileInfo
	info, err = os.Stat(f.name)
	if err != nil {
		return
	}
	if info.IsDir() {
		err = fmt.Errorf("%s is a directory, not a file", f.name)
		return
	}

	size = info.Size()
	return
}

// dummyDir is a fake dir with custom latency delay.
// Latency can be used for make a listing latency for testing Context
type dummyDir struct {
	name string

	// latency can be used for make a listing latency for testing Context
	latency time.Duration
	// files  denote a set of files for Ls() result
	files []File
	// dirs  denote a set of dirs for Ls() result
	dirs []Dir
}

func NewDummyDir(root string, lat time.Duration, dirs []Dir, files []File) Dir {
	return &dummyDir{
		name: root, latency: lat, dirs: dirs, files: files,
	}
}

func (d *dummyDir) Name() string {
	return d.name
}

func (d *dummyDir) Ls(ctx context.Context) (dirs []Dir, files []File, err error) {
	defer func() {
		if ctxErr := ctx.Err(); ctxErr != nil {
			err = ctxErr
		}
	}()

	time.Sleep(d.latency)

	dirs = d.dirs
	files = d.files

	return
}

type dummyFile struct {
	name string
	size int64
}

func NewDummyFile(name string, size int64) File {
	return &dummyFile{
		name: name,
		size: size,
	}
}

func (f *dummyFile) Name() string {
	return f.name
}

func (f *dummyFile) Stat(ctx context.Context) (size int64, err error) {
	defer func() {
		if ctxErr := ctx.Err(); ctxErr != nil {
			err = ctxErr
		}
	}()

	size = f.size
	return
}
