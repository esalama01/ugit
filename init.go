package main
import (
	"fmt"
	"strconv"

	//"io/fs"
	"log"
	"os"

	//"path/filepath"
	"gopkg.in/ini.v1"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func mkdirrr() {
	//creating the .ugit directory
	err := os.Mkdir(".ugit", 0755)
	check(err)

	//creating the hooks subd
	err = os.MkdirAll(".ugit/hooks", 0755)
	check(err)

	//creating the info subd
	err = os.MkdirAll(".ugit/info", 0755)
	check(err)

	//creating the objects subd
	err = os.MkdirAll(".ugit/objects", 0755)
	check(err)
	
	//creating the objects/info subd
	err = os.MkdirAll(".ugit/objects/info", 0755)
	check(err)

	//creating the objects/pack subd
	err = os.MkdirAll(".ugit/objects/pack", 0755)
	check(err)

	//creating the refs subd
	err = os.MkdirAll(".ugit/refs", 0755)
	check(err)
	//creating the refs/heads subd
	err = os.MkdirAll(".ugit/refs/heads", 0755)
	check(err)
}

func create() {

	targetDir := ".ugit/" //go inside the .ugit directory
	if err := os.Chdir(targetDir); err != nil {
		log.Fatalf("Error : %v\n", err)
	}
	//creating the HEAD file (empty)
	file_name := "HEAD"
	p, err := os.Create(file_name)
	check(err)
	defer p.Close()

	//creating the config file (empty)
	file_name = "config.ini"
	p, err = os.Create(file_name)
	config(p)
	check(err)
	defer p.Close()

	//creating the description file (empty)
	file_name = "description"
	p, err = os.Create(file_name)
	description(p)
	check(err)
	defer p.Close()

	//creating the index file (empty)
	file_name = "index"
	p, err = os.Create(file_name)
	config(p)
	check(err)
	defer p.Close()
	
	targetDir = ".." //go back to the parent directory
	if err := os.Chdir(targetDir); err != nil {
		log.Fatalf("Error : %v\n", err)
	}
}

func config(file *os.File) {
	cfg, err := ini.Load(file)
	if err != nil {
		log.Fatalf("Something's wrong with the config file.")
	}
	/*
		core_section := cfg.Section("core")
		core_section.Key("repositoryformatversion").SetValue(0)
		core_section.Key("filemode").SetValue(true)
		core_section.Key("bare").SetValue(false)
		core_section.Key("logallrefupdates").SetValue(true)
	*/
	//creating the core section
	newSection, err := cfg.NewSection("core")
	if err != nil {
		log.Fatalf("Cannot create the new section.")
	}
	//adding mappings between keys and values.
	_, err = newSection.NewKey("repositoryformatversion", strconv.Itoa(0))
	if err != nil {
		log.Fatalf("Fail to add new key to new section: %v", err)
	}

	_, err = newSection.NewKey("filemode", strconv.FormatBool(true))
	if err != nil {
		log.Fatalf("Fail to add new key to new section: %v", err)
	}

	_, err = newSection.NewKey("bare", strconv.FormatBool(false))
	if err != nil {
		log.Fatalf("Fail to add new key to new section: %v", err)
	}

	_, err = newSection.NewKey("logallrefupdates", strconv.FormatBool(true))
	if err != nil {
		log.Fatalf("Fail to add new key to new section: %v", err)
	}

	if err := cfg.SaveTo("config.ini"); err != nil {
		log.Fatalf("Failed to save the configuration file.")
	}
}

func description(file *os.File) {
	data := []byte("Unnamed repository; edit this file 'description' to name the repository.")
	err := os.WriteFile(file.Name(), data, 0644)
	if err != nil {
		log.Fatalf("Couldn't write data to the description file.")
	}
}
func Init() {
	mkdirrr()
	create()
	fmt.Println("The .ugit folder created succesfully.")
}
