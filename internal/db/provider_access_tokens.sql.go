// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: provider_access_tokens.sql

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createAccessToken = `-- name: CreateAccessToken :one
INSERT INTO provider_access_tokens (group_id, provider_id, encrypted_token, expiration_time, owner_filter) VALUES ($1, $2, $3, $4, $5) RETURNING id, provider_id, group_id, owner_filter, encrypted_token, expiration_time, created_at, updated_at
`

type CreateAccessTokenParams struct {
	GroupID        int32          `json:"group_id"`
	ProviderID     uuid.UUID      `json:"provider_id"`
	EncryptedToken string         `json:"encrypted_token"`
	ExpirationTime time.Time      `json:"expiration_time"`
	OwnerFilter    sql.NullString `json:"owner_filter"`
}

func (q *Queries) CreateAccessToken(ctx context.Context, arg CreateAccessTokenParams) (ProviderAccessToken, error) {
	row := q.db.QueryRowContext(ctx, createAccessToken,
		arg.GroupID,
		arg.ProviderID,
		arg.EncryptedToken,
		arg.ExpirationTime,
		arg.OwnerFilter,
	)
	var i ProviderAccessToken
	err := row.Scan(
		&i.ID,
		&i.ProviderID,
		&i.GroupID,
		&i.OwnerFilter,
		&i.EncryptedToken,
		&i.ExpirationTime,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteAccessToken = `-- name: DeleteAccessToken :exec
DELETE FROM provider_access_tokens WHERE provider_id = $1 AND group_id = $2
`

type DeleteAccessTokenParams struct {
	ProviderID uuid.UUID `json:"provider_id"`
	GroupID    int32     `json:"group_id"`
}

func (q *Queries) DeleteAccessToken(ctx context.Context, arg DeleteAccessTokenParams) error {
	_, err := q.db.ExecContext(ctx, deleteAccessToken, arg.ProviderID, arg.GroupID)
	return err
}

const getAccessTokenByGroupID = `-- name: GetAccessTokenByGroupID :one
SELECT id, provider_id, group_id, owner_filter, encrypted_token, expiration_time, created_at, updated_at FROM provider_access_tokens WHERE provider_id = $1 AND group_id = $2
`

type GetAccessTokenByGroupIDParams struct {
	ProviderID uuid.UUID `json:"provider_id"`
	GroupID    int32     `json:"group_id"`
}

func (q *Queries) GetAccessTokenByGroupID(ctx context.Context, arg GetAccessTokenByGroupIDParams) (ProviderAccessToken, error) {
	row := q.db.QueryRowContext(ctx, getAccessTokenByGroupID, arg.ProviderID, arg.GroupID)
	var i ProviderAccessToken
	err := row.Scan(
		&i.ID,
		&i.ProviderID,
		&i.GroupID,
		&i.OwnerFilter,
		&i.EncryptedToken,
		&i.ExpirationTime,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getAccessTokenByProvider = `-- name: GetAccessTokenByProvider :many
SELECT id, provider_id, group_id, owner_filter, encrypted_token, expiration_time, created_at, updated_at FROM provider_access_tokens WHERE provider_id = $1
`

func (q *Queries) GetAccessTokenByProvider(ctx context.Context, providerID uuid.UUID) ([]ProviderAccessToken, error) {
	rows, err := q.db.QueryContext(ctx, getAccessTokenByProvider, providerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ProviderAccessToken{}
	for rows.Next() {
		var i ProviderAccessToken
		if err := rows.Scan(
			&i.ID,
			&i.ProviderID,
			&i.GroupID,
			&i.OwnerFilter,
			&i.EncryptedToken,
			&i.ExpirationTime,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAccessTokenSinceDate = `-- name: GetAccessTokenSinceDate :one
SELECT id, provider_id, group_id, owner_filter, encrypted_token, expiration_time, created_at, updated_at FROM provider_access_tokens WHERE provider_id = $1 AND group_id = $2 AND created_at >= $3
`

type GetAccessTokenSinceDateParams struct {
	ProviderID uuid.UUID `json:"provider_id"`
	GroupID    int32     `json:"group_id"`
	CreatedAt  time.Time `json:"created_at"`
}

func (q *Queries) GetAccessTokenSinceDate(ctx context.Context, arg GetAccessTokenSinceDateParams) (ProviderAccessToken, error) {
	row := q.db.QueryRowContext(ctx, getAccessTokenSinceDate, arg.ProviderID, arg.GroupID, arg.CreatedAt)
	var i ProviderAccessToken
	err := row.Scan(
		&i.ID,
		&i.ProviderID,
		&i.GroupID,
		&i.OwnerFilter,
		&i.EncryptedToken,
		&i.ExpirationTime,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateAccessToken = `-- name: UpdateAccessToken :one
UPDATE provider_access_tokens SET encrypted_token = $3, expiration_time = $4, owner_filter = $5, updated_at = NOW() WHERE provider_id = $1 AND group_id = $2 RETURNING id, provider_id, group_id, owner_filter, encrypted_token, expiration_time, created_at, updated_at
`

type UpdateAccessTokenParams struct {
	ProviderID     uuid.UUID      `json:"provider_id"`
	GroupID        int32          `json:"group_id"`
	EncryptedToken string         `json:"encrypted_token"`
	ExpirationTime time.Time      `json:"expiration_time"`
	OwnerFilter    sql.NullString `json:"owner_filter"`
}

func (q *Queries) UpdateAccessToken(ctx context.Context, arg UpdateAccessTokenParams) (ProviderAccessToken, error) {
	row := q.db.QueryRowContext(ctx, updateAccessToken,
		arg.ProviderID,
		arg.GroupID,
		arg.EncryptedToken,
		arg.ExpirationTime,
		arg.OwnerFilter,
	)
	var i ProviderAccessToken
	err := row.Scan(
		&i.ID,
		&i.ProviderID,
		&i.GroupID,
		&i.OwnerFilter,
		&i.EncryptedToken,
		&i.ExpirationTime,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
