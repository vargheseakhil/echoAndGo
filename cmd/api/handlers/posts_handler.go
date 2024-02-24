package handlers

import (
	"echoGoApp/cmd/api/service"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func PostIndexHandler(c echo.Context) error {
	data, err := service.GetAll()
	fmt.Println(err)
	if err != nil {
		c.String(http.StatusBadGateway, "Unable to process")
	}
	res := make(map[string]any)
	res["status"] = "ok"
	res["data"] = data
	return c.JSON(http.StatusOK, res)
}

func PostSingleHandler(c echo.Context) error {
	id := c.Param("id")
	idx, err := strconv.Atoi(id)
	if err != nil {
		return c.String(http.StatusBadGateway, "Unable to process the id")
	}
	data, err := service.GetById(idx)
	if err != nil {
		return c.String(http.StatusBadGateway, "Unable to fetch data by ID")
	}
	res := make(map[string]any)
	res["status"] = "ok"
	res["data"] = data
	return c.JSON(http.StatusOK, res)
}
