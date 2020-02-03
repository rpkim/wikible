package wiki

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"

	rand "wikible/util"
	api "wikible/wiki/api"
)

//WikiNode is a tree for wiki contents
type WikiNode struct {
	Parent    *WikiNode
	Child     []*WikiNode
	ContentID string
	Numbering string
	Title     string
	Sequence  int
}

func (node *WikiNode) addChild(child *WikiNode) {
	node.Child = append(node.Child, child)
}

//CreateNode is for creating wiki node
func CreateNode(parent *WikiNode, sequence int, numbering string, title string) *WikiNode {
	node := WikiNode{
		Parent:    parent,
		Numbering: numbering,
		Sequence:  sequence,
		Title:     title,
	}
	if parent != nil {
		parent.addChild(&node)
	}

	return &node
}

func syntaxCheck(str string) bool {
	match, _ := regexp.MatchString("^-+\\s.+", str)
	return match
}

func generateWikiNode(scanner *bufio.Scanner) (root *WikiNode) {
	var currentNode *WikiNode
	prev := 0
	for scanner.Scan() {
		text := scanner.Text()

		//syntax check
		if syntaxCheck(text) {
			//getting the depth(length)
			depth := len(strings.Split(text, " ")[0])
			//trim two characters("-"," ")
			text = strings.Trim(text, "- ")

			var parent *WikiNode
			//first time
			if prev == 0 {
				//first node
				root = CreateNode(nil, 1, "1.", text)
				currentNode = root
			} else {
				//compare the previous depth and current depth
				if prev == depth {
					parent = currentNode.Parent
				} else if prev > depth {
					parent = currentNode.Parent.Parent
				} else if prev < depth {
					parent = currentNode
				}

				sequence := len(parent.Child) + 1
				numbering := parent.Numbering + strconv.Itoa(sequence) + "."
				newNode := CreateNode(parent, sequence, numbering, text)
				currentNode = newNode
			}

			//update prev depth
			prev = depth

		} else {
			fmt.Println("File syntax Error")
		}

	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return root
}

//GenerateWikiNode is creating wiki node
func GenerateWikiNodeFromString(str string) (root *WikiNode) {
	//file read
	scanner := bufio.NewScanner(strings.NewReader(str))
	return generateWikiNode(scanner)
}

//GenerateWikiNode is creating wiki node
func GenerateWikiNodeFromFile(filepath string) (root *WikiNode) {
	//file read
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	return generateWikiNode(scanner)
}

//PrintWikiNode is for printing the wiki node
func PrintWikiNode(wikinode *WikiNode) {
	fmt.Println(wikinode.Numbering + " " + wikinode.Title)

	for _, element := range wikinode.Child {
		child := element
		PrintWikiNode(child)
	}
}

func CreateWikiNode(wg *sync.WaitGroup, wiki *api.Wiki, space api.Space, contentID string, wikinode *WikiNode) bool {
	defer wg.Done()
	res, err := wiki.CreateContent(wikinode.Numbering+" "+wikinode.Title, contentID, space)

	if err != nil {
		var errorResponse api.ErrorResponse
		err = json.Unmarshal(res, &errorResponse)
		if err != nil {
			//other error
			fmt.Println("ErrorResponse Unmarshal Error")
			return false
		} else {
			if strings.Contains(errorResponse.Message, "A page with this title already exists") {
				newTitle := wikinode.Numbering+" "+wikinode.Title+"_"+rand.String(4)
				res, err = wiki.CreateContent(newTitle, contentID, space)				
				if err != nil {
					var errorResponse api.ErrorResponse
					err = json.Unmarshal(res, &errorResponse)
					fmt.Println("Error Message:", err)
					return false
				} else {
					fmt.Println("A page with this title already existes")
					fmt.Println("create a page with the below title")
					fmt.Println(newTitle)
				}
			} else {
				fmt.Println("Error Message:", errorResponse.Message)
				return false	
			}
		}
	}

	var content api.Content

	err = json.Unmarshal(res, &content)
	if err != nil {
		fmt.Println("Error Message:", err)
		return false
	}

	wikinode.ContentID = content.ID

	for _, element := range wikinode.Child {
		child := element
		wg.Add(1)
		go CreateWikiNode(wg, wiki, space, wikinode.ContentID, child)
	}

	return true
}

//CreateWikiAPI is a wrap API for wikiapi
func CreateWikiAPI(address string, basicAuth string) (*api.Wiki, error) {
	return api.CreateWikiAPI(address, basicAuth)
}
