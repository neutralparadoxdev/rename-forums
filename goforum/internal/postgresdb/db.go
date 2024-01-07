package postgresdb
import (
	"context"
	"github.com/jackc/pgx/v5"
	"log"
)

type PostgresDatabase struct {
	conn *pgx.Conn
}

func New(databaseUrl string) (*PostgresDatabase, error) {
	conn, err := pgx.Connect(context.Background(), databaseUrl)

	if err != nil {
		log.Printf("Unable to connect to database")
		return nil, err
	}
	return &PostgresDatabase{
		conn: conn,
	}, nil
}

func (db *PostgresDatabase) Close() {
	db.conn.Close(context.Background())
}
