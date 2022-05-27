package routes

import (
	"net/http"
	"server/database/cache"
	"server/globals"
)

func Reference(route *Route) {
	refTable := route.Request.URL.Query().Get("tableName")
	if len(refTable) == 0 {
		route.ResponseWriter.WriteErrorMessage(http.StatusBadRequest, globals.ERROR_INVALID_PAYLOAD, "Provide 'tableName' arg.")
		return
	}

	refItem := cache.Cache.GetReference(refTable)
	if refItem == nil {
		route.ResponseWriter.WriteErrorMessage(http.StatusBadRequest, globals.ERROR_INVALID_PAYLOAD, "Inexistent reference table: "+refTable+"!")
		return
	}

	route.Response.Data = refItem
	route.WriteData()
}
