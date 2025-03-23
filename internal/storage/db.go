package storage

import (
	"fmt"
	"log"
	"os"

	"github.com/EgorikA4/golang-message-broker-lab/config"
	"github.com/EgorikA4/golang-message-broker-lab/internal/consts"
	"github.com/gocql/gocql"
)

var cluster *gocql.ClusterConfig

func InitDB() error {
    cfg := config.GetConfig()
    cluster = gocql.NewCluster(cfg.CassandraHost)

    cluster.Consistency = gocql.Quorum
    cluster.ProtoVersion = 4

    session, err := cluster.CreateSession()
	if err != nil {
        return err
	}
	defer session.Close()

    keyspace := os.Getenv("KEYSPACE")
	keyspaceQuery := fmt.Sprintf(consts.CREATE_KEYSPACE, keyspace)
	if err := session.Query(keyspaceQuery).Exec(); err != nil {
        return err
	}
    log.Printf("the keyspace: '%s' has been created.\n", keyspace)

    createTableQuery := fmt.Sprintf(consts.CREATE_TABLE, keyspace)
    if err := session.Query(createTableQuery).Exec(); err != nil {
        return err
	}
    log.Println("the table 'logs' has been created.")
    return nil
}

func GetCluster() *gocql.ClusterConfig {
    return cluster
}
