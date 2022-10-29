package books

import (
	"net/http"

	"github.com/Prosp3r/test_prep/bookstore/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h Handler) GetBooks(c *gin.Context) {
	var books []models.Book

	if result := h.DB.Find(&books); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	c.JSON(http.StatusOK, &books)
}
