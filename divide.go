package filedivdir

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

// Divide divides the files in the base directory by
// creating a directory for each file timestamp.
func Divide(base string) error {
	if files, err := ioutil.ReadDir(base); err != nil {
		return err
	} else if err := divide(base, files); err != nil {
		return err
	}
	return nil
}

func divide(base string, files []os.FileInfo) error {
	var made = make(map[string]bool)
	for _, file := range files {
		if !file.IsDir() && file.Name()[0:1] != "." {
			date := file.ModTime().Format("20060102")
			dst := filepath.Join(base, date)
			if _, ok := made[date]; !ok {
				if b, err := mkdir(dst); err != nil {
					return err
				} else if b {
					made[date] = true
				}
			}
			src := filepath.Join(base, file.Name())
			if err := exec.Command(moveCmd(), src, dst).Run(); err != nil {
				return err
			}
		}
	}
	return nil
}

func mkdir(dir string) (bool, error) {
	if f, err := os.Stat(dir); err != nil {
		// dir does not exist.
		if err := os.Mkdir(dir, 0755); err != nil {
			return false, err
		}
		return true, nil
	} else if !f.IsDir() {
		// dir is not direcroty.
		if err := os.Mkdir(dir, 0755); err != nil {
			return false, err
		}
		return true, nil
	}
	return false, nil
}

func moveCmd() string {
	if runtime.GOOS == "windows" {
		return "move"
	}
	return "mv"
}
