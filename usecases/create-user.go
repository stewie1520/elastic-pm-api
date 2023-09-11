package usecases

import (
	"context"
	"database/sql"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
	"github.com/stewie1520/elasticpmapi/core"
	"github.com/stewie1520/elasticpmapi/daos/dao_user"
	"github.com/stewie1520/elasticpmapi/tools/types"
)

var _ Command = (*CreateUserCommand)(nil)

type CreateUserCommand struct {
	app core.App
	dao *dao_user.Queries

	ID         string `json:"id"`
	TimeJoined uint64 `json:"timeJoined"`
	Email      string `json:"email"`
	ThirdParty *struct {
		ID     string `json:"id"`
		UserID string `json:"userId"`
	} `json:"thirdParty"`
	TenantIds []string `json:"tenantIds"`
}

func NewCreateUserCommand(app core.App) *CreateUserCommand {
	return &CreateUserCommand{
		app: app,
		dao: app.Dao().User,
	}
}

func (command *CreateUserCommand) Validate() error {
	return validation.ValidateStruct(command,
		validation.Field(&command.ID, validation.Required),
		validation.Field(&command.TimeJoined, validation.Required),
		validation.Field(&command.Email, validation.Required, is.Email, validation.Length(1, 255)),
	)
}

func (command *CreateUserCommand) Execute() error {
	if err := command.Validate(); err != nil {
		return err
	}

	createdAt, err := types.ParseDateTime(command.TimeJoined)
	if err != nil {
		return err
	}

	updatedAt, err := types.ParseDateTime(command.TimeJoined)

	command.dao.CreateUser(context.Background(), dao_user.CreateUserParams{
		ID:        uuid.New(),
		FullName:  sql.NullString{}, // TODO: add fake name
		AccountId: command.ID,
		CreatedAt: createdAt.Time(),
		UpdatedAt: updatedAt.Time(),
	})

	return nil
}
