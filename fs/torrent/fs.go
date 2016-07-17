package torrent

import (
	"errors"
	"os"
	"time"

	"github.com/jpillora/cloud-torrent/fs"
	"github.com/spf13/afero"
)

func New() fs.FS {
	return &torrentFS{}
}

type torrentFS struct {
	config struct {
	}
}

func (t *torrentFS) Name() string {
	return "Torrents"
}

func (t *torrentFS) Mode() fs.FSMode {
	return fs.RW
}

func (t *torrentFS) Config() interface{} {
	return &t.config
}

func (t *torrentFS) Create(name string) (afero.File, error) {
	return &file{}, nil
}

func (t *torrentFS) Open(name string) (afero.File, error) {
	return &file{}, nil
}

func (t *torrentFS) OpenFile(name string, flag int, perm os.FileMode) (afero.File, error) {
	return t.Open(name)
}

func (t *torrentFS) Mkdir(name string, perm os.FileMode) error {
	return errors.New("not supported yet")
}

func (t *torrentFS) MkdirAll(path string, perm os.FileMode) error {
	return errors.New("not supported yet")
}

func (t *torrentFS) Remove(name string) error {
	return errors.New("not supported yet")
}

func (t *torrentFS) RemoveAll(path string) error {
	return errors.New("not supported yet")
}

func (t *torrentFS) Rename(oldname, newname string) error {
	return errors.New("not supported yet")
}

func (t *torrentFS) Stat(name string) (os.FileInfo, error) {
	return nil, errors.New("not supported yet")
}

func (t *torrentFS) Chmod(name string, mode os.FileMode) error {
	return errors.New("not supported yet")
}

func (t *torrentFS) Chtimes(name string, atime time.Time, mtime time.Time) error {
	return errors.New("not supported yet")
}