package main

import(
	"os"
	"fmt"
)

type Index struct {
	Name string
	Path_id string 
	Permissions int
	Stage_number int
}

func (i Index) get_name(f *os.File) {
	i.name := f.Name()
}

func (i index) get_id(f *os.File) {
	i.Path_id := Get_Hash(f)
}

func (i index) get_permissions(f os.File) { //each file need to have the 644 | 755 permissions.
	
}