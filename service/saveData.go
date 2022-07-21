package service

import (
	"fmt"

	"github.com/AlejandroWaiz/Middlewares/model"
)

func (s *Service) SaveDataIntoDatabase(dataToSave model.Novel) error {

	err := s.firestore.SaveDataIntoDatabase(dataToSave)

	if err != nil {

		fmt.Printf("Can not save data into database: %v", err)

		return err

	}

	return nil

}
