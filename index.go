package main
import (
	"os"
	"fmt"
	"strconv"
	"path/filepath"
	"bufio"
	"strings"
	"log"
)

type Metadata struct {
	Permissions *string
	Last_modification *string
	size *int64
}

func get_metadata(f *os.File) (*string, *string, *int64){

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
	return &mod_time, &permissions, &size
}

type Index struct {
	Path string
	Id string
	Type *string
	metadata *Metadata
	Stage_number int
}

func get_path(f *os.File) string{
	return f.Name()
}

func get_stage_number(number int) int {
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
	entries map[string] *Index //a mapping between a file's name and it s Index format
}

func indexfileline(idx *Index)(string){ //a method for the staging area.
	s1 := string(*idx.metadata.Permissions)
	s2 := string(idx.Id)
	s3 := strconv.Itoa(idx.Stage_number)
	s4 := idx.Path
	result := s1 + " " + s2 + " " + s3 + "         " + s4 +"\n"
	return result
}

func indexfileadd(area *StagingArea){ //writing to the index file
	file, err := os.OpenFile(".ugit/index", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	check(err)

	for _, idx := range area.entries {
		_, err = file.WriteString(indexfileline(idx))
		check(err)
	}
	
}

func Loadfromtheindex(area *StagingArea)(*StagingArea) { //a function that load entries from the index file and puts them into the staging area
	index, err := os.Open(".ugit/index")
	check(err)
	defer index.Close()
	scanner := bufio.NewScanner(index) //to scan the index file
	s := "blob"
	for scanner.Scan() {//scan line by line
		// Get the current line as a string
		line := scanner.Text()
		//from ech line i have to retrieve the path + the hash_id + type(blob by default) + stage number(0 by default) and i ll let the metadata empty.
		arr := strings.Fields(line) //i ll make an array of those vals.
		
		idx := Index { //filling the index struct
		Path : arr[3],
		Id : arr[2],
		Type : &s,
		Stage_number : get_stage_number(0),
		}

		area.entries[filepath.Clean(arr[3])] = &idx
	}

	if err := scanner.Err(); err != nil {//checks for errs during scanning
		log.Fatal(err)
	}
	return area
}

func Ugit_update_index(f *os.File, area *StagingArea) { // i will be implementing the git update-index --add command
	val1, _ := Get_Hash_Blob(f)
	mdata1, mdata2, mdata3 :=  get_metadata(f)
	s := "blob"
	mdata := Metadata{
		Permissions : mdata2,
		Last_modification : mdata1,
		size : mdata3,
	}
	idx := Index {
		Path : get_path(f),
		Id : val1,
		Type : &s,
		metadata : &mdata,
		Stage_number : get_stage_number(0),
	}
	area.entries[filepath.Clean(f.Name())] = &idx
}