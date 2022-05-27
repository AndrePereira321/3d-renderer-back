package routes

import (
	"net/http"
	"server/database/cache"
	"server/globals"
)

func Reference(route *Route) {
	args := route.Request.URL.Query()
	refTable := args.Get("tableName")
	itemCode := args.Get("code")

	if len(refTable) == 0 {
		route.ResponseWriter.WriteErrorMessage(http.StatusBadRequest, globals.ERROR_INVALID_PAYLOAD, "Provide 'tableName' arg.")
		return
	}

	if len(itemCode) == 0 {
		refItem := cache.Cache.GetReference(refTable)
		if refItem == nil {
			route.ResponseWriter.WriteErrorMessage(http.StatusBadRequest, globals.ERROR_INVALID_PAYLOAD, "Inexistent reference table: "+refTable+"!")
			return
		}

		route.Response.Data = refItem
	} else {
		item := cache.Cache.References.GetReferenceValue(refTable, itemCode)
		if item == nil {
			route.ResponseWriter.WriteErrorMessage(http.StatusBadRequest, globals.ERROR_INVALID_PAYLOAD, "Inexistent item with table: "+refTable+" and code "+itemCode+"!")
			return
		}
		route.Response.Data = item
	}
	route.WriteData()
}
