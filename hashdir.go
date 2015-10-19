package gohashdir

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

const SLASH string = "/"
const HASH_FORMAT = "%016x"

func uint64ToHashPath(width int, n uint64) (string, error) {
	if width < 1 {
		return "", errors.New("Width cannot be less than zero")
	}
	s := fmt.Sprintf(HASH_FORMAT, n)
	return StringToDirsString(width, s)

}

func StringToDirsString(width int, s string) (string, error) {
	if width < 1 {
		return "", errors.New("Width cannot be less than zero")
	}
	dirs := ""
	for i, c := range s {
		if i != 0 && i%width == 0 {
			dirs += SLASH
		}
		ch, err := strconv.Unquote(strconv.QuoteRune(c))
		if err != nil {
			return "", err
		}
		dirs += ch
	}
	dirs += SLASH
	return filepath.FromSlash(dirs), nil
}

func baseDirExists(baseDir string) error {
	_, err := os.Stat(baseDir)
	if err != nil {
		return err
	}
	return nil
}

func makeDir(dir string) (existed bool, err error) {
	var stat os.FileInfo
	stat, err = os.Stat(dir)
	if os.IsNotExist(err) {
		existed = false
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			dir = ""
		}
	} else {
		if stat != nil {
			existed = true
		}
	}
	return
}

func HashDirString(baseDir string, width int, s string) (dir string, existed bool, err error) {
	err = baseDirExists(baseDir)
	if err != nil {
		return
	}
	dir, err = StringToDirsString(width, s)
	if err != nil {
		return
	}
	dir = baseDir + "/" + dir
	existed, err = makeDir(dir)
	log.Println(dir)
	return
}

func HashDir(baseDir string, width int, n uint64) (dir string, existed bool, err error) {
	if width < 1 {
		return "", false, errors.New("Width cannot be less than zero")
	}
	err = baseDirExists(baseDir)
	if err != nil {
		return
	}

	existed = false
	dir, err = uint64ToHashPath(width, n)
	dir = baseDir + "/" + dir

	existed, err = makeDir(dir)

	return
}
