package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	. "main/common"
	. "main/models"
	"net/http"
	"strconv"
)

func IndexApi(c *gin.Context) {
	c.JSON(http.StatusOK, Result{
		Data:    "",
		Message: "测试",
	})
}

/*
添加Persion
*/
func AddPersonApi(c *gin.Context) {
	firstName := c.Request.FormValue("first_name")
	lastName := c.Request.FormValue("last_name")
	p := Person{FirstName: firstName, LastName: lastName}

	ra, err := p.AddPerson()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("insert successful %d", ra)
	c.JSON(http.StatusOK, gin.H{
		"data": true,
		"msg":  msg,
	})
}

/*
修改Persion
*/
func ModPersonApi(c *gin.Context) {
	firstName := c.Request.FormValue("first_name")
	lastName := c.Request.FormValue("last_name")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatalln(err)
	}
	p := Person{Id: id}

	p.GetPerson()
	if p.FirstName != "" {
		p.FirstName = firstName
		p.LastName = lastName
		ra, err := p.ModPerson()
		if err != nil {
			log.Fatalln(err)
		}
		msg := fmt.Sprintf("update successful %d", ra)
		c.JSON(http.StatusOK, gin.H{
			"data": true,
			"msg":  msg,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": nil,
			"msg":  "Person not found",
		})
	}
}

func DelPersonApi(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	p := Person{Id: id}

	ra, err := p.DelPerson()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("delete successful %d", ra)
	c.JSON(http.StatusOK, gin.H{
		"data": true,
		"msg":  msg,
	})
}

func GetPersonApi(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatalln(err)
	}
	person := Person{Id: id}

	person.GetPerson()
	if person.FirstName != "" {
		c.JSON(http.StatusOK, gin.H{
			"data": person,
			"msg":  "success",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": nil,
			"msg":  "Person not found",
		})
	}

}

func GetPersonsApi(c *gin.Context) {
	var p Person
	persons, err := p.GetPersons()
	if err != nil {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": persons,
		"msg":  "success",
	})
}

func GetPersonAsNameApi(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	name := c.Param("name")
	if err != nil {
		log.Fatalln(err)
	}
	person := Person{Id: id, FirstName: name}

	person.GetPerson()
	if person.FirstName != "" {
		c.JSON(http.StatusOK, gin.H{
			"data": person,
			"msg":  "success",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": nil,
			"msg":  "Person not found",
		})
	}

}
