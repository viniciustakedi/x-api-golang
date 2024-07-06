package responseJson

import "github.com/gin-gonic/gin"

type ResponseData struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

type ResponseText struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func Error(
	c *gin.Context,
	statusCode int,
	message string,
	err error,
) {
	c.JSON(
		statusCode,
		ResponseData{
			Status:  statusCode,
			Message: message,
			Data: map[string]interface{}{
				"data": err.Error(),
			},
		},
	)
}

func Data(
	c *gin.Context,
	statusCode int,
	message string,
	data map[string]interface{},
) {
	c.JSON(
		statusCode,
		ResponseData{
			Status:  statusCode,
			Message: message,
			Data:    data,
		},
	)
}

func Text(
	c *gin.Context,
	statusCode int,
	message string,
) {
	c.JSON(
		statusCode,
		ResponseData{
			Status:  statusCode,
			Message: message,
		},
	)
}
