package cache

import (
	"server/database"
	"server/globals"
	"server/logger"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/xerrors"
)

const reference_collection = "References"

type ReferenceCache struct {
	RefsByName          map[string]*ReferenceItemDTO
	RefsCodeByTableName map[string]map[string]*ReferenceValue
}

func (cache *ReferenceCache) Init() error {
	refs, err := GetAllReferencesFromDB()
	if err != nil {
		return err
	}
	for i := 0; i < len(refs); i++ {
		tableCode := refs[i].Table.TableCode
		cache.RefsByName[tableCode] = &refs[i]
		cache.RefsCodeByTableName[tableCode] = map[string]*ReferenceValue{}
		for j := 0; j < len(refs[i].Values); j++ {
			cache.RefsCodeByTableName[tableCode][refs[i].Values[j].Code] = &refs[i].Values[j]
		}
	}
	return nil
}

func (cache *ReferenceCache) GetReference(tableName string) *ReferenceItemDTO {
	ref, ok := cache.RefsByName[tableName]
	if ok {
		return ref
	}
	return nil
}

func (cache *ReferenceCache) GetReferenceValue(tableName string, code string) *ReferenceValue {
	table, ok := cache.RefsCodeByTableName[tableName]
	if !ok {
		return nil
	}
	value := table[code]
	if !ok {
		return nil
	}
	return value
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
	results, err := database.Database.Collection(reference_collection).Find(*database.ClientContext, bson.M{})
	if err != nil {
		go logger.LogError("Error retrieving references from the database!", "cache.references.GetAllReferencesFromDB", globals.ERROR_DATABASE_ERROR)
		return nil, xerrors.Errorf("Error retrieving references from the database: %w", err)
	}
	references := make([]ReferenceItemDTO, results.RemainingBatchLength())
	err = results.All(*database.ClientContext, &references)
	if err != nil {
		return nil, xerrors.Errorf("Error parsing references from the database: %w", err)
	}
	return references, nil
}
