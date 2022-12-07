package answers

import (
	"strconv"
	"strings"
)

func Day7() []interface{} {
	data := ReadInputAsStr(7)
	return []interface{}{q7part1(data), q7part2(data)}
}

type Directory struct {
	Name             string
	ParentDirectory  *Directory
	ChildDirectories []*Directory
	Files            []*File
}

func (d Directory) DirectSize() int {
	size := 0
	for _, file := range d.Files {
		size += file.Size
	}
	return size
}

func (d Directory) TotalSize() int {
	size := d.DirectSize()
	for _, dir := range d.ChildDirectories {
		size += dir.TotalSize()
	}
	return size
}

type File struct {
	Name string
	Size int
}

func NewDirectory(name string, parent *Directory) *Directory {
	return &Directory{
		Name:            name,
		ParentDirectory: parent,
	}
}

func ParseDirectory(data []string) ([]*Directory, *Directory) {
	rootDirectory := NewDirectory("/", &Directory{})
	activeDirectory := rootDirectory
	all := []*Directory{rootDirectory}
	for _, instruction := range data[1:] {
		instructionSplit := strings.Split(instruction, " ")
		if instructionSplit[0] == "$" && instructionSplit[1] == "cd" {
			folderName := instructionSplit[2]
			origin := activeDirectory
			if folderName == ".." {
				activeDirectory = activeDirectory.ParentDirectory
			} else {
				activeDirectory = NewDirectory(folderName, activeDirectory)
				all = append(all, activeDirectory)
				origin.ChildDirectories = append(origin.ChildDirectories, activeDirectory)
			}
		}
		if instructionSplit[0] != "$" && instructionSplit[0] != "dir" {
			// Listing files in active Directory
			size, err := strconv.Atoi(instructionSplit[0])
			if err != nil {
				panic(err)
			}
			newFile := File{Name: instructionSplit[1], Size: size}
			activeDirectory.Files = append(activeDirectory.Files, &newFile)
		}
	}
	return all, rootDirectory
}

func q7part1(data []string) int {
	all, _ := ParseDirectory(data)
	solution := 0
	for _, dir := range all {
		size := dir.TotalSize()
		if size <= 100000 {
			solution += size
		}
	}
	return solution
}

func q7part2(data []string) int {
	all, rootDirectory := ParseDirectory(data)
	unused := 70000000 - rootDirectory.TotalSize()
	necessaryToDelete := 30000000 - unused

	smallest := 999999999
	for _, dir := range all {
		if dir.TotalSize() > necessaryToDelete && dir.TotalSize() < smallest {
			smallest = dir.TotalSize()
		}
	}
	return smallest
}
