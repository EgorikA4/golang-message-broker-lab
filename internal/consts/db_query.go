package consts

var (
    CREATE_KEYSPACE = `CREATE KEYSPACE IF NOT EXISTS %s WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1};`
    CREATE_TABLE = `CREATE TABLE IF NOT EXISTS %s.logs (id UUID PRIMARY KEY, line TEXT);`
    INSERT_LINE = `INSERT INTO %s.logs (id, line) VALUES (?, ?);`
)
