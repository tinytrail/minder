// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: session_store.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createSessionState = `-- name: CreateSessionState :one
INSERT INTO session_store (provider, grp_id, port, session_state, owner_filter) VALUES ($1, $2, $3, $4, $5) RETURNING id, provider, grp_id, port, owner_filter, session_state, created_at
`

type CreateSessionStateParams struct {
	Provider     uuid.UUID      `json:"provider"`
	GrpID        sql.NullInt32  `json:"grp_id"`
	Port         sql.NullInt32  `json:"port"`
	SessionState string         `json:"session_state"`
	OwnerFilter  sql.NullString `json:"owner_filter"`
}

func (q *Queries) CreateSessionState(ctx context.Context, arg CreateSessionStateParams) (SessionStore, error) {
	row := q.db.QueryRowContext(ctx, createSessionState,
		arg.Provider,
		arg.GrpID,
		arg.Port,
		arg.SessionState,
		arg.OwnerFilter,
	)
	var i SessionStore
	err := row.Scan(
		&i.ID,
		&i.Provider,
		&i.GrpID,
		&i.Port,
		&i.OwnerFilter,
		&i.SessionState,
		&i.CreatedAt,
	)
	return i, err
}

const deleteExpiredSessionStates = `-- name: DeleteExpiredSessionStates :exec
DELETE FROM session_store WHERE created_at < NOW() - INTERVAL '1 day'
`

func (q *Queries) DeleteExpiredSessionStates(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteExpiredSessionStates)
	return err
}

const deleteSessionState = `-- name: DeleteSessionState :exec
DELETE FROM session_store WHERE id = $1
`

func (q *Queries) DeleteSessionState(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteSessionState, id)
	return err
}

const deleteSessionStateByGroupID = `-- name: DeleteSessionStateByGroupID :exec
DELETE FROM session_store WHERE provider=$1 AND grp_id = $2
`

type DeleteSessionStateByGroupIDParams struct {
	Provider uuid.UUID     `json:"provider"`
	GrpID    sql.NullInt32 `json:"grp_id"`
}

func (q *Queries) DeleteSessionStateByGroupID(ctx context.Context, arg DeleteSessionStateByGroupIDParams) error {
	_, err := q.db.ExecContext(ctx, deleteSessionStateByGroupID, arg.Provider, arg.GrpID)
	return err
}

const getGroupIDPortBySessionState = `-- name: GetGroupIDPortBySessionState :one
SELECT provider, grp_id, port, owner_filter FROM session_store WHERE session_state = $1
`

type GetGroupIDPortBySessionStateRow struct {
	Provider    uuid.UUID      `json:"provider"`
	GrpID       sql.NullInt32  `json:"grp_id"`
	Port        sql.NullInt32  `json:"port"`
	OwnerFilter sql.NullString `json:"owner_filter"`
}

func (q *Queries) GetGroupIDPortBySessionState(ctx context.Context, sessionState string) (GetGroupIDPortBySessionStateRow, error) {
	row := q.db.QueryRowContext(ctx, getGroupIDPortBySessionState, sessionState)
	var i GetGroupIDPortBySessionStateRow
	err := row.Scan(
		&i.Provider,
		&i.GrpID,
		&i.Port,
		&i.OwnerFilter,
	)
	return i, err
}

const getSessionState = `-- name: GetSessionState :one
SELECT id, provider, grp_id, port, owner_filter, session_state, created_at FROM session_store WHERE id = $1
`

func (q *Queries) GetSessionState(ctx context.Context, id int32) (SessionStore, error) {
	row := q.db.QueryRowContext(ctx, getSessionState, id)
	var i SessionStore
	err := row.Scan(
		&i.ID,
		&i.Provider,
		&i.GrpID,
		&i.Port,
		&i.OwnerFilter,
		&i.SessionState,
		&i.CreatedAt,
	)
	return i, err
}

const getSessionStateByGroupID = `-- name: GetSessionStateByGroupID :one
SELECT id, provider, grp_id, port, owner_filter, session_state, created_at FROM session_store WHERE grp_id = $1
`

func (q *Queries) GetSessionStateByGroupID(ctx context.Context, grpID sql.NullInt32) (SessionStore, error) {
	row := q.db.QueryRowContext(ctx, getSessionStateByGroupID, grpID)
	var i SessionStore
	err := row.Scan(
		&i.ID,
		&i.Provider,
		&i.GrpID,
		&i.Port,
		&i.OwnerFilter,
		&i.SessionState,
		&i.CreatedAt,
	)
	return i, err
}
