package controllers

import (
	// "fmt"
	"fmt"
	"math"
	"net/http"

	// "github.com/anhhuy1010/cms-menu/constant"
	"github.com/anhhuy1010/cms-menu/constant"
	"github.com/anhhuy1010/cms-menu/helpers/respond"
	"github.com/anhhuy1010/cms-menu/helpers/util"

	// "github.com/anhhuy1010/cms-menu/helpers/util"
	"github.com/anhhuy1010/cms-menu/models"
	request "github.com/anhhuy1010/cms-menu/request/user"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.mongodb.org/mongo-driver/bson"
)

type ProductController struct {
}

// khởi tạo hàm get List
func (productClt ProductController) List(c *gin.Context) {
	productModel := new(models.Product)
	var req request.GetListRequest

	// kiểm tra đầu vào
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
	if req.Quantity != nil {
		cond["quantity"] = bson.M{"$gt": req.Quantity}
	}
	if req.Quantity != nil {
		cond["quantity"] = bson.M{"$lt": req.Quantity}
	}
	if req.StartDate != nil {
		cond["start_date"] = req.StartDate
	}
	if req.EndDate != nil {
		cond["end_date"] = req.EndDate
	}
	optionsQuery, page, limit := models.GetPagingOption(req.Page, req.Limit, req.Sort)
	var respData []request.ListResponse
	productt, _ := productModel.Pagination(c, cond, optionsQuery)

	for _, productt := range productt {

		res := request.ListResponse{
			Uuid:       productt.Uuid,
			Name:       productt.Name,
			Price:      productt.Price,
			ClientUuid: productt.ClientUuid,
			IsActive:   productt.IsActive,
			Image:      productt.Image,
			Sequence:   productt.Sequence,
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

// hàm get detail list
func (productClt ProductController) Detail(c *gin.Context) {
	productModel := new(models.Product)
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
		c.JSON(http.StatusOK, respond.ErrorCommon("User no found!"))
		return
	}

	response := request.GetDetailResponse{
		Name:  productt.Name,
		Image: productt.Image,
		Price: productt.Price,
	}
	c.JSON(http.StatusOK, respond.Success(response, "Successfully"))
}

// hàm update list
func (productClt ProductController) Update(c *gin.Context) {
	productModel := new(models.Product)
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
	if req.IsActive != nil {
		productt.IsActive = *req.IsActive
	}
	if req.IsDelete != nil {
		productt.IsDelete = *req.IsDelete
	}

	_, err = productt.Update()
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusOK, respond.UpdatedFail())
		return
	}
	c.JSON(http.StatusOK, respond.Success(productt.IsActive, "update successfully"))
}

// hàm delete list
func (productClt ProductController) Delete(c *gin.Context) {
	productModel := new(models.Product)
	var reqUri request.DeleteUri
	err := c.ShouldBindUri(&reqUri)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, respond.MissingParams())
		return
	}

	condition := bson.M{"uuid": reqUri.Uuid}
	user, err := productModel.FindOne(condition)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusOK, respond.ErrorCommon("User no found!"))
		return
	}

	user.IsDelete = constant.DELETE

	_, err = user.Update()
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusOK, respond.UpdatedFail())
		return
	}
	c.JSON(http.StatusOK, respond.Success(user.Uuid, "Delete successfully"))
}

// hàm create list
func (productClt ProductController) Create(c *gin.Context) {
	var req request.GetInsertRequest
	err := c.ShouldBindWith(&req, binding.Query)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, respond.MissingParams())
		return
	}
	err = c.ShouldBindJSON(&req)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, respond.MissingParams())
		return
	}
	productData := models.Product{}
	productData.Uuid = util.GenerateUUID()
	productData.Name = req.Name
	productData.Image = req.Image
	productData.IsActive = req.IsActive
	productData.Price = req.Price
	productData.Quantity = req.Quantity
	productData.Sequence = req.Sequence
	productData.StartDate = req.StartDate
	productData.EndDate = req.EndDate
	_, err = productData.Insert()
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusOK, respond.UpdatedFail())
		return
	}
	c.JSON(http.StatusOK, respond.Success(productData.Uuid, "update successfully"))
}

// hàm create detail
func (productClt ProductController) CreateDetail(c *gin.Context) {
	var req request.GetCreateRequest
	err := c.ShouldBindWith(&req, binding.Query)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, respond.MissingParams())
		return
	}
	err = c.ShouldBindJSON(&req)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, respond.MissingParams())
		return
	}
	productData := models.Product{}
	productData.Uuid = util.GenerateUUID()
	productData.Name = req.Name
	productData.Image = req.Image
	productData.IsActive = req.IsActive
	productData.Price = req.Price
	productData.Quantity = req.Quantity
	productData.Sequence = req.Sequence
	productData.Description = req.Description
	productData.Gallery = req.Gallery
	_, err = productData.Insert()
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusOK, respond.UpdatedFail())
		return
	}
	c.JSON(http.StatusOK, respond.Success(productData.Uuid, "update successfully"))
}

func (productClt ProductController) UpdateDetail(c *gin.Context) {
	productModel := new(models.Product)
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
		c.JSON(http.StatusOK, respond.ErrorCommon("User no found!"))
		return
	}
	if req.IsActive != nil {
		productt.IsActive = *req.IsActive
	}
	if req.IsDelete != nil {
		productt.IsDelete = *req.IsDelete
	}
	if req.Name != "" {
		productt.Name = req.Name
	}
	if req.Image != "" {
		productt.Image = req.Image
	}
	if req.Description != "" {
		productt.Description = req.Description
	}
	if req.Price != 0 {
		productt.Price = req.Price
	}
	if req.Sequence != 0 {
		productt.Sequence = req.Sequence
	}
	if req.Quantity != 0 {
		productt.Quantity = req.Quantity
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
	c.JSON(http.StatusOK, respond.Success(productt.IsActive, "update successfully"))
}
