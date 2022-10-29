package books

import (
	"net/http"

	"github.com/Prosp3r/test_prep/bookstore/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h Handler) GetBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book

	if result := h.DB.Find(&book, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	c.JSON(http.StatusOK, &book)
}
