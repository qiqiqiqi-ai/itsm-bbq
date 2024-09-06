package handler

import (
	"github.com/gin-gonic/gin"
	"itsm/internal/service"
)

type AccidentHandler struct {
	*Handler
	accidentService service.AccidentService
}

func NewAccidentHandler(
	handler *Handler,
	accidentService service.AccidentService,
) *AccidentHandler {
	return &AccidentHandler{
		Handler:         handler,
		accidentService: accidentService,
	}
}

func (h *AccidentHandler) GetAccidentList(ctx *gin.Context) {
	accidentList, err := h.accidentService.GetAccidentList(ctx)
}
