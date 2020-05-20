package utils

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

// FileExists returns true if a file at the passed string exists.
func FileExists(name string) bool {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		return false
	}
	return true
}

func TooPermissive(fileMode, maxAllowedPerms os.FileMode) bool {
	return fileMode&^maxAllowedPerms != 0
}

// Ensures that the given path exists, that it's a directory, and that it has
// permissions that are no more permissive than the given ones.
//
// - If the path does not exist, it is created
// - If the path exists, but is not a directory, an error is returned
// - If the path exists, and is a directory, but has the wrong perms, it is chmod'ed
func EnsureDirAndMaxPerms(path string, perms os.FileMode) error {
	stat, err := os.Stat(path)
	if err != nil && !os.IsNotExist(err) {
		// Regular error
		return err
	} else if os.IsNotExist(err) {
		// Dir doesn't exist, create it with desired perms
		return os.MkdirAll(path, perms|os.ModeDir)
	} else if !stat.IsDir() {
		// Path exists, but it's a file, so don't clobber
		return errors.Errorf("%v already exists and is not a directory", path)
	} else if stat.Mode() != perms {
		// Dir exists, but wrong perms, so chmod
		return os.Chmod(path, (stat.Mode()&perms)|os.ModeDir)
	}
	return nil
}

// Writes `data` to `path` and ensures that the file has permissions that
// are no more permissive than the given ones.
func WriteFileWithMaxPerms(path string, data []byte, perms os.FileMode) error {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, perms)
	if err != nil {
		return err
	}
	defer f.Close()
	err = EnsureFileMaxPerms(f, perms)
	if err != nil {
		return err
	}
	_, err = f.Write(data)
	return err
}

// Copies the file at `srcPath` to `dstPath` and ensures that it has
// permissions that are no more permissive than the given ones.
func CopyFileWithMaxPerms(srcPath, dstPath string, perms os.FileMode) error {
	src, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.OpenFile(dstPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, perms)
	if err != nil {
		return err
	}
	defer dst.Close()

	err = EnsureFileMaxPerms(dst, perms)
	if err != nil {
		return err
	}

	_, err = io.Copy(dst, src)

	return err
}

// Ensures that the given file has permissions that are no more
// permissive than the given ones.
func EnsureFileMaxPerms(file *os.File, perms os.FileMode) error {
	stat, err := file.Stat()
	if err != nil {
		return err
	}
	if stat.Mode() == perms {
		return nil
	}
	return file.Chmod(stat.Mode() & perms)
}

// Ensures that the file at the given filepath has permissions that are
// no more permissive than the given ones.
func EnsureFilepathMaxPerms(filepath string, perms os.FileMode) error {
	dst, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, perms)
	if err != nil {
		return err
	}
	defer dst.Close()

	return EnsureFileMaxPerms(dst, perms)
}

// FilesInDir returns an array of filenames in the directory.
func FilesInDir(dir string) ([]string, error) {
	f, err := os.Open(dir)
	if err != nil {
		return []string{}, err
	}
	defer f.Close()

	r, err := f.Readdirnames(-1)
	if err != nil {
		return []string{}, err
	}

	return r, nil
}

// FileContents returns the contents of a file as a string.
func FileContents(path string) (string, error) {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(dat), nil
}
