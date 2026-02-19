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
		if args[2] == "w"{
			Ugit_hash_object_w(&args[3])
		}else{
			Ugit_hash_object(&args[2])
		}
	default:
		fmt.Println("holaa")
	}
}