package controller

import (
	"net/http"
	"to-do-list-app/entity"
	"to-do-list-app/helper"
	"to-do-list-app/service"

	"github.com/gin-gonic/gin"
)

type checklistController struct {
	checklistService service.ChecklistService
}

func NewChecklistController(checklistService service.ChecklistService) *checklistController {
	return &checklistController{checklistService}
}

func (h *checklistController) GetAllChecklists(c *gin.Context) {
	checklists, err := h.checklistService.GetAllChecklists()
	if err != nil {
		response := helper.APIResponse("Error to get checklists", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of checklists", http.StatusOK, "success", helper.FormatChecklists(checklists))
	c.JSON(http.StatusOK, response)
}

func (h *checklistController) GetChecklist(c *gin.Context) {
	var input service.GetChecklistDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of checklist", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	checklistDetail, err := h.checklistService.GetChecklistByID(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Failed to get detail of checklist", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Checklist detail", http.StatusOK, "success", helper.FormatChecklistDetail(checklistDetail))
	c.JSON(http.StatusOK, response)
}

func (h *checklistController) CreateChecklist(c *gin.Context) {
	var input service.CreateChecklistInput

	err := c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create checklist", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(entity.User)

	input.User = currentUser

	newChecklist, err := h.checklistService.CreateChecklist(input)
	if err != nil {
		response := helper.APIResponse("Failed to create checklist", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create checklist", http.StatusOK, "success", helper.FormatChecklist(newChecklist))
	c.JSON(http.StatusOK, response)
}

func (h *checklistController) DeleteChecklist(c *gin.Context) {
	var input service.GetChecklistDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to delete checklist", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	deletedChecklist, err := h.checklistService.DeleteChecklist(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Failed to delete checklist", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to delete checklist", http.StatusOK, "success", helper.FormatChecklistDetail(deletedChecklist))
	c.JSON(http.StatusOK, response)
}
