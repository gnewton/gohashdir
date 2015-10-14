package gohashdir

import (
	//"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"testing"
)

func TestWidthsAndDirWrite(t *testing.T) {
	for width := 2; width < 8; width++ {
		tmpDir, err := ioutil.TempDir(".", "testDir")
		if err != nil {
			panic(err)
		}
		defer func() {
			// err = os.Chdir("..")
			// if err != nil {
			// 	log.Fatal(err)
			// }
			err := os.RemoveAll(tmpDir)
			if err != nil {
				log.Fatal(err)
			}
		}()

		// err = os.Chdir(tmpDir)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		var i uint64

		for i = 1; i < 2000; i++ {
			_, _, err := HashDir(tmpDir, width, uint64(rand.Int63()))
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
