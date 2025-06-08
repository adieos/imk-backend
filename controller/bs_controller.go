package controller

import (
	"net/http"

	"github.com/adieos/imk-backend/dto"
	"github.com/adieos/imk-backend/service"
	"github.com/adieos/imk-backend/utils"
	"github.com/gin-gonic/gin"
)

type (
	BSController interface {
		CreateBS(ctx *gin.Context)
		GetBSById(ctx *gin.Context)
		GetAllBS(ctx *gin.Context)
		GetAllBSByUserId(ctx *gin.Context)
		UpdateBS(ctx *gin.Context)
		ChangeStatusBS(ctx *gin.Context)
	}

	bSController struct {
		BSService service.BSService
	}
)

func NewBSController(us service.BSService) BSController {
	return &bSController{
		BSService: us,
	}
}

func (c *bSController) CreateBS(ctx *gin.Context) {
	var bs dto.BSCreateRequest
	if err := ctx.ShouldBind(&bs); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	userId := ctx.MustGet("user_id").(string)

	result, err := c.BSService.CreateBS(ctx.Request.Context(), bs, userId)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_BS, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_CREATE_BS, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *bSController) GetBSById(ctx *gin.Context) {
	bsId := ctx.Param("id")
	if bsId == "" {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_PATH, "Bank Sampah ID is required", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.BSService.GetBSById(ctx.Request.Context(), bsId)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_BS_BY_ID, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_BS_BY_ID, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *bSController) GetAllBS(ctx *gin.Context) {
	result, err := c.BSService.GetAllBS(ctx.Request.Context())
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_ALL_BS, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_ALL_BS, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *bSController) GetAllBSByUserId(ctx *gin.Context) {
	userId := ctx.MustGet("user_id").(string)

	result, err := c.BSService.GetAllBSByUserId(ctx.Request.Context(), userId)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_ALL_BS_BY_USER_ID, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_ALL_BS_BY_USER_ID, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *bSController) UpdateBS(ctx *gin.Context) {
	var bs dto.BSUpdateRequest
	if err := ctx.ShouldBind(&bs); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.BSService.UpdateBS(ctx.Request.Context(), bs)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_BS, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATE_BS, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *bSController) ChangeStatusBS(ctx *gin.Context) {
	bsId := ctx.Param("id")
	if bsId == "" {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_PATH, "Bank Sampah ID is required", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	status := ctx.Query("status")
	if status == "" {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_QUERY, "Status is required", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.BSService.ChangeStatusBS(ctx.Request.Context(), bsId, status)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CHANGE_STATUS_BS, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_CHANGE_STATUS_BS, result)
	ctx.JSON(http.StatusOK, res)
}

// func (r *bsRepository) ChangeStatusBS(ctx context.Context, tx *gorm.DB, bsId string, status string) (entity.BankSampah, error) {
