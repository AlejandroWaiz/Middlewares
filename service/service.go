package service

import (
	"github.com/AlejandroWaiz/Middlewares/firestore"
	"github.com/AlejandroWaiz/Middlewares/model"
)

type Service struct {
	firestore firestore.IFirestore
}

type IService interface {
	SaveDataIntoDatabase(dataToSave []model.Novel) error
}

func GetServiceInstance(f firestore.IFirestore) IService {

	return &Service{firestore: f}

}
