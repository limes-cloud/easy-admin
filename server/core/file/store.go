package file

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/limeschool/easy-admin/server/config"
	"github.com/limeschool/easy-admin/server/tools"
	"io"
	"os"
	"strings"
)

type Store interface {
	Put(key string, r io.Reader) error
	PutFromLocal(key string, localPath string) error
	Get(key string) (io.ReadCloser, error)
	Delete(key string) error
	Size(key string) (int64, error)
	Exists(key string) (bool, error)
}

type File interface {
	CheckType(name string) error
	CheckSize(size int64) error
	Upload(key string, r io.Reader) (string, error)
	UploadFromLocal(key string, localPath string) (string, error)
	GetString(key string) (string, error)
	GetBytes(key string) ([]byte, error)
	GetToFile(key string, localPath string) error
	Get(key string) (io.ReadCloser, error)
	Delete(key string) error
	Size(key string) (int64, error)
	Exists(key string) (bool, error)
}

type file struct {
	sub  string
	conf *config.File
	Store
}

// NewFile create new File instance.
func NewFile(conf *config.File, sub string) (File, error) {
	_, is := conf.SubDir[sub]
	if !is {
		return nil, errors.New("not support sub-dir")
	}
	var store Store
	var err error
	switch conf.DriveType {
	case "tencent":
		store, err = NewTencent(conf)
	case "qiniu":
		store, err = NewQiniu(conf)
	case "aliyun":
		store, err = NewAliyun(conf)
	case "huawei":
		store, err = NewHuawei(conf)
	case "minio":
		store, err = NewMinio(conf)
	case "local":
		store, err = NewLocal(conf)
	default:
		err = errors.New("not support drive")
	}
	return &file{
		conf:  conf,
		Store: store,
		sub:   sub,
	}, err
}

func (s *file) CheckType(name string) error {
	sub := s.conf.SubDir[s.sub]
	index := strings.LastIndex(name, ".")
	if index == -1 {
		return errors.New("file type error")
	}
	if !tools.InList(sub.Accepts, name[index+1:]) {
		return errors.New("file type not support")
	}
	return nil
}

func (s *file) CheckSize(size int64) error {
	sub := s.conf.SubDir[s.sub]
	if sub.MaxSize > 0 && int(size/1024) > sub.MaxSize {
		return errors.New("file size overflow allow size")
	}
	return nil
}

func (s *file) Rename(fileName string) string {
	sub := s.conf.SubDir[s.sub]
	if !sub.Rename {
		return fileName
	}
	uid := fmt.Sprintf("%x", md5.Sum([]byte(uuid.New().String())))
	fileType := fileName[strings.LastIndex(fileName, "."):]
	return uid + fileType
}

func (s *file) Upload(key string, r io.Reader) (string, error) {
	// 检验类型
	if err := s.CheckType(key); err != nil {
		return "", err
	}

	key = s.sub + "/" + s.Rename(key)
	if s.conf.LocalDir != "" {
		key = s.conf.LocalDir + "/" + key
	}
	return key, s.Store.Put(key, r)
}

func (s *file) UploadFromLocal(key string, localPath string) (string, error) {
	// 检验类型
	if err := s.CheckType(key); err != nil {
		return "", err
	}

	// 检验大小
	info, err := os.Stat(localPath)
	if err != nil {
		return "", err
	}
	if err = s.CheckSize(info.Size()); err != nil {
		return "", err
	}

	key = s.sub + "/" + s.Rename(key)
	if s.conf.LocalDir != "" {
		key = s.conf.LocalDir + "/" + key
	}
	return key, s.Store.PutFromLocal(key, localPath)
}

// GetString gets the file pointed to by key and returns a string.
func (s *file) GetString(key string) (string, error) {
	bs, err := s.GetBytes(key)
	if err != nil {
		return "", err
	}

	return string(bs), nil
}

// GetBytes gets the file pointed to by key and returns a byte array.
func (s *file) GetBytes(key string) (bytes []byte, err error) {
	rc, err := s.Get(key)
	if err != nil {
		return
	}

	defer func() {
		err = rc.Close()
	}()

	return io.ReadAll(rc)
}

// GetToFile saves the file pointed to by key to the localPath.
func (s *file) GetToFile(key string, localPath string) (err error) {
	rc, err := s.Get(key)
	if err != nil {
		return err
	}

	defer func(rc io.ReadCloser) {
		err = rc.Close()
	}(rc)

	f, _ := os.OpenFile(localPath, os.O_CREATE|os.O_WRONLY, 0o644)
	defer func(f *os.File) {
		err = f.Close()
	}(f)

	_, err = io.Copy(f, rc)

	return err
}
