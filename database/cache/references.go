package cache

import (
	"context"
	"server/database"
	"server/globals"
	"server/logger"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/xerrors"
)

const reference_collection = "References"

type ReferenceCache struct {
	RefsByName map[string]ReferenceItemDTO
}

func (cache *ReferenceCache) Init() error {
	refs, err := GetAllReferencesFromDB()
	if err != nil {
		return err
	}
	for i := 0; i < len(refs); i++ {
		cache.RefsByName[refs[i].Table.TableCode] = refs[i]
	}
	return nil
}

func (cache *ReferenceCache) GetReference(tableName string) *ReferenceItemDTO {
	ref, ok := cache.RefsByName[tableName]
	if ok {
		return &ref
	}
	return nil
}

type ReferenceItemDTO struct {
	database.DTO `json:"-"`
	Table        ReferenceTable   `json:"table" bson:"table"`
	Values       []ReferenceValue `json:"values" bson:"values"`
}

type ReferenceTable struct {
	TableCode string  `json:"tableCode" bson:"tableCode"`
	Text1     string  `json:"text1" bson:"text1"`
	Text2     string  `json:"text2" bson:"text2"`
	Text3     string  `json:"text3" bson:"text3"`
	Num1      float32 `json:"num1" bson:"num1"`
	Num2      float32 `json:"num2" bson:"num2"`
	Num3      float32 `json:"num3" bson:"num3"`
	Disabled  bool    `json:"disabled" bson:"disabled"`
}

type ReferenceValue struct {
	Code     string  `json:"code" bson:"code"`
	Text1    string  `json:"text1" bson:"text1"`
	Text2    string  `json:"text2" bson:"text2"`
	Text3    string  `json:"text3" bson:"text3"`
	Text4    string  `json:"text4" bson:"text4"`
	Text5    string  `json:"text5" bson:"text5"`
	Num1     float32 `json:"num1" bson:"num1"`
	Num2     float32 `json:"num2" bson:"num2"`
	Num3     float32 `json:"num3" bson:"num3"`
	Pos      int     `json:"pos" bson:"pos"`
	Disabled bool    `json:"disabled" bson:"disabled"`
}

func GetAllReferencesFromDB() ([]ReferenceItemDTO, error) {
	db, err := database.GetDatabase()
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	results, err := db.Collection(reference_collection).Find(ctx, bson.M{})
	if err != nil {
		go logger.LogError("Error retrieving references from the database!", "cache.references.GetAllReferencesFromDB", globals.ERROR_DATABASE_ERROR)
		return nil, xerrors.Errorf("Error retrieving references from the database: %w", err)
	}
	references := make([]ReferenceItemDTO, 0)
	err = results.All(ctx, &references)
	if err != nil {
		return nil, xerrors.Errorf("Error parsing references from the database: %w", err)
	}
	return references, nil
}
