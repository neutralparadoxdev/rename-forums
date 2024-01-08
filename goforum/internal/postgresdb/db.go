package postgresdb
import (
	"context"
//	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

type PostgresDatabase struct {
	conn *pgxpool.Pool
}

func New(databaseUrl string) (*PostgresDatabase, error) {
	conn, err := pgxpool.New(context.Background(), databaseUrl)

	if err != nil {
		log.Printf("Unable to connect to database")
		return nil, err
	}
	return &PostgresDatabase{
		conn: conn,
	}, nil
}

func (db *PostgresDatabase) Close() {
	db.conn.Close()
}
