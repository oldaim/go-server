package repository

import (
	"context"
	"database/sql"
	"errors"
	"go-server/model"
)

type MemberRepository interface {
	Create(ctx context.Context, member *model.Member) error
	FindById(ctx context.Context, id int64) (*model.Member, error)
	FindAll(ctx context.Context) ([]*model.Member, error)
	Update(ctx context.Context, member *model.Member) error
	Delete(ctx context.Context, id int64) error
}

type MemberRepositoryImpl struct {
	db *sql.DB
}

// NewMemberRepositoryImpl 인스턴스 생성?
func NewMemberRepositoryImpl(db *sql.DB) *MemberRepositoryImpl {
	return &MemberRepositoryImpl{
		db: db,
	}
}

func (r *MemberRepositoryImpl) Create(ctx context.Context, member model.Member) error {

	query := `
		INSERT INTO members(username, password, email, nickname, status, role, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id
    `

	return r.
		db.
		QueryRowContext(ctx, query, member.Username, member.Password, member.Email, member.Nickname, member.Status, member.Role, member.CreatedAt, member.UpdatedAt).
		Scan(&member.ID)
}

// FindById 메서드 구현
func (r *MemberRepositoryImpl) FindById(ctx context.Context, id int64) (*model.Member, error) {
	member := &model.Member{}
	query := `
        SELECT username, email, nickname, status, role, created_at, updated_at
        FROM members
        WHERE id = $1`

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&member.ID,
		&member.Username,
		&member.Email,
		&member.Nickname,
		&member.Status,
		&member.Role,
		&member.CreatedAt,
		&member.UpdatedAt,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return member, nil
}

func (r *MemberRepositoryImpl) FindAll(ctx context.Context) ([]*model.Member, error) {

	query := `SELECT id, username, email, nickname, status, role, created_at, updated_at FROM members`

	rows, err := r.db.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	members := make([]*model.Member, 0)

	for rows.Next() {
		member := &model.Member{}
		err := rows.Scan(
			&member.ID,
			&member.Username,
			&member.Email,
			&member.Nickname,
			&member.Status,
			&member.Role,
			&member.CreatedAt,
			&member.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		members = append(members, member)
	}

	return members, nil
}
