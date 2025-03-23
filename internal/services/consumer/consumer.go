package consumer

import (
	"fmt"
	"log"
	"os"

	"github.com/EgorikA4/golang-message-broker-lab/config"
	"github.com/EgorikA4/golang-message-broker-lab/internal/consts"
	"github.com/EgorikA4/golang-message-broker-lab/internal/storage"
	"github.com/gocql/gocql"

	amqp "github.com/rabbitmq/amqp091-go"
)

func Listen() error {
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
    if err != nil {
        return err
    }

    q, err := ch.QueueDeclare(
            "",
            false,
            false,
            true,
            false,
            nil,
    )
    if err != nil {
        return err
    }

    err = ch.QueueBind(
        q.Name,
        "",
        cfg.ExchangeName,
        false,
        nil,
    )
    if err != nil {
        return err
    }

    msgs, err := ch.Consume(
        q.Name,
        "",
        true,
        false,
        false,
        false,
        nil,
    )
    if err != nil {
        return err
    }

    cluster := storage.GetCluster()
    session, err := cluster.CreateSession()
    if err != nil {
        return err
    }
    defer session.Close()

    keyspace := os.Getenv("KEYSPACE")
    for msg := range msgs {
        insertQuery := fmt.Sprintf(consts.INSERT_LINE, keyspace)
        line := string(msg.Body)
        if err := session.Query(insertQuery, gocql.TimeUUID(), line).Exec(); err != nil {
            return err
        }
        log.Printf("the line: %s was successfully inserted into DB\n", line)
    }
    return nil
}
