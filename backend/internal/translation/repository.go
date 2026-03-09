package translation

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	collection *mongo.Collection
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		collection: db.Collection("translations"),
	}
}

func (r *Repository) CreateTranslation(ctx context.Context, translation *Translation) error {
	_, err := r.collection.InsertOne(ctx, translation)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetTranslations(ctx context.Context) ([]Translation, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var translations []Translation
	if err := cursor.All(ctx, &translations); err != nil {
		return nil, err
	}

	return translations, nil
}

func (r *Repository) GetTranslation(ctx context.Context, tag, lang string) (*Translation, error) {
	var translation Translation
	err := r.collection.FindOne(ctx, bson.M{"tag": tag, "lang": lang}).Decode(&translation)
	if err != nil {
		return nil, err
	}

	return &translation, nil
}
