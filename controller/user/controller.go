package user

import (
	"apigorm/model"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	Model model.UserModel
}

func (uc UserController) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		tmp := uc.Model.GetAll()

		if tmp == nil {
			return c.JSON(http.StatusInternalServerError, "error from database")
		}

		res := map[string]interface{}{
			"message": "Get All Data",
			"data":    tmp,
		}

		return c.JSON(http.StatusOK, res)
	}
}

func (uc *UserController) GetSpesificUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := c.Param("id")
		cnv, err := strconv.Atoi(param)
		if err != nil {
			log.Println("Cannot convert to int", err.Error())
			return c.JSON(http.StatusInternalServerError, "Cannot convert id")
		}

		data := uc.Model.GetSpecific(cnv)

		if data.Id == 0 {
			return c.JSON(http.StatusBadRequest, "no data")
		}

		res := map[string]interface{}{
			"message": "Get all data",
			"data":    data,
		}

		return c.JSON(http.StatusOK, res)
	}
}

func (uc UserController) InsertUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmp model.User
		err := c.Bind(&tmp)
		if err != nil {
			log.Println("Cannot parse input to object", err.Error())
			return c.JSON(http.StatusInternalServerError, "Error from server")

		}

		data := uc.Model.Insert(tmp)

		if data.Id == 0 {
			return c.JSON(http.StatusInternalServerError, "Error from server")
		}

		res := map[string]interface{}{
			"message": "Succes input data",
			"data":    data,
		}

		return c.JSON(http.StatusOK, res)
	}
}

func (uc UserController) UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		qry := map[string]interface{}{}
		param := c.Param("id")
		cnv, err := strconv.Atoi(param)
		if err != nil {
			log.Println("Cannot concert to int", err.Error())
			return c.JSON(http.StatusInternalServerError, "cannot convert id")
		}

		var tmp model.User
		err = c.Bind(&tmp)
		if err != nil {
			log.Println("Cannot parse input to object", err.Error())
			return c.JSON(http.StatusInternalServerError, "Error from Server")
		}

		if tmp.Name != "" {
			qry["name"] = tmp.Name
		}
		if tmp.Email != "" {
			qry["email"] = tmp.Email
		}
		if tmp.Password != "" {
			qry["password"] = tmp.Password
		}

		data := uc.Model.Update(cnv, tmp)

		if data.Id == 0 {
			return c.JSON(http.StatusInternalServerError, "cannot update data")
		}

		res := map[string]interface{}{
			"message": "Succes update data",
			"data":    data,
		}

		return c.JSON(http.StatusOK, res)
	}
}

func (uc UserController) DeleteData() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := c.Param("id")
		cnv, err := strconv.Atoi(param)
		if err != nil {
			log.Println("Cannot convert to int", err.Error())
			return c.JSON(http.StatusInternalServerError, "cannot convert id")
		}

		if !uc.Model.Delete(cnv) {
			return c.JSON(http.StatusInternalServerError, "cannot delete")
		}

		res := map[string]interface{}{
			"message": "Succes delete data",
		}

		return c.JSON(http.StatusOK, res)
	}
}
