package models

import (
	"config"
	"fmt"
)

type Space struct {
	Id          int
	Name        string
	AccessToken string
}

func (space *Space) Save() {
	if space.Id == 0 {
		space.Connection()
	}
	conig.connection.Query()
}
