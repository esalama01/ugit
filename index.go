package main
import (
	"os"
	"path/filepath"
)

type Index struct {
	Path string
	Id string 
	Permissions int
	Stage_number int
}

func (i Index) get_name(f *os.File) string{
	return filepath.Dir(f.Name())
}

func (i Index) get_id(f *os.File) string{
	hash, _ := Get_hash(f)
	return hash
}

func (i Index) get_permissions(f os.File) int{ //each file need to have the 644 | 755 permissions.
	
}