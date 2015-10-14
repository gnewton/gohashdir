package gohashdir

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

const SLASH string = "/"
const HASH_FORMAT = "%016x"

func uint64ToHashPath(width int, n uint64) (string, error) {
	s := fmt.Sprintf(HASH_FORMAT, n)
	v := ""
	for i, c := range s {
		if i != 0 && i%width == 0 {
			v += SLASH
		}
		ch, err := strconv.Unquote(strconv.QuoteRune(c))
		if err != nil {
			return "", err
		}
		v += ch
	}
	v += SLASH
	return filepath.FromSlash(v), nil
}

func HashDir(baseDir string, width int, n uint64) (dir string, existed bool, err error) {
	if width < 1 {
		return "", false, errors.New("Width cannot be less than zero")
	}
	_, err = os.Stat(baseDir)
	if err != nil {
		return "", false, err
	}

	existed = false
	dir, err = uint64ToHashPath(width, n)
	if err != nil {
		dir = ""
	} else {
		dir = baseDir + "/" + dir
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
	}
	return
}
