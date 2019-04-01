package demo

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mrlyc/gin-gorm-demo/models"
)

// RetrievePairView :
func RetrievePairView(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatus(400)
		return
	}

	result := new(models.Pair)
	err = models.NewPairQuerySet(models.DB).WithoutSoftDelete().IDEq(id).One(result)
	if err != nil {
		ResponseFromError(ctx, err)
		return
	}

	ResponseFromData(ctx, result)
}

// ListPairView :
func ListPairView(ctx *gin.Context) {
	form := new(PaginationForm)
	err := ctx.ShouldBindQuery(form)
	if err != nil {
		ctx.AbortWithStatus(400)
		return
	}

	result := new([]models.Pair)
	err = models.NewPairQuerySet(models.DB).WithoutSoftDelete().Offset(form.Offset).Limit(form.Limit).All(result)
	if err != nil {
		ResponseFromError(ctx, err)
		return
	}

	ResponseFromData(ctx, result)
}

// CreatePairView :
func CreatePairView(ctx *gin.Context) {
	result := new(models.Pair)
	err := ctx.BindJSON(result)
	if err != nil {
		ctx.AbortWithStatus(400)
		return
	}

	err = models.DB.Create(result).Error
	if err != nil {
		ResponseFromError(ctx, err)
		return
	}

	ResponseFromData(ctx, result)
}

// UpdatePairView :
func UpdatePairView(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatus(400)
		return
	}

	result := new(models.Pair)
	err = models.NewPairQuerySet(models.DB).WithoutSoftDelete().IDEq(id).One(result)
	if err != nil {
		ResponseFromError(ctx, err)
		return
	}

	err = ctx.BindJSON(result)
	if err != nil {
		ctx.AbortWithStatus(400)
		return
	}

	result.ID = id
	models.DB.Save(result)
	ResponseFromData(ctx, result)
}

// DeletePairView :
func DeletePairView(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatus(400)
		return
	}

	querySet := models.NewPairQuerySet(models.DB).WithoutSoftDelete().IDEq(id)
	count, err := querySet.Count()
	if err != nil {
		ResponseFromError(ctx, err)
		return
	}

	now := time.Now()
	updater := querySet.GetUpdater()
	updater.SetDeletedAt(&now)
	err = updater.Update()
	if err != nil {
		ResponseFromError(ctx, err)
		return
	}

	ResponseFromData(ctx, count)
}
