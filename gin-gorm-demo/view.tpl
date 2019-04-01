package demo

import "github.com/cheekybits/genny/generic"

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mrlyc/gin-gorm-demo/models"
)

type TYPE generic.Type

// RetrieveTYPEView :
func RetrieveTYPEView(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatus(400)
		return
	}

	result := new(models.TYPE)
	err = models.NewTYPEQuerySet(models.DB).WithoutSoftDelete().IDEq(id).One(result)
	if err != nil {
		ResponseFromError(ctx, err)
		return
	}

	ResponseFromData(ctx, result)
}

// ListTYPEView :
func ListTYPEView(ctx *gin.Context) {
	form := new(PaginationForm)
	err := ctx.ShouldBindQuery(form)
	if err != nil {
		ctx.AbortWithStatus(400)
		return
	}

	result := new([]models.TYPE)
	err = models.NewTYPEQuerySet(models.DB).WithoutSoftDelete().Offset(form.Offset).Limit(form.Limit).All(result)
	if err != nil {
		ResponseFromError(ctx, err)
		return
	}

	ResponseFromData(ctx, result)
}

// CreateTYPEView :
func CreateTYPEView(ctx *gin.Context) {
	result := new(models.TYPE)
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

// UpdateTYPEView :
func UpdateTYPEView(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatus(400)
		return
	}

	result := new(models.TYPE)
	err = models.NewTYPEQuerySet(models.DB).WithoutSoftDelete().IDEq(id).One(result)
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

// DeleteTYPEView :
func DeleteTYPEView(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatus(400)
		return
	}

	querySet := models.NewTYPEQuerySet(models.DB).WithoutSoftDelete().IDEq(id)
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
