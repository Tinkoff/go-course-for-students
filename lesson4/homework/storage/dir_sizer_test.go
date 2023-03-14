package storage

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_DirSizer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("ok, file storage", func(t *testing.T) {
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

		sizer := NewSizer()

		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		result, err := sizer.Size(ctx, NewLocalDir(td))
		assert.NoError(t, err)
		assert.Equal(t, int64(3), result.Count)
		assert.Equal(t, int64(37), result.Size)
	})

	t.Run("fail, file storage, dir not exist", func(t *testing.T) {
		sizer := NewSizer()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		_, err := sizer.Size(ctx, NewLocalDir("/path/to/nonexistent"))
		assert.Error(t, err)
	})

	t.Run("fail, file does not exist", func(t *testing.T) {
		sizer := NewSizer()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		f := NewMockFile(ctrl)
		f.EXPECT().Stat(gomock.Any()).Return(int64(0), errors.New("file does not exist"))

		d := NewMockDir(ctrl)
		d.EXPECT().Ls(gomock.Any()).Return(nil, []File{f}, nil)

		_, err := sizer.Size(ctx, d)
		assert.ErrorContains(t, err, "file does not exist")
	})

	t.Run("ok, dummy storage", func(t *testing.T) {
		sizer := NewSizer()

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		result, err := sizer.Size(ctx, getDummySet())
		assert.NoError(t, err)
		assert.Equal(t, int64(14), result.Count)
		assert.Equal(t, int64(37254162), result.Size)
	})
}

func Test_DirSizerAsync(t *testing.T) {
	sizer := NewSizer()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	result, err := sizer.Size(ctx, getDummySet())
	assert.NoError(t, err)
	assert.Equal(t, int64(14), result.Count)
	assert.Equal(t, int64(37254162), result.Size)
}

func getDummySet() Dir {
	return NewDummyDir("/root", 10*time.Millisecond, []Dir{
		NewDummyDir("/root/foo", 40*time.Millisecond, []Dir{
			NewDummyDir("/root/foo/bar", 359*time.Millisecond, nil, []File{
				NewDummyFile("/root/foo/bar/f9.txt", 6353),
				NewDummyFile("/root/foo/bar/f10.txt", 235621),
			}),
			NewDummyDir("/root/foo/baz", 590*time.Millisecond, nil, []File{
				NewDummyFile("/root/foo/baz/f11.txt", 76504),
				NewDummyFile("/root/foo/baz/f12.txt", 252446),
			}),
			NewDummyDir("/root/foo/baz", 690*time.Millisecond, nil, []File{
				NewDummyFile("/root/foo/baz/f13.txt", 3455),
				NewDummyFile("/root/foo/baz/f14.txt", 121223),
			}),
		}, []File{
			NewDummyFile("/root/foo/f5.txt", 735635),
			NewDummyFile("/root/foo/f6.txt", 372650),
			NewDummyFile("/root/foo/f7.txt", 467),
			NewDummyFile("/root/foo/f8.txt", 35415264),
		}),
	}, []File{
		NewDummyFile("/root/f1.txt", 1249),
		NewDummyFile("/root/f2.txt", 3523),
		NewDummyFile("/root/f3.txt", 8542),
		NewDummyFile("/root/f4.txt", 21230),
	})
}
