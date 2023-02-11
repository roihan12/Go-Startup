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

	campaign, err := ch.campaignService.GetCampaigns(userID)
	if err != nil {
		response := helpers.ApiResponse("Error to get campaigns", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helpers.ApiResponse("Success get list campaigns", http.StatusOK, "success", campaigns.FormatCampaigns(campaign))
	c.JSON(http.StatusOK, response)
	return
}

func (ch *campaignHandler) GetCampaign(c *gin.Context) {

	var input campaigns.GetCampaignDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helpers.ApiResponse("Error to get detail campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	campaign, err := ch.campaignService.GetCampaignByID(input)
	if err != nil {
		response := helpers.ApiResponse("Error to get detail campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.ApiResponse("Success get detail campaign", http.StatusOK, "success", campaign)
	c.JSON(http.StatusOK, response)
	return
}
