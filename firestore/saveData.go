package firestore

import (
	"fmt"
	"log"
	"os"

	"github.com/AlejandroWaiz/Middlewares/model"
)

func (f *Firestore) SaveDataIntoDatabase(dataToSave model.Novel) error {

	_, _, err := f.client.Collection(os.Getenv("Firestore-Collections-name")).Add(f.ctx, dataToSave)

	if err != nil {
		log.Printf(("Firestore | ERROR) Err: %v"), err)
		return err
	}

	fmt.Println("¡Data saved into database!")

	return nil

}
