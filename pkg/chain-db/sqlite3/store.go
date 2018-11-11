package sqlite3

import (
	"database/sql"

	pb "github.com/carapace/core/pkg/chain-db/proto"
	"github.com/golang/protobuf/proto"
)

type Engine struct {
	db *sql.DB
}

func (e *Engine) Migrate() error {
	_, err := e.db.Exec(`
	CREATE TABLE chainstore (id INTEGER PRIMARY KEY, time TIMESTAMP DEFAULT CURRENT_TIMESTAMP, key string, data BLOB)
	`)
	return err
}

func (e *Engine) Put(key string, data *pb.Chunk) error {
	serialized, err := proto.Marshal(data)
	if err != nil {
		return err
	}
	_, err = e.db.Exec(`INSERT INTO chainstore (key, data) VALUES (?, ?)`, key, serialized)
	return err
}

func (e *Engine) Get(key string) (chunks []pb.Chunk, err error) {
	rows, err := e.db.Query(`SELECT data FROM chainstore WHERE key = ? ORDER BY ID ASC`, key)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var chunk pb.Chunk
		var data []byte
		err = rows.Scan(&data)
		if err != nil {
			return nil, err
		}

		err = proto.Unmarshal(data, &chunk)
		if err != nil {
			return nil, err
		}

		chunks = append(chunks, chunk)
	}
	return chunks, nil
}

func (e *Engine) Close() error {
	return e.db.Close()
}

func New(db *sql.DB) *Engine {
	return &Engine{db: db}
}
