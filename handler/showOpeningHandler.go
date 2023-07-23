package handler

import (
	"api-opportunities/schemas"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowOpeningHandler(ctx *gin.Context) {
	id  := ctx.Query("id")

	if id == ""{
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id","queryParameter").Error())
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil{
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
	}

	sendSuccess(ctx, "showing-opening", opening)
}
