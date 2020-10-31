package main

import (
	"cloudstorageapi.com/models/spaceModel"
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

	case "info":
		RunWithSpaceId(getDetails)

	case "edit":
		RunWithSpaceId(editSpace)

	case "delete":
		RunWithSpaceId(deleteSpace)

	default:
		fmt.Printf("The program requires a valid argument. Valid arguments are:\n" +
			"create: Create a new space\n" +
			"list: List all spaces\n" +
			"info: Show details by ID\n" +
			"edit: Edit space name by ID\n" +
			"delete: Delete a space by ID\n")
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
	space := spaceModel.Space{Name: spacename}
	err := space.Save()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("New space entry created.\n"+
		"Space ID: %d\nName: %s\nAccess Token: %s\n",
		space.Id, space.Name, space.AccessToken)
}

func showList() {
	/** Show list of all spaces from database */
	spaces, err := spaceModel.All()
	if err != nil {
		fmt.Println(err)
		return
	}

	for i, sp := range spaces {
		fmt.Printf("%d. %s (ID:%d)\n", i+1, sp.Name, sp.Id)
	}
}

func getDetails(spaceId int) {
	/** Get access key from generated space stored in database */
	space, err := spaceModel.FindSpaceById(spaceId)
	if err != nil {
		fmt.Println(err)
		return
	}

	if space == nil {
		fmt.Println("No space found with this id")
	} else {
		fmt.Printf("Space ID: %d\nName: %s\nAccess Token: %s\n",
			space.Id, space.Name, space.AccessToken)
	}
}

func editSpace(spaceId int) {
	/** Edit a existing space name and update the database */

	//finding space in database
	space, err := spaceModel.FindSpaceById(spaceId)
	if err != nil {
		fmt.Println(err)
		return
	}
	if space == nil {
		fmt.Println("No space found with this id")
		return
	}
	fmt.Printf("Space ID: %d\nCurrent Name: %s\n", space.Id, space.Name)

	//collectin new name from user input
	var newName string
	fmt.Printf("New Name: ")
	fmt.Scanf("%s", &newName)

	//updating
	space.UpdateName(newName)
	fmt.Printf("Successfully change named to %s\n", space.Name)
}

func deleteSpace(spaceId int) {
	/** Removes a space from database and system */

	//finding
	fmt.Printf("Delete space: %d\n", spaceId)
	space, err := spaceModel.FindSpaceById(spaceId)
	if err != nil {
		fmt.Println(err)
		return
	}
	if space == nil {
		fmt.Println("No space found with this id")
		return
	}
	fmt.Printf("Space ID: %d\nCurrent Name: %s\n", space.Id, space.Name)

	//confirm
	var confirm string
	fmt.Printf("Please enter \"%s\" to confirm delete: ", space.Name)
	fmt.Scanf("%s", &confirm)
	if confirm == space.Name {
		//delete
		err = space.Delete()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Successfully deleted space")
	} else {
		fmt.Println("Space is not deleted")
	}
}
