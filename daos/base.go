package daos

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/stewie1520/elasticpmapi/daos/dao_user"
)

type Dao struct {
	db *sql.DB

	Builder squirrel.StatementBuilderType
	User    *dao_user.Queries
}

func New(db *sql.DB) *Dao {
	return &Dao{
		db:      db,
		User:    dao_user.New(db),
		Builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}

func (dao *Dao) DB() *sql.DB {
	return dao.db
}
