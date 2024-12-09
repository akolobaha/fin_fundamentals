package transport

import (
	"fin_fundamentals/internal/config"
	"fin_fundamentals/internal/entity"
	"fin_fundamentals/internal/log"
	"fmt"
	"github.com/streadway/amqp"
)

type Rabbitmq struct {
	Chan  *amqp.Channel
	Queue amqp.Queue
}

func New() *Rabbitmq {
	return &Rabbitmq{}
}

func (rabbit *Rabbitmq) InitConn(cfg *config.Config) {
	// Установите соединение с RabbitMQ
	conn, err := amqp.Dial(cfg.GetRabbitDSN())
	if err != nil {
		log.Error("Failed to connect to RabbitMQ: ", err)
	}

	// Создайте канал
	ch, err := conn.Channel()
	rabbit.Chan = ch

	if err != nil {
		log.Error("Failed to open a channel: ", err)
	}
}

func (rabbit *Rabbitmq) ConnClose() {
	rabbit.Chan.Close()
}

func (rabbit *Rabbitmq) DeclareQueue(name string) {
	// Объявите очередь, в которую будете отправлять сообщения
	queue, err := rabbit.Chan.QueueDeclare(
		name,  // имя очереди
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // аргументы
	)
	rabbit.Queue = queue
	if err != nil {
		log.Error("Failed to declare a queue: ", err)
	}
}

func (rabbit *Rabbitmq) SendMsg(data []byte, header entity.FundamentalHeader) {
	err := rabbit.Chan.Publish(
		"",                // обменник
		rabbit.Queue.Name, // ключ маршрутизации (имя очереди)
		false,             // обязательное
		false,             // немедленное
		amqp.Publishing{
			DeliveryMode: amqp.Persistent, // сохранять сообщение
			ContentType:  "text/json",
			Body:         data,
			Headers: amqp.Table{
				"Period":       header.Period,
				"Ticker":       header.Ticker,
				"Report":       header.Report,
				"ReportMethod": header.ReportMethod,
				"ReportUrl":    header.ReportUrl,
				"SourceUrl":    header.SourceUrl,
			},
		})
	if err != nil {
		log.Error("Failed to publish a message: ", err)
	} else {
		log.Info(fmt.Sprintf("Message sent to RabbitMQ %s %s %s", header.Ticker, header.Report, header.ReportMethod))
	}

}
