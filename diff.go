package main

import(
	"github.com/neticdk/go-stdlib/diff/myers"
	"os"
	"bytes"
	"compress/zlib"
	"strings"
	"io"
)

type Sha1_path struct{
	Checksum string
	Path string
}


func Compare_with_sha1() ([]*Sha1_path){//a function to compare whatever s in the index file with my current working directory. But adds the sha1 of the file(the one in the index) to the result
	area := StagingArea{
        entries: make(map[string]*Index),
    }
	Loadfromtheindex(&area)
	paths := Traversal()
	var m []*Sha1_path
	var exists []string //i ll need it to point to removed files
	//i ll traverse the paths slice  and check if each path exists in my index
	for _, path := range paths { 
		f, err := os.Open(path) //open the file to pass it as a parameter in the Get_Hash function.
		check(err)
		value, ok := area.entries[path]
		if ok { //if it exists
			val1, _ := Get_Hash_Blob(f)
			f.Close()
			exists = append(exists, path)
			if val1 != value.Id{ //if they re the same
				instance := Sha1_path{
					Checksum : value.Id,
					Path : path,
				}
				m = append(m, &instance)
			}else{
			}
		}else{//if it doesn t exist
		}
		
	}
	return m
}

func Ugit_cat_file(path string) string{
	content, err := os.ReadFile(path)
	check(err)
	compressed := []byte(content)
	b := bytes.NewReader(compressed)
	r, err := zlib.NewReader(b)
	check(err)
	defer r.Close()
	var sb strings.Builder
	_, err = io.Copy(&sb, r)
	check(err)
	text := sb.String()
	//i need to remove the header from the text.
}