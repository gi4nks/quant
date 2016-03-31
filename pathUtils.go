package quant

import (
	"os"
	"path/filepath"

	"github.com/kardianos/osext"
)

type PathUtils struct {
}

func NewPathUtils() *PathUtils {
	return &PathUtils{}
}

func (t PathUtils) ExistsPath(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func (t PathUtils) CreatePath(path string) error {
	pa, err := t.ExecutableFolder()
	if err == nil {
		return err
	}

	return os.Mkdir(pa+string(filepath.Separator)+path, 0777)
}

func (t PathUtils) ExecutableFolder() (string, error) {
	return osext.ExecutableFolder()
}
