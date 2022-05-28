package grid

import (
	"server/database"
	"server/database/repositories"
	"server/server/router/routes"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Test(route *routes.Route, filter bson.M, order bson.M) (*GridResponse, error) {
	var orderOpts *options.FindOptions
	results := make([]repositories.LogMessageDTO, 0)

	if order != nil {
		orderOpts = options.Find().SetSort(order)
	}
	err := database.GetManyDTOs(&results, "LogMessages", filter, orderOpts)
	if err != nil {
		return nil, err
	}

	result := GridResponse{
		Rows:     results,
		RowCount: len(results),
	}
	return &result, nil
}
