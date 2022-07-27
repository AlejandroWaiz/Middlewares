package firestore

import (
	"fmt"
	"log"
	"os"

	"github.com/AlejandroWaiz/Middlewares/model"
)

func (f *Firestore) SaveDataIntoDatabase(dataToSave []model.Novel) error {

	for _, novel := range dataToSave {

		log.Println()

		_, _, err := f.client.Collection(os.Getenv("Firestore-Collections-name")).Add(f.ctx, novel)

		if err != nil {
			log.Printf(("Firestore | ERROR) Err: %v"), err)
			return err
		}

		fmt.Println("¡Data saved into database!")

	}

	return nil

}
