package service

import (
	"github.com/AlejandroWaiz/Middlewares/firestore"
)

type Service struct {
	firestore firestore.IFirestore
}

type IService interface {
	SaveDataIntoDatabase()
}

func GetServiceInstance(f firestore.IFirestore) IService {

	return &Service{firestore: f}

}
