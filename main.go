package main

import (
	"encoding/gob"
	"flag"
	"fmt"
	"os"
)

// parameters: filename
// returns: saved items
func saveToFile(items map[string][]string, file string) {
	// create new file
	encodeFile, err := os.Create(file)
	if err != nil {
		panic(err)
	}
	//
	encoder := gob.NewEncoder(encodeFile)
	// write data to the file
	if err := encoder.Encode(items); err != nil {
		panic(err)
	}
	encodeFile.Close()
}

func fileExists(file string) bool {
	_, err := os.Stat(file)
	return !os.IsNotExist(err)
}

func readFromFile(file string) map[string][]string {
	items := make(map[string][]string)

	// check if file exists
	exist := fileExists(file)
	if !exist {
		os.Create(file)
	}
	// open db file
	decodeFile, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer decodeFile.Close()

	// create decoder
	decoder := gob.NewDecoder(decodeFile)
	// decode read items
	decoder.Decode(&items)
	return items
}

func printItems(items map[string][]string) {
	for project, list := range items {
		fmt.Printf("\n------" + project + "------")
		for key, item := range list {
			fmt.Printf("\n[%d]- %s", key, item)
		}
	}
}

// check if some flag was provided on cli
func isFlagPassed(name string) bool {
	passed := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			passed = true
		}
	})
	return passed
}

func flagCheck(removeKey int) {
	// changes the value of the variable
	flag.IntVar(&removeKey, "r", -1, "Key to remove")
	flag.Parse()
}

func removeItem(removeKey int, project string, items map[string][]string) map[string][]string {
	fmt.Println(items[project])
	items[project] = append(items[project][:removeKey], items[project][removeKey+1:]...)

	// delete if project is empty
	if len(items[project]) == 0 {
		delete(items, project)
	}
	return items
}

func main() {
	items := make(map[string][]string)
	file := "./todo"
	items = readFromFile(file)
	var removeKey int

	// get parameters passed
	// (not counting file name)
	data := os.Args[1:]
	// check if parameters were passed
	if len(data) > 0 {
		project := data[0]
		fmt.Println("len")

		// r -> {project} -r {key} -> remove key from project
		// check if it's to remove or add item
		if true {
			fmt.Println("len")
			flagCheck(removeKey)
			items = removeItem(removeKey, project, items)
		} else {
			// ./sstodo "project" "do this do that"
			newItem := data[1]

			// append item to project
			items[project] = append(items[project], newItem)
		}
		saveToFile(items, file)
	} else {
		fmt.Printf("supersimpletodo")
		printItems(items)
	}
}
