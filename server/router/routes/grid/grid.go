package grid

import (
	"net/http"
	"server/globals"
	"server/logger"
	"server/server/router/routes"
)

type GridResponse struct {
	RowCount int   `json:"rowCount"`
	Rows     []any `json:"rows"`
}
type ListHandler func() (*GridResponse, error)

var lists = map[string]ListHandler{
	"test": Test,
}

func Grid(route *routes.Route) {
	args := route.Request.URL.Query()
	listName := args.Get("listName")

	if len(listName) == 0 {
		route.ResponseWriter.WriteError(http.StatusBadRequest, globals.ERROR_INVALID_PAYLOAD)
		return
	}

	handler, ok := lists[listName]
	if !ok {
		route.ResponseWriter.WriteError(http.StatusBadRequest, globals.ERROR_INVALID_LIST_NAME)
		return
	}

	data, err := handler()
	if err != nil {
		go logger.LogError("Error handling list "+listName+" :"+err.Error(), "routes.grid.Grid", globals.ERROR_DATABASE_ERROR)
		return
	}

	route.WriteObject(data)
}
