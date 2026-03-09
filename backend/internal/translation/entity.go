package translation

import "go.mongodb.org/mongo-driver/bson/primitive"

type Translation struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Tag         string             `bson:"tag" json:"tag"`
	Lang        string             `bson:"lang" json:"lang"`
	Translation string             `bson:"translation" json:"translation"`
}
