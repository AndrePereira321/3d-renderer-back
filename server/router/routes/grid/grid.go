package grid

import (
	"net/http"
	"server/globals"
	"server/logger"
	"server/server/router/routes"

	"go.mongodb.org/mongo-driver/bson"
)

type GridPayload struct {
	GridName string `json:"gridName"`
	Filter   bson.M `json:"filter"`
	Order    bson.M `json:"order"`
	StartRow int    `json:"startRow"`
	EndRow   int    `json:"endRow"`
	Search   string `json:"search"`
}

type GridResponse struct {
	RowCount int `json:"rowCount"`
	Rows     any `json:"rows"`
}
type ListHandler func(route *routes.Route, filter bson.M, order bson.M) (*GridResponse, error)

const argName = "gridName"

var grids = map[string]ListHandler{
	"test": Test,
}

func Grid(route *routes.Route) {
	payload := GridPayload{}
	err := route.Request.UnmarshallBody(&payload)
	if err != nil {
		route.ResponseWriter.WriteError(http.StatusBadRequest, globals.ERROR_INVALID_PAYLOAD)
		return
	}

	handler, ok := grids[payload.GridName]
	if !ok {
		route.ResponseWriter.WriteError(http.StatusBadRequest, globals.ERROR_INVALID_LIST_NAME)
		return
	}

	data, err := handler(route, payload.Filter, payload.Order)
	if err != nil {
		go logger.LogError("Error handling list "+payload.GridName+" :"+err.Error(), "routes.grid.Grid", globals.ERROR_DATABASE_ERROR)
		return
	}

	route.WriteObject(data)
}
