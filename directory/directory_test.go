package directory

import (
	"testing"
)

func TestCreateDir(t *testing.T) {
	root = &Directory{Name: "", Subdirectories: make(map[string]*Directory)}

	CreateDir("fruits")
	CreateDir("vegetables")
	CreateDir("fruits/apples")
	CreateDir("fruits/apples/fuji")

	if _, exists := root.Subdirectories["fruits"]; !exists {
		t.Errorf("Expected 'fruits' directory to be created")
	}

	if _, exists := root.Subdirectories["vegetables"]; !exists {
		t.Errorf("Expected 'vegetables' directory to be created")
	}

	if _, exists := root.Subdirectories["fruits"].Subdirectories["apples"]; !exists {
		t.Errorf("Expected 'apples' directory to be created under 'fruits'")
	}

	if _, exists := root.Subdirectories["fruits"].Subdirectories["apples"].Subdirectories["fuji"]; !exists {
		t.Errorf("Expected 'fuji' directory to be created under 'apples'")
	}
}

func TestMoveDir(t *testing.T) {
	root = &Directory{Name: "", Subdirectories: make(map[string]*Directory)}

	CreateDir("fruits")
	CreateDir("vegetables")
	CreateDir("fruits/apples")

	MoveDir("fruits/apples", "vegetables")

	if _, exists := root.Subdirectories["fruits"].Subdirectories["apples"]; exists {
		t.Errorf("Expected 'apples' to be moved out of 'fruits'")
	}

	if _, exists := root.Subdirectories["vegetables"].Subdirectories["apples"]; !exists {
		t.Errorf("Expected 'apples' to be moved under 'vegetables'")
	}
}

func TestDeleteDir(t *testing.T) {
	root = &Directory{Name: "", Subdirectories: make(map[string]*Directory)}

	CreateDir("fruits")
	CreateDir("fruits/apples")
	DeleteDir("fruits/apples")

	if _, exists := root.Subdirectories["fruits"].Subdirectories["apples"]; exists {
		t.Errorf("Expected 'apples' directory to be deleted")
	}
}

func TestListDirs(t *testing.T) {
	root = &Directory{Name: "", Subdirectories: make(map[string]*Directory)}

	CreateDir("fruits")
	CreateDir("vegetables")
	CreateDir("fruits/apples")

	ListDirs(root, "")
}

func TestMoveDir_NonExistent(t *testing.T) {
	root = &Directory{Name: "", Subdirectories: make(map[string]*Directory)}

	CreateDir("fruits")
	CreateDir("vegetables")

	MoveDir("fruits/nonexistent", "vegetables")

	if _, exists := root.Subdirectories["vegetables"].Subdirectories["nonexistent"]; exists {
		t.Errorf("Did not expect 'nonexistent' directory to be moved")
	}
}

func TestDeleteDir_DeleteChilde(t *testing.T) {
	root = &Directory{Name: "", Subdirectories: make(map[string]*Directory)}

	CreateDir("fruits")

	DeleteDir("fruits/nonexistent")

	if _, exists := root.Subdirectories["fruits"]; !exists {
		t.Errorf("Expected 'fruits' directory to still exist")
	}
}
