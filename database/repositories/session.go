package repositories

import (
	"server/database"
	"server/utils"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/xerrors"
)

const session_collection = "Sessions"

type SessionDTO struct {
	database.DTO
	UserEmail   string `bson:"userEmail"`
	SessionCode string `bson:"sessionCode"`
	Active      bool   `bson:"active"`
}

type CheckSessionDto struct {
	Active bool `bson:"active"`
}

func NewSessionDTO() SessionDTO {
	return SessionDTO{
		DTO: database.DTO{
			CollectionName: session_collection,
		},
	}
}

func NewSessionDTOFill(userEmail string) SessionDTO {
	return SessionDTO{
		UserEmail:   userEmail,
		SessionCode: utils.RandomString(8) + "-" + utils.RandomString(8) + "-" + utils.RandomString(8),
		Active:      true,
		DTO: database.DTO{
			CollectionName: session_collection,
		},
	}
}

func DisableUserSessions(userEmail string) error {
	db, err := database.GetDatabase()
	if err != nil {
		return err
	}
	_, err = db.Collection(session_collection).UpdateMany(*database.GetClientContext(),
		bson.D{
			{"userEmail", userEmail},
		},
		bson.D{
			{"$set", bson.D{
				{"active", false},
			}},
		})
	return err
}

func DisableUserSession(sessionCode string) error {
	db, err := database.GetDatabase()
	if err != nil {
		return err
	}
	_, err = db.Collection(session_collection).UpdateOne(*database.GetClientContext(),
		bson.D{
			{"sessionCode", sessionCode},
		},
		bson.D{
			{"$set", bson.D{
				{"active", false},
			}},
		})
	return err
}

func IsActiveSession(sessionCode string) (bool, error) {
	db, err := database.GetDatabase()
	if err != nil {
		return false, err
	}
	projection := options.FindOne().SetProjection(bson.D{
		{"active", 1},
	})
	result := db.Collection(session_collection).FindOne(*database.GetClientContext(), bson.M{"sessionCode": sessionCode}, projection)
	if result.Err() != nil {
		if strings.Contains(result.Err().Error(), "no result") {
			return false, nil
		}
		return false, xerrors.Errorf("Error accessing database for retrieving session with code %s: %w", sessionCode, result.Err())
	}
	dto := CheckSessionDto{}
	err = result.Decode(&dto)
	return dto.Active, err
}
