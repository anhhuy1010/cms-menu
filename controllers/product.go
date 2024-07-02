package controllers

import (
	"fmt"
	"math"
	"net/http"

	"github.com/anhhuy1010/cms-menu/constant"
	"github.com/anhhuy1010/cms-menu/helpers/respond"
	"github.com/anhhuy1010/cms-menu/helpers/util"

	"github.com/anhhuy1010/cms-menu/models"
	request "github.com/anhhuy1010/cms-menu/request/products"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.mongodb.org/mongo-driver/bson"
)

type ProductController struct {
}

// Function Get List
func (productClt ProductController) List(c *gin.Context) {
	productModel := new(models.Products)
	var req request.GetListRequest
	err := c.ShouldBindWith(&req, binding.Query)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, respond.MissingParams())
		return
	}
	cond := bson.M{}
	if req.Name != nil {
		cond["name"] = req.Name
	}
	if req.IsActive != nil {
		cond["is_active"] = req.IsActive
	}

	if req.MaxQuantity != nil && req.MinQuantity != nil {
		cond["quantity"] = bson.M{"$gte": req.MinQuantity, "$lte": req.MaxQuantity}
	} else if req.MinQuantity != nil {
		cond["quantity"] = bson.M{"$lte": req.MinQuantity}
	} else if req.MaxQuantity != nil {
		cond["quantity"] = bson.M{"$gte": req.MaxQuantity}
	}
	if req.Date != nil {
		cond["start_date"] = bson.M{"$gte": req.Date}
		cond["end_date"] = bson.M{"$lte": req.Date}
	}
	optionsQuery, page, limit := models.GetPagingOption(req.Page, req.Limit, req.Sort)
	var respData []request.ListResponse
	productt, _ := productModel.Pagination(c, cond, optionsQuery)

	for _, productt := range productt {

		res := request.ListResponse{
			Uuid:      productt.Uuid,
			Name:      productt.Name,
			Image:     productt.Image,
			Price:     productt.Price,
			IsActive:  productt.IsActive,
			Sequence:  productt.Sequence,
			StartDate: productt.StartDate,
			EndDate:   productt.EndDate,
			Quantity:  productt.Quantity,
		}
		respData = append(respData, res)
	}
	total, err := productModel.Count(c, cond)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, respond.MissingParams())
		return
	}
	pages := int(math.Ceil(float64(total) / float64(limit)))
	c.JSON(http.StatusOK, respond.SuccessPagination(respData, page, limit, pages, total))
}

// Function Get Detail

func (productClt ProductController) Detail(c *gin.Context) {
	productModel := new(models.Products)
	var reqUri request.GetDetailUri
	err := c.ShouldBindUri(&reqUri)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, respond.MissingParams())
		return
	}

	condition := bson.M{"uuid": reqUri.Uuid}
	productt, err := productModel.FindOne(condition)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusOK, respond.ErrorCommon("Product no found!"))
		return
	}

	response := request.GetDetailResponse{
		Uuid:        productt.Uuid,
		Price:       productt.Price,
		Image:       productt.Image,
		Name:        productt.Name,
		Sequence:    productt.Sequence,
		Quantity:    productt.Quantity,
		Description: productt.Description,
		Gallery:     productt.Gallery,
		IsActive:    productt.IsActive,
		StartDate:   productt.StartDate,
		EndDate:     productt.EndDate,
	}
	c.JSON(http.StatusOK, respond.Success(response, "Successfully"))
}

// Function Update Status
func (productClt ProductController) UpdateStatus(c *gin.Context) {
	productModel := new(models.Products)
	var reqUri request.UpdateUri

	err := c.ShouldBindUri(&reqUri)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, respond.MissingParams())
		return
	}
	var req request.UpdateRequest
	err = c.ShouldBindJSON(&req)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, respond.MissingParams())
		return
	}

	condition := bson.M{"uuid": reqUri.Uuid}
	productt, err := productModel.FindOne(condition)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusOK, respond.ErrorCommon("Product no found!"))
		return
	}
	productt.IsActive = *req.IsActive

	_, err = productt.Update()
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusOK, respond.UpdatedFail())
		return
	}
	c.JSON(http.StatusOK, respond.Success(productt.Uuid, "update successfully"))
}

// Function Delete
func (productClt ProductController) Delete(c *gin.Context) {
	productModel := new(models.Products)
	var reqUri request.DeleteUri
	err := c.ShouldBindUri(&reqUri)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, respond.MissingParams())
		return
	}

	condition := bson.M{"uuid": reqUri.Uuid}
	productt, err := productModel.FindOne(condition)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusOK, respond.ErrorCommon("Product no found!"))
		return
	}

	productt.IsDelete = constant.DELETE

	_, err = productt.Update()
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusOK, respond.UpdatedFail())
		return
	}
	c.JSON(http.StatusOK, respond.Success(productt.Uuid, "Delete successfully"))
}

// Function Create
func (productClt ProductController) Create(c *gin.Context) {
	var req request.GetInsertRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, respond.MissingParams())
		return
	}
	productData := models.Products{}
	productData.Uuid = util.GenerateUUID()
	productData.Name = req.Name
	productData.Image = req.Image
	productData.IsActive = req.IsActive
	productData.Price = req.Price
	productData.Quantity = req.Quantity
	productData.Sequence = req.Sequence
	productData.StartDate = req.StartDate
	productData.EndDate = req.EndDate
	productData.Gallery = req.Gallery
	productData.Description = req.Description
	_, err = productData.Insert()
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusOK, respond.UpdatedFail())
		return
	}
	c.JSON(http.StatusOK, respond.Success(productData.Uuid, "Successfully"))
}

// Function Update
func (productClt ProductController) Update(c *gin.Context) {
	productModel := new(models.Products)
	var reqUri request.UpdateUriDetail
	err := c.ShouldBindUri(&reqUri)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, respond.MissingParams())
		return
	}
	var req request.UpdateDetailRequest
	err = c.ShouldBindJSON(&req)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, respond.MissingParams())
		return
	}

	condition := bson.M{"uuid": reqUri.Uuid}
	productt, err := productModel.FindOne(condition)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusOK, respond.ErrorCommon("Product no found!"))
		return
	}
	if req.IsActive != nil {
		productt.IsActive = *req.IsActive
	}
	if req.Name != "" {
		productt.Name = req.Name
	}
	if req.Image != nil {
		productt.Image = *req.Image
	}
	if req.Description != nil {
		productt.Description = *req.Description
	}
	if req.Price != nil {
		productt.Price = *req.Price
	}
	if req.Sequence != nil {
		productt.Sequence = *req.Sequence
	}
	if req.Quantity != nil {
		productt.Quantity = *req.Quantity
	}
	if req.StartDate != nil {
		productt.StartDate = *req.StartDate
	}
	if req.EndDate != nil {
		productt.EndDate = *req.EndDate
	}

	_, err = productt.Update()

	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusOK, respond.UpdatedFail())
		return
	}
	c.JSON(http.StatusOK, respond.Success(productt.Uuid, "Update successfully"))
}
