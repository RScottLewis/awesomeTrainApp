//package awesomeTrainApp
// Fetch prints the content found at each specified URL.
package main

import (
//"io/ioutil"
"log"
"net/http"
"os/exec"
"path/filepath"
"os"
"fmt"


)


func run() ([]string, error) {
	searchDir := "."

	fileList := make([]string, 0)
	e := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return err
	})

	if e != nil {
		panic(e)
	}

	for _, file := range fileList {
		fmt.Println(file)
	}

	return fileList, nil
}

func main() {

	for _, url := range os.Args[1:] {
	resp, err := http.Get(url)
	if err != nil {
	fmt.Fprintf(os.Stderr, "exam: %v\n", err)
	os.Exit(1)
	}
//err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
	fmt.Fprintf(os.Stderr, "exam: reading %s: %v\n", url, err)
	os.Exit(1)
	}


	//for _, Dirs := range os.Dirs[1:]

	dirs:= make([]string, 0)
	dirs, _ = run()

		var buffer string
		for _, dir := range dirs {
			buffer="java -jar jawk.jar -F, -f "
			buffer=buffer+dir
			buffer=buffer+"dijkstra.awk"
			buffer=buffer+dir
			buffer=buffer+"exam.txt"

			cmd := exec.Command(buffer)
			err = cmd.Start()
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Waiting for _, command to finish...")
			err = cmd.Wait()
			log.Printf("Command finished with error: %v", err)
		}
	}

  //cmd := exec.Command("gawk -F, -f dijkstra.awk exam.txt")


}

//!-
