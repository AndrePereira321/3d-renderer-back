package database

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/xerrors"
)

type DTO struct {
	ID             primitive.ObjectID  `bson:"_id,omitempty"`
	CollectionName string              `bson:"-"`
	CreatedDate    primitive.Timestamp `bson:"createdDate"`
	UpdatedDate    primitive.Timestamp `bson:"updatedDate"`
}

func (dto *DTO) Save(saveDTO interface{}) (*primitive.ObjectID, error) {
	dto.CreatedDate = primitive.Timestamp{T: uint32(time.Now().Unix())}
	result, err := Database.Collection(dto.CollectionName).InsertOne(*ClientContext, saveDTO)
	if err != nil {
		return nil, xerrors.Errorf("Error saving "+dto.CollectionName+": %w", err)
	}
	if objectID, ok := result.InsertedID.(primitive.ObjectID); ok {
		return &objectID, nil
	}

	return nil, xerrors.New("Invalid generated ID when saving " + dto.CollectionName + " !")
}

func (dto *DTO) Update(saveDTO interface{}, key primitive.ObjectID) (int, error) {
	dto.UpdatedDate = primitive.Timestamp{T: uint32(time.Now().Unix())}
	result, err := Database.Collection(dto.CollectionName).UpdateByID(*ClientContext, key, saveDTO)
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
	result := Database.Collection(collectionName).FindOne(*ClientContext, filter, opts...)
	if result.Err() != nil {
		return nil, result.Err()
	}
	return result, nil
}

func GetManyDTOs(dtos any, collectionName string, filter any, opts ...*options.FindOptions) error {
	cursor, err := GetMany(collectionName, filter, opts...)
	if err != nil {
		return err
	}
	return cursor.All(*ClientContext, dtos)
}

func GetMany(collectionName string, filter any, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	return Database.Collection(collectionName).Find(*ClientContext, filter, opts...)
}

func GetAggregation(dtos any, collectionName string, stages ...bson.D) error {
	cursor, err := Aggregate(collectionName, stages...)
	if err != nil {
		return err
	}
	return cursor.All(*ClientContext, dtos)
}

func Aggregate(collectionName string, stages ...bson.D) (*mongo.Cursor, error) {
	return Database.Collection(collectionName).Aggregate(*ClientContext, stages)
}
