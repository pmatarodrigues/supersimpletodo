package main

import (
	"encoding/gob"
	"flag"
	"fmt"
	"os"
)

var (
	removeKey   *int
	listProject *string
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
		fmt.Printf("\n------------" + project + "------------\n")
		for key, item := range list {
			fmt.Printf("%d > %s \n", key, item)
		}
	}
}

func printItemsFromProject(projectToList string, items map[string][]string) {
	for project, list := range items {
		if project == projectToList {
			fmt.Printf("\n------------" + project + "------------\n")
			for key, item := range list {
				fmt.Printf("%d > %s \n", key, item)
			}
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

func removeItem(removeKey int, project string, items map[string][]string) map[string][]string {
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
	// changes the value of the variable
	removeKey = flag.Int("rm", -1, "Item key to Remove")
	listProject = flag.String("p", "", "Project to List")

	// get parameters passed
	// (not counting file name)
	data := os.Args[1:]
	// check if parameters were passed
	if len(data) > 0 {
		flag.Parse()
		isToRemove := isFlagPassed("rm")
		isToListProject := isFlagPassed("p")
		// r -> -r {key} {project} -> remove key from project
		// check if it's to remove or add item
		if isToListProject {
			fmt.Printf("______________# sst #______________")
			printItemsFromProject(*listProject, items)
		} else if isToRemove {
			project := data[2]
			items = removeItem(*removeKey, project, items)
		} else {
			// ./sst "do this do that" "project"
			newItem := data[0]
			project := data[1]

			// append item to project
			items[project] = append(items[project], newItem)
		}
		saveToFile(items, file)
	} else {
		fmt.Printf("______________# sst #______________")
		printItems(items)
	}
}
