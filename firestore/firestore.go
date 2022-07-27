package firestore

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/AlejandroWaiz/Middlewares/model"
	"google.golang.org/api/option"
)

type Firestore struct {
	client *firestore.Client
	ctx    context.Context
}

type IFirestore interface {
	SaveDataIntoDatabase(dataToSave []model.Novel) error
}

func CreateFirestoreInstance() IFirestore {

	ctx := context.Background()

	client, err := firestore.NewClient(ctx, os.Getenv("Project-id"), option.WithCredentialsFile(os.Getenv("sa-filepath")))

	if err != nil {
		log.Fatalf("(Firestore | ERROR) Err: %v ", err)
		return nil
	}

	return &Firestore{client: client, ctx: ctx}

}
