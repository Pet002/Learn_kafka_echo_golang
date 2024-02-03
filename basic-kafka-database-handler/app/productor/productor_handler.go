package productor

import (
	"encoding/json"
	"fmt"

	"github.com/IBM/sarama"
	"github.com/labstack/echo/v4"
)

type producerHandler struct {
	productor sarama.SyncProducer
}

func NewProducerHandler(prod sarama.SyncProducer) *producerHandler {
	return &producerHandler{
		productor: prod,
	}
}

func (p *producerHandler) ProductorHandler(c echo.Context) error {

	user := new(User)
	if err := c.Bind(user); err != nil {
		return c.JSON(400, map[string]string{
			"err": err.Error(),
		})
	}

	userByte, err := json.Marshal(user)
	if err != nil {
		return c.JSON(400, map[string]string{
			"err": err.Error(),
		})
	}
	msg := &sarama.ProducerMessage{
		Topic: "topic1",
		Value: sarama.ByteEncoder(userByte),
	}

	partition, offset, err := p.productor.SendMessage(msg)
	if err != nil {
		return c.JSON(400, map[string]string{
			"error": err.Error(),
		})

	}
	return c.JSON(200, map[string]string{
		"data": fmt.Sprintf("> message sent to partition %d at offset %d\n", partition, offset),
	})
}
