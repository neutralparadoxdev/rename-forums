package postgresdb
import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/neutralparadoxdev/rename-forums/goforum/internal/core"

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


func (db *PostgresDatabase) Init() error {
	return nil
}

func (db *PostgresDatabase) GetUserRepository() core.UserRepository {
	return NewUserRepository(db)
}

func (db *PostgresDatabase) GetCommentRepository() core.CommentRepository {
	return NewCommentRepository(db)
}

func (db *PostgresDatabase) Close() {
	db.conn.Close()
}
