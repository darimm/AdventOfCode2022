package Day7all

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/darimm/AOCFunctions"
)

type Day7 struct {
	data []string
}

func (d *Day7) New() {
	lines, err := AOCFunctions.ReadFile(".\\Day7\\a\\day7a.txt")
	if err != nil {
		panic("Could not read input")
	}
	d.data = lines
}

func Day7all() {
	var day Day7
	day.New()
	day.PopulateFilesAndFolders()
}

type Node struct {
	Name        string
	Parent      *Node
	Directories []*Node
	Files       []int
	Size        int
}

func changeDirectory(dir string, node *Node) *Node {
	if dir == ".." { // go up a level.
		return node.Parent
	}
	for _, v := range node.Directories { // check to see if child already exists, move current pointer to it.
		if v.Name == dir {
			return v
		}
	}
	panic("You shouldn't get here. If you do, your output is wrong man.") //something has gone wrong
}

func (d *Day7) PopulateFilesAndFolders() {
	var currentNode *Node
	var root = Node{
		Name: "/",
	}
	currentNode = &root

	for i := 1; i < len(d.data); i++ {

		if len(d.data[i]) == 0 { // ignore empty lines
			continue
		}

		parse := strings.Fields(d.data[i])

		if len(parse) == 3 && parse[0] == "$" { //cd command
			currentNode = changeDirectory(parse[2], currentNode)
		}

		if len(parse) == 2 && parse[0] == "$" { //ls command, we don't actually need to do anything with this
			continue
		}

		if len(parse) == 2 && parse[0] != "$" { //results of ls command
			if parse[0] == "dir" { //either a directory
				createDirectory(parse[1], currentNode)
			} else { //or a file
				size, err := strconv.Atoi(parse[0])
				if err != nil {
					panic(fmt.Sprintf("An error has occurred converting %v, your output is wrong", parse[0]))
				}
				currentNode.Files = append(currentNode.Files, size)
			}
		}
	}
	CalculateFoldserSize(&root)
	fmt.Println(Test7a(&root))
	fmt.Println("Root.Size: ", root.Size)
	spaceneeded := (30000000 - (70000000 - root.Size))
	fmt.Println("Space needed: ", spaceneeded)
	fmt.Println(Test7b(&root, spaceneeded))
}

func Test7b(node *Node, spaceNeeded int) int {
	var result int

	myValue := spaceNeeded - node.Size
	fmt.Println("My Value: ", myValue)
	if node.Directories != nil {
		for _, v := range node.Directories {
			if spaceNeeded-v.Size < 0 && spaceNeeded-v.Size > myValue {
				result = Test7b(v, spaceNeeded)
			}
		}
	}
	if result != 0 {
		return result
	}
	return node.Size
}

func Test7a(node *Node) int {
	size := 0
	if node.Directories != nil {
		for _, v := range node.Directories {
			result := Test7a(v)
			size += result
		}
	}
	if node.Size <= 100000 {
		return node.Size + size
	}
	return size
}

func CalculateFoldserSize(node *Node) int {
	size := 0
	if node.Directories != nil {
		for _, v := range node.Directories {
			size += CalculateFoldserSize(v)
		}
	}
	if len(node.Files) != 0 {
		for _, v := range node.Files {
			size += v
		}
	}
	node.Size = size
	return size
}

func createDirectory(dir string, node *Node) {
	if node.Directories == nil { // node has no children recorded, create a new child node.
		var newNodedir = Node{
			Name:   dir,
			Parent: node,
		}
		node.Directories = append(node.Directories, &newNodedir)
		return
	} else { // node already has children recorded
		for _, v := range node.Directories { // check to see if child already exists, move current pointer to it.
			if v.Name == dir {
				return
			}
		}
		var newNodedir = Node{ // child does not exist since the for loop completed, create a new one.
			Name:   dir,
			Parent: node,
		}
		node.Directories = append(node.Directories, &newNodedir)
	}
}
