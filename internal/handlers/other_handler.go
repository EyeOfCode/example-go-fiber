package handlers

import (
	"go-fiber-api/internal/service"
	"go-fiber-api/pkg/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type OtherHandler struct {
	otherService *service.ArtworkApiService
}

func NewOtherHandler(otherService *service.ArtworkApiService) *OtherHandler {
	return &OtherHandler{
		otherService: otherService,
	}
}

// @Summary Get Gallery endpoint
// @Description Get the API's get Gallery
// @Tags other
// @Accept json
// @Produce json
// @Router /other/example/gallery [get]
func (o *OtherHandler) GetListImages(c *fiber.Ctx) error {
	res, err := o.otherService.GetListImages(c.Context())
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, err.Error())
	}
	if res == nil {
		return utils.SendError(c, http.StatusNotFound, "Failed to find gallery")
	}

	return utils.SendSuccess(c, http.StatusOK, res)
}
