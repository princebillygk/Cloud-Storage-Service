package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	choice := ""
	if len(os.Args) > 1 {
		choice = os.Args[1]
	}

	switch choice {
	case "create":
		RunWithSpaceName(createSpace)

	case "list":
		showList()

	case "key":
		RunWithSpaceId(getKey)

	case "edit":
		RunWithSpaceId(editSpace)

	case "delete":
		RunWithSpaceId(deleteSpace)

	default:
		fmt.Printf("The program requires a valid argument. Valid arguments are:\n" +
			"create: Create a new space\n" +
			"list: List all spaces\n" +
			"key: Show access key\n" +
			"edit: Edit space name\n" +
			"delete: Delete a space\n")
	}
}

func RunWithSpaceName(fn func(string)) {
	if len(os.Args) < 3 {
		fmt.Println("A name is required as argument")
		return
	}
	spacename := os.Args[2]
	fn(spacename)
}

func RunWithSpaceId(fn func(int)) {
	if len(os.Args) < 3 {
		fmt.Println("A space id is required as argument")
		return
	}
	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Space id is not valid")
		return
	}
	fn(id)
}

func createSpace(spacename string) {
	/**
	* Creates a new space and save it to database, add to ram
	* Creates a folder with the unique id in the system
	 */
	fmt.Printf("Create space: %s\n", spacename)
}

func showList() {
	/** Show list of all spaces from database */
	fmt.Println("Show list of all spaces")

}

func getKey(spaceId int) {
	/** Get access key from generated space stored in database */
	fmt.Printf("Get Key: %d\n", spaceId)
}

func editSpace(spaceId int) {
	/** Edit a existing space name and update the database */
	fmt.Printf("Edit space name: %d\n", spaceId)
}

func deleteSpace(spaceId int) {
	/** Removes a space from database and system */
	fmt.Printf("Delete space: %d\n", spaceId)
}
