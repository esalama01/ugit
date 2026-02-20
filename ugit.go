package main


import(
	"fmt"
	"os"
)

func main(){
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
	default:
		fmt.Println("holaa")
	}
}