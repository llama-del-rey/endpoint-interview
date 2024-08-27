package directory

import (
	"fmt"
	"sort"
	"strings"
)

// Directory represents a directory in the filesystem.
type Directory struct {
	Name           string
	Subdirectories map[string]*Directory
}

var root = &Directory{Name: "", Subdirectories: make(map[string]*Directory)}

// Root returns the root directory.
func Root() *Directory {
	return root
}

// CreateDir creates a directory at the specified path.
func CreateDir(path string) {
	parts := strings.Split(path, "/")
	currentDir := root
	for _, part := range parts {
		if _, exists := currentDir.Subdirectories[part]; !exists {
			currentDir.Subdirectories[part] = &Directory{Name: part, Subdirectories: make(map[string]*Directory)}
		}
		currentDir = currentDir.Subdirectories[part]
	}
	fmt.Printf("CREATE %s\n", path)
}

// MoveDir moves a directory from sourcePath to destPath.
func MoveDir(sourcePath, destPath string) {
	sourceParts := strings.Split(sourcePath, "/")
	destParts := strings.Split(destPath, "/")

	sourceParent := getParentDir(sourceParts[:len(sourceParts)-1])
	destDir := getDir(destParts)

	if sourceParent == nil || destDir == nil {
		fmt.Printf("Cannot move %s - path does not exist\n", sourcePath)
		return
	}

	dirName := sourceParts[len(sourceParts)-1]
	dirToMove, exists := sourceParent.Subdirectories[dirName]
	if !exists {
		fmt.Printf("Cannot move %s - path does not exist\n", sourcePath)
		return
	}

	destDir.Subdirectories[dirName] = dirToMove
	delete(sourceParent.Subdirectories, dirName)
	fmt.Printf("MOVE %s %s\n", sourcePath, destPath)
}

// DeleteDir deletes the directory at the specified path.
func DeleteDir(path string) {
	parts := strings.Split(path, "/")
	parentDir := getParentDir(parts[:len(parts)-1])

	if parentDir == nil {
		fmt.Printf("Cannot delete %s - %s does not exist\n", path, parts[0])
		return
	}

	dirName := parts[len(parts)-1]
	if _, exists := parentDir.Subdirectories[dirName]; exists {
		delete(parentDir.Subdirectories, dirName)
	} else {
		fmt.Printf("Cannot delete %s - %s does not exist\n", path, parts[len(parts)-2])
	}
}

// ListDirs lists all directories starting from the specified directory, sorted alphabetically.
func ListDirs(dir *Directory, prefix string) {
	if dir.Name != "" {
		fmt.Println(prefix + dir.Name)
		prefix += "  "
	}

	// Sort the subdirectory names alphabetically
	var subdirNames []string
	for name := range dir.Subdirectories {
		subdirNames = append(subdirNames, name)
	}
	sort.Strings(subdirNames)

	for _, name := range subdirNames {
		ListDirs(dir.Subdirectories[name], prefix)
	}
}

func getDir(parts []string) *Directory {
	currentDir := root
	for _, part := range parts {
		if nextDir, exists := currentDir.Subdirectories[part]; exists {
			currentDir = nextDir
		} else {
			return nil
		}
	}
	return currentDir
}

func getParentDir(parts []string) *Directory {
	if len(parts) == 0 {
		return root
	}
	return getDir(parts)
}
