package handler

import (
	"crud-categories/internal/model"
	sampledata "crud-categories/internal/sample_data"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllCategories(c *gin.Context) {

	categories := sampledata.Categories

	response := model.Response{
		Message: "success",
		Data:    categories,
	}

	c.JSON(http.StatusOK, response)
}

func GetCategoryByID(c *gin.Context) {

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Message: "invalid id",
		})
		return
	}

	categories := sampledata.Categories
	for _, v := range categories {
		if v.ID == id {
			c.JSON(http.StatusOK, model.Response{
				Message: "success",
				Data:    v,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, model.Response{
		Message: "category not found",
	})
}

func CreateCategory(c *gin.Context) {
	var category model.Category

	// validate request
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Message: "invalid request",
		})
		return
	}

	// generate id
	lastID := 0
	for _, v := range sampledata.Categories {
		if v.ID > lastID {
			lastID = v.ID
		}
	}
	category.ID = lastID + 1

	// add category to sample data
	sampledata.Categories = append(sampledata.Categories, category)

	c.JSON(http.StatusOK, model.Response{
		Message: "success",
		Data:    category,
	})
}

func UpdateCategory(c *gin.Context) {
	var category model.Category

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Message: "invalid id",
		})
		return
	}

	// validate request
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Message: "invalid request",
		})
		return
	}

	category.ID = id

	// update category
	for i, v := range sampledata.Categories {
		if v.ID == category.ID {
			sampledata.Categories[i] = category
			c.JSON(http.StatusOK, model.Response{
				Message: "success",
				Data:    category,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, model.Response{
		Message: "category not found",
	})
}

func DeleteCategory(c *gin.Context) {
	var category model.Category

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Message: "invalid id",
		})
		return
	}

	category.ID = id

	// delete category
	for i, v := range sampledata.Categories {
		if v.ID == category.ID {
			sampledata.Categories = append(sampledata.Categories[:i], sampledata.Categories[i+1:]...)
			c.JSON(http.StatusOK, model.Response{
				Message: "success",
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, model.Response{
		Message: "category not found",
	})
}
