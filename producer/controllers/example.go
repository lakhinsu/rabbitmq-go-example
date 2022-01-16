package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lakhinsu/rabbitmq-go-example/producer/consts"
	"github.com/lakhinsu/rabbitmq-go-example/producer/models"
	"github.com/lakhinsu/rabbitmq-go-example/producer/utils"
	"github.com/rs/zerolog/log"
)

func Example(c *gin.Context) {
	var msg models.Message

	request_id := c.GetString("x-request-id")

	// Bind request payload with our model
	if binderr := c.ShouldBindJSON(&msg); binderr != nil {

		log.Error().Err(binderr).Str("request_id", request_id).
			Msg("Error occurred while binding request data")

		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": binderr.Error(),
		})
		return
	}

	connectionString := utils.GetEnvVar("RMQ_URL")

	rmqProducer := utils.RMQProducer{
		consts.EXAMPLE_QUEUE,
		connectionString,
	}

	rmqProducer.PublishMessage("text/plain", []byte(msg.Message))

	c.JSON(http.StatusOK, gin.H{
		"response": "Message received",
	})

}
