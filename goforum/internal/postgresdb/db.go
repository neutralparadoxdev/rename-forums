package postgresdb
import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/neutralparadoxdev/rename-forums/goforum/internal/core"

	"log"
)

type PostgresDatabase struct {
	pool *pgxpool.Pool
}

func New(databaseUrl string) (*PostgresDatabase, error) {
	pool, err := pgxpool.New(context.Background(), databaseUrl)

	if err != nil {
		log.Printf("Unable to connect to database")
		return nil, err
	}
	return &PostgresDatabase{
		pool: pool,
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

func (db *PostgresDatabase) GetForumRepository() core.ForumRepository {
	return NewForumRepository(db)
}

func (db *PostgresDatabase) GetPostRepository() core.PostRepository {
	return NewPostRepository(db)
}

func (db *PostgresDatabase) GetVoteRepository() core.VoteRepository {
	return NewVoteRepository(db)
}

func (db *PostgresDatabase) Close() {
	db.pool.Close()
}

func (db *PostgresDatabase) GetSessionRepository() core.SessionRepository {
	return NewSessionRepository(db)
}
