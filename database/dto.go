package database

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/xerrors"
)

type DTO struct {
	ID          primitive.ObjectID  `bson:"_id,omitempty"`
	Name        string              `bson:"name"`
	UpdatedDate primitive.Timestamp `bson:"updatedDate"`
}

func (dto *DTO) Save(saveDTO interface{}) (*primitive.ObjectID, error) {
	db, err := GetDatabase()
	if err != nil {
		return nil, xerrors.Errorf("Error retrieving database for saving "+dto.Name+": %w", err)
	}

	result, err := db.Collection(dto.Name, nil).InsertOne(clientContext(), saveDTO)
	if err != nil {
		return nil, xerrors.Errorf("Error saving "+dto.Name+": %w", err)
	}
	if objectID, ok := result.InsertedID.(primitive.ObjectID); ok {
		return &objectID, nil
	}

	return nil, xerrors.New("Invalid generated ID when saving " + dto.Name + " !")
}
