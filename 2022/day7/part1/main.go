package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Sizer interface {
	getSize() int
}

type folder struct {
	name     string
	size     int
	father   *folder
	children map[string]Sizer
}

func (f *folder) getSize() int {
	var (
		size      int
		resultsCh = make(chan int)
	)

	var wg sync.WaitGroup

	for _, kid := range f.children {
		localKid := kid // avoid race condition
		wg.Add(1)
		go func() {
			defer wg.Done()
			resultsCh <- localKid.getSize()
		}()
	}

	go func() {
		wg.Wait()
		close(resultsCh)
	}()

	for res := range resultsCh {
		size += res
	}

	f.size = size
	fmt.Println("Directory:", f.name, "| size:", f.size)
	return size
}

type goFile struct {
	name string
	size int
}

func (f *goFile) getSize() int {
	return f.size
}

// Used to avoid a collision if there's a folder and file with the same name
func pathHash(path string) string {
	return path + "dir"
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var topDirectory *folder = &folder{name: "/", children: make(map[string]Sizer)}

	var currentDirectory *folder

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		switch parts[0] {
		case "$":
			command := parts[1]
			switch command {
			case "cd":
				path := parts[2]
				switch path {
				case "/":
					currentDirectory = topDirectory
				case "..":
					currentDirectory = currentDirectory.father // can fail if current dir is topDirectory
				default:
					currentDirectory = currentDirectory.children[pathHash(path)].(*folder)
				}
			case "ls":
				//Do nothing for now
			}
		case "dir":
			path := parts[1]
			if _, ok := currentDirectory.children[pathHash(path)]; !ok {
				currentDirectory.children[pathHash(path)] = &folder{name: path, children: make(map[string]Sizer), father: currentDirectory}
			}
		default:
			fileSize, _ := strconv.Atoi(parts[0])
			fileName := parts[1]
			currentDirectory.children[fileName] = &goFile{name: fileName, size: fileSize}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	topDirectory.getSize()

	smallFolders := make(chan *folder)

	var (
		walkFolder func(dir *folder)
		wg         sync.WaitGroup
	)

	walkFolder = func(dir *folder) {
		for _, children := range dir.children {
			folder, isFolder := children.(*folder)

			if !isFolder {
				continue
			}

			wg.Add(1)

			go func() {
				walkFolder(folder)
				wg.Done()
			}()

			if folder.size <= 100000 {
				smallFolders <- folder
			}
		}
	}

	wg.Add(1)
	go func() {
		walkFolder(topDirectory)
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(smallFolders)
	}()

	sum := 0

	for smallfolder := range smallFolders {
		sum += smallfolder.size
	}

	fmt.Println(sum)

}
