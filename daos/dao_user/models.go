// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0

package dao_user

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID      `json:"id"`
	FullName  sql.NullString `json:"fullName"`
	AccountId string         `json:"accountId"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
}
