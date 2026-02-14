package main

import (
	//"strconv"
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("wc", "-c", "init.go")

	// The `Output` method executes the command and
	// collects the output, returning its value
	out, err := cmd.Output()
	if err != nil {
		// if there was any error, print it here
		fmt.Println("could not run command: ", err)
	}
	// otherwise, print the output from running the command
	//trimmed := bytes.Trim(out, "init.go")
	hola := string(out)
	var myInt int
	_, err = fmt.Sscanf(hola, "%d", &myInt)
	if err != nil {
   		fmt.Println(err)
	}else{
		fmt.Println(string(myInt))
	}
}
