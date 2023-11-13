package http

import (
	"log"
	// "strings"

	"github.com/gin-gonic/gin"

	"svc-boilerplate-golang/utils/config"
	"svc-boilerplate-golang/utils/message"
)

func (handler *HttpBoilerplateHandler) GetAllEvaluation(ctx *gin.Context) {
	var param = map[string]interface{}{
		"AND": map[string]interface{}{}}
	response, err := handler.boilerplateUsecase.GetAllEvaluation(param)

	if err != nil {
		if err.Error() == config.ERROR_BIND_JSON {
			message.ReturnOk(ctx, make(map[string]interface{}), param)
			return
		}
		message.ReturnInternalServerError(ctx, err.Error())
		log.Println(err.Error())
		return
	}

	message.ReturnOk(ctx, response, param)
}
