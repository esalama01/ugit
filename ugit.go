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
	default:
		fmt.Println("holaa")
	}
}