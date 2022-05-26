package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/xerrors"
)

type DTO struct {
	ID             primitive.ObjectID  `bson:"_id,omitempty"`
	CollectionName string              `bson:"-"`
	UpdatedDate    primitive.Timestamp `bson:"updatedDate"`
}

func (dto *DTO) Save(saveDTO interface{}) (*primitive.ObjectID, error) {
	db, err := GetDatabase()

	if err != nil {
		return nil, xerrors.Errorf("Error retrieving database for saving "+dto.CollectionName+": %w", err)
	}

	result, err := db.Collection(dto.CollectionName, nil).InsertOne(*GetClientContext(), saveDTO)
	if err != nil {
		return nil, xerrors.Errorf("Error saving "+dto.CollectionName+": %w", err)
	}
	if objectID, ok := result.InsertedID.(primitive.ObjectID); ok {
		return &objectID, nil
	}

	return nil, xerrors.New("Invalid generated ID when saving " + dto.CollectionName + " !")
}

func GetObjectByID(collectionName string, id primitive.ObjectID, dto interface{}) error {
	db, err := GetDatabase()
	if err != nil {
		return xerrors.Errorf("Error retrieving database for loading: %w", err)
	}
	result := db.Collection(collectionName).FindOne(context.Background(), bson.M{"_id": id})
	if result.Err() != nil {
		return xerrors.Errorf("Error accessing collection %s for finding object by it's id [%s]: %w", collectionName, id, result.Err())
	}
	err = result.Decode(dto)
	if err != nil {
		return xerrors.Errorf("Error decoding result from collection %s: %w", collectionName, err)
	}
	return nil
}

func GetDTOByField(collectionName string, fieldName string, fieldValue interface{}, dto interface{}) error {
	db, err := GetDatabase()
	if err != nil {
		return xerrors.Errorf("Error retrieving database for loading: %w", err)
	}
	result := db.Collection(collectionName).FindOne(context.Background(), bson.M{fieldName: fieldValue})
	if result.Err() != nil {
		return xerrors.Errorf("Error accessing collection %s for finding object with the field %s containging %s: %w", collectionName, fieldName, fieldValue, result.Err())
	}
	err = result.Decode(dto)
	if err != nil {
		return xerrors.Errorf("Error decoding result from collection %s: %w", collectionName, err)
	}
	return nil
}
