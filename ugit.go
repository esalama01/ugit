package main


import(
	"fmt"
	"os"
)

func main(){
	area := StagingArea{
        entries: make(map[string]*Index),
    }
	args := os.Args
	switch args[1]{
	case "init":
		Init()
	case "hash-object":
		if args[2] == "-w"{
			file, err := os.Open(args[3])
			check(err)
			defer file.Close()
			Ugit_hash_object_w(file)
		}else{
			file, err := os.Open(args[2])
			check(err)
			defer file.Close()
			Ugit_hash_object(file)
		}
	case "update-index":
		if args[2] == "--add"{
			file, err := os.Open(args[3])
			check(err)
			Ugit_update_index(file,&area)
			indexfileadd(&area)
		}else{
			file, err := os.Open(args[2])
			check(err)
			Ugit_update_index(file,&area)
		}

	default:
		fmt.Println("holaa")
	}
}