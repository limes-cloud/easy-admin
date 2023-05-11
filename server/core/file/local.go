package file

import (
	"errors"
	"github.com/limeschool/easy-admin/server/config"
	"io"
	"os"
	"strings"
)

type local struct {
	conf *config.File
}

func NewLocal(conf *config.File) (Store, error) {
	if conf.LocalDir == "" {
		return nil, errors.New("upload config error")
	}
	return &local{
		conf: conf,
	}, nil
}

func (s *local) Put(path string, r io.Reader) error {
	if err := s.makeDir(path); err != nil {
		return err
	}

	saveFile, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer saveFile.Close()
	_, err = io.Copy(saveFile, r)
	return nil
}

func (s *local) PutFromLocal(path string, localPath string) error {

	if err := s.makeDir(path); err != nil {
		return err
	}
	return os.Rename(localPath, path)
}

func (s *local) Get(key string) (io.ReadCloser, error) {
	path := s.conf.LocalDir + "/" + key
	return os.OpenFile(path, os.O_RDONLY, os.ModePerm)
}

func (s *local) Delete(key string) error {
	return os.Remove(s.conf.LocalDir + "/" + key)
}

func (s *local) Size(key string) (int64, error) {
	path := s.conf.LocalDir + "/" + key
	info, err := os.Stat(path)
	if err != nil {
		return 0, err
	}
	return info.Size(), nil
}

func (s *local) Exists(key string) (bool, error) {
	_, err := os.Stat(key)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func (s *local) makeDir(path string) error {
	dir := path[:strings.LastIndex(path, "/")]
	if is, err := s.Exists(dir); !is {
		if err != nil {
			return err
		}
		if err = os.MkdirAll(dir, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}
