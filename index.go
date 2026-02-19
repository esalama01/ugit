package main
import (
	"os"
	"fmt"
)

type Metadata struct {
	Permissions string
	Last_modification string
	size int
}

func (data Metadata) get_metadata(f *os.File) (string, string, int64){

	//--------the date of the last modification
	fileInfo, err := os.Stat(f.Name())
	check(err)
	modificationTime := fileInfo.ModTime()
	mod_time := modificationTime.Format("2006-01-02 15:04:05") 
	
	//-------- The file's permissions
	mode := fileInfo.Mode()
	permissions :=  fmt.Sprintf("%o", mode.Perm())
	
	//-------- the file's size
	size := fileInfo.Size()
	//------------------------------------------------------
	return mod_time, permissions, size
}

type Index struct {
	Path string
	Id string
	Type string
	metadata *Metadata
	Stage_number int
}

func (blob Blob) get_object_type() string {
	return "Blob"
}

func (blob Blob) get_path(f *os.File) string{
	return f.Name()
}

func (blob Blob) get_id(f *os.File) string{
	hash, _ := Get_Hash(f)
	return hash
}

func (blob Blob) get_stage_number(number int) int {
	return number
}
/*
func (blob Blob) get_permissions(f os.File) string{ //each file need to have the 644 | 755 permissions.
	info, err := os.Stat(f.Name())
	check(err)
	mode := info.Mode()
	return mode.String() //return the permissions on the file in string format
}
*/

type StagingArea struct {
	entries map[string] *Index //a mapping between a file name and it s Index format
}

