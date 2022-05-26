package repositories

import "server/database"

const user_table = "Users"

type UserCredentials struct {
	Hash string `bson:"hash"`
	Salt string `bson:"sal"`
}

type UserDTO struct {
	database.DTO
	Email       string          `bson:"email"`
	Credentials UserCredentials `bson:"credentials"`
	FirstName   string          `bson:"firstName"`
	LastName    string          `bson:"lastName"`
	PhoneNumber string          `bson:"phoneNumber"`
	Active      bool            `bson:"active"`
	Verified    bool            `bson:"verified"`
}

func NewUserDTO() UserDTO {
	return UserDTO{
		DTO: database.DTO{
			CollectionName: user_table,
		},
	}
}

func GetUserByEmail(email string) (*UserDTO, error) {
	user := NewUserDTO()
	err := database.GetDTOByField(user_table, "email", email, &user)
	if err != nil {
		return nil, err
	}
	return &user, err
}
