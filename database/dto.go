package database

import (
	"encoding/hex"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (dto *DTO) Update(saveDTO interface{}, key primitive.ObjectID) (int, error) {
	db, err := GetDatabase()

	if err != nil {
		return 0, xerrors.Errorf("Error retrieving database for upadting "+dto.CollectionName+" "+hex.EncodeToString(dto.ID[:])+": %w", err)
	}

	result, err := db.Collection(dto.CollectionName, nil).UpdateByID(*GetClientContext(), key, saveDTO)
	if err != nil {
		return int(result.ModifiedCount), xerrors.Errorf("Error saving "+dto.CollectionName+": %w", err)
	}
	return int(result.ModifiedCount), nil
}

func GetObjectByID(collectionName string, id primitive.ObjectID, dto interface{}) error {
	return GetDTO(dto, collectionName, bson.M{"_id": id})
}

func GetDTOByField(collectionName string, fieldName string, fieldValue any, dto any) error {
	return GetDTO(dto, collectionName, bson.M{fieldName: fieldValue})
}

func GetDTO(dto any, collectionName string, filter any, opts ...*options.FindOneOptions) error {
	result, err := GetDTOResult(collectionName, filter, opts...)
	if err != nil {
		return err
	}
	err = result.Decode(dto)
	if err != nil {
		return xerrors.Errorf("Error decoding result from collection %s: %w", collectionName, err)
	}
	return nil
}

func GetDTOResult(collectionName string, filter any, opts ...*options.FindOneOptions) (*mongo.SingleResult, error) {
	db, err := GetDatabase()
	if err != nil {
		return nil, err
	}
	result := db.Collection(collectionName).FindOne(*GetClientContext(), filter, opts...)
	if result.Err() != nil {
		return nil, result.Err()
	}
	return result, nil
}
