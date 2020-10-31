package spaceModel

import (
	"cloudstorageapi.com/configs"
	"errors"
	"github.com/lib/pq"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

type Space struct {
	Id          int
	Name        string
	AccessToken string
}

func (space *Space) Save() error {
	err := configs.Connection.
		QueryRow("INSERT INTO spaces(space_name) values($1) Returning id, access_key", space.Name).
		Scan(&space.Id, &space.AccessToken)
	if err != nil {
		log.Fatal(err)
		if err, ok := err.(*pq.Error); ok && err.Code.Name() == "unique_violation" {
			return errors.New("Space name already exists. Please choose a different space name")
		} else {
			return errors.New("Something went wrong while creating new space.")
		}
	}

	err = space.createSpaceInFileSystem()
	if err != nil {
		log.Fatal(err)
		space.Delete()
		return errors.New("Failed to create new folder")
	}
	return nil
}

func (space *Space) UpdateName(newName string) error {
	_, err := configs.Connection.Exec("UPDATE spaces SET space_name=$1 WHERE id=$2", newName, space.Id)
	if err != nil {
		log.Fatal(err)
		return errors.New("Something went wrong")
	}
	space.Name = newName
	return nil
}

func (space *Space) Delete() error {
	err := space.removeSpaceInFileSystem()
	if err != nil {
		log.Fatal(err)
		return errors.New("Couldn't remove directory")
	}

	_, err = configs.Connection.Exec("DELETE FROM spaces WHERE id=$1", space.Id)
	if err != nil {
		log.Fatal(err)
		return errors.New("Something went wrong")
	}
	return nil
}

func FindSpaceById(spaceId int) (*Space, error) {
	space := Space{}
	err := configs.Connection.
		QueryRow("SELECT * FROM spaces WHERE id = $1", spaceId).
		Scan(&space.Id, &space.Name, &space.AccessToken)

	if err != nil {
		log.Fatal(err)
		if _, ok := err.(*pq.Error); ok {
			return nil, errors.New("Some error occured")
		}
		return nil, nil
	}
	return &space, nil
}

func All() ([]Space, error) {
	spaces := make([]Space, 0, 10)
	rows, err := configs.Connection.
		Query("SELECT id, space_name from spaces ORDER BY space_name")
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("Something went wrong")
	}
	for rows.Next() {
		space := Space{}
		err := rows.Scan(&space.Id, &space.Name)
		if err != nil {
			return nil, errors.New("Some error occured while parsing data")
		}
		spaces = append(spaces, space)
	}
	return spaces, nil
}

func (space *Space) createSpaceInFileSystem() error {
	return os.MkdirAll(space.GetFilePath(), os.ModePerm)
}

func (space *Space) removeSpaceInFileSystem() error {
	return os.RemoveAll(space.GetFilePath())
}

func (space *Space) GetFilePath() string {
	return filepath.Join(configs.STORAGE_ROOT_PATH, strconv.Itoa(space.Id))
}
