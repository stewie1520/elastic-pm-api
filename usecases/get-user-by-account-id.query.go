package usecases

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/stewie1520/elasticpmapi/core"
	"github.com/stewie1520/elasticpmapi/daos/dao_user"
)

var _ Query[dao_user.User] = (*GetUserByAccountIDQuery)(nil)

func NewGetUserByAccountIDQuery(app core.App) *GetUserByAccountIDQuery {
	return &GetUserByAccountIDQuery{
		app: app,
		dao: app.Dao().User,
	}
}

type GetUserByAccountIDQuery struct {
	app core.App
	dao *dao_user.Queries

	AccountID string `json:"accountId"`
}

// Execute implements Query.
func (q *GetUserByAccountIDQuery) Execute() (dao_user.User, error) {
	if err := q.Validate(); err != nil {
		return dao_user.User{}, err
	}

	user, err := q.dao.GetUserByAccountID(context.Background(), q.AccountID)
	return user, err
}

// Validate implements Query.
func (q *GetUserByAccountIDQuery) Validate() error {
	return validation.ValidateStruct(q,
		validation.Field(&q.AccountID, validation.Required, is.UUIDv4),
	)
}
