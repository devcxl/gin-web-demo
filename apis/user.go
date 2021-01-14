package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	. "main/common"
	. "main/models"
	"net/http"
)

func AddUserApi(c *gin.Context) {
	u := User{}
	err := c.BindJSON(u)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}
	u.Password = GetSha256Code([]byte(u.Password))
	ra, err := u.AddUser()
	if err != nil {
		return
	}

	msg := fmt.Sprintf("insert successful %d", ra)
	c.JSON(http.StatusOK, Result{
		Message: msg,
	})

}
