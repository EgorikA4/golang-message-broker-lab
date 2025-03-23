package producer

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/EgorikA4/golang-message-broker-lab/config"

	amqp "github.com/rabbitmq/amqp091-go"
)

func ProcessFile(filePath string) error {
    cfg := config.GetConfig()
    url := fmt.Sprintf("amqp://%s:%s@%s:%s/", cfg.RabbitUser, cfg.RabbitPassword, cfg.RabbitHost, cfg.RabbitPort)
    conn, err := amqp.Dial(url)
    if err != nil {
        return err
    }
    defer conn.Close()

    ch, err := conn.Channel()
    if err != nil {
        return err
    }
    defer ch.Close()

    err = ch.ExchangeDeclare(
        cfg.ExchangeName,
        "fanout",
        true,
        false,
        false,
        false,
        nil,
    )

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    file, err := os.Open(filePath)
    if err != nil {
        return err
    }

    fileScanner := bufio.NewScanner(file)
    for fileScanner.Scan() {
        line := fileScanner.Text()
        err = ch.PublishWithContext(
            ctx,
            cfg.ExchangeName,
            "",
            false,
            false,
            amqp.Publishing{
                ContentType: "text/plain",
                Body:        []byte(line),
        })
        if err != nil {
            return err
        }
        log.Printf("the line: %s was successfully sent.\n", line)
    }
    return nil
}
