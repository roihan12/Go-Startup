package handlers

import (
	"bwastartup/campaigns"
	"bwastartup/helpers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	campaignService campaigns.Service
}

func NewCampaignHandler(campaignService campaigns.Service) *campaignHandler {
	return &campaignHandler{campaignService}
}

func (ch *campaignHandler) GetCampaigns(c *gin.Context) {

	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := ch.campaignService.GetCampaigns(userID)
	if err != nil {
		response := helpers.ApiResponse("Error to get campaigns", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helpers.ApiResponse("Success get list campaigns", http.StatusOK, "success", campaigns)
	c.JSON(http.StatusOK, response)
	return
}
