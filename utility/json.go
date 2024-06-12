package utility

import (
	"server-chat/domain/response"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func ReadFromJSON(c *gin.Context, data interface{}) error {
	err := c.BindJSON(data)
	return err
}

func ParseRequestBody(c *gin.Context, data interface{}) error {

	var err error

	// body, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	logrus.Errorf("ErrorParsing request body: %v", err)
	// 	// c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to read request body"})
	// 	return err
	// }

	// if len(body) > 0 {
	err = c.BindJSON(data)
	if err != nil {
		logrus.Errorf("ErrorParsing request body: %v", err)
		return err
	}
	// }

	return err
}

func WriteResponseListJSON(c *gin.Context, data response.GlobalListDataTableResponse, err error) {
	var code int = 200
	var status string = "OK"
	var message string = "success"

	if err != nil {
		if badRequest, ok := err.(*BadRequestError); ok {
			code = 400
			status = "BadRequest"
			message = badRequest.Message
		} else if notFound, ok := err.(*BadRequestError); ok {
			code = 404
			status = "NotFound"
			message = notFound.Message
		} else {
			code = 500
			status = "ERROR"
			message = err.Error()
		}
	} else {

	}
	response := response.GlobalListResponse{
		Code:    code,
		Status:  status,
		Message: message,
		Data:    data,
	}

	c.JSON(code, response)
}

func WriteResponseSingleJSON(c *gin.Context, data interface{}, err error) {
	var code int = 200
	var status string = "OK"
	var message string = "success"

	if err != nil {
		if badRequest, ok := err.(*BadRequestError); ok {
			code = 400
			status = "BadRequest"
			message = badRequest.Message
		} else {
			code = 500
			status = "ERROR"
			message = err.Error()
		}
	} else {

	}
	response := response.GlobalSingleResponse{
		Code:    code,
		Status:  status,
		Message: message,
		Data:    data,
	}

	c.JSON(code, response)
}
