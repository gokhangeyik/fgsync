package tools

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

type FileList struct {
	ID       int32  `json:"id"`
	Name     string `json:"name"`
	FullPath string `json:"fullpath"`
	RelPath  string `json:"relpath"`
	Size     string `json:"size"`
}

// ListDir walks through the provided directory and returns a slice of FileList objects
// containing information about each file found (excluding directories).
//
// Parameters:
//   - dir_path: The path to the directory to list files from
//
// Returns:
//   - []FileList: A slice of FileList objects with file information
//   - error: An error if the directory doesn't exist or can't be read
func ListDir(dirPath string) ([]FileList, error) {
	dirPath = filepath.Clean(dirPath)
	if !os.IsPathSeparator(dirPath[len(dirPath)-1]) {
		dirPath = dirPath + string(os.PathSeparator)
	}

	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("folder not found: %s", dirPath)
	}

	files := make([]FileList, 0, 10)

	var fileID int32 = 1

	fileSystem := os.DirFS(dirPath)
	err := fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil // Skip all directories including root directory
		}

		info, err := d.Info()
		if err != nil {
			return err
		}
		fileSize := info.Size()

		files = append(files, FileList{
			ID:       fileID,
			Name:     info.Name(),
			FullPath: filepath.Join(dirPath, path),
			RelPath:  path,
			Size:     FormatBytes(fileSize),
		})

		fileID++

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("unable to scan directory: %v", err)
	}

	return files, nil
}
