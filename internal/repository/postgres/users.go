package postgres

import (
	"context"
	"fmt"

	"final/internal/repository/postgres/sqlc"
)

func (p *Postgres) CreateUser(ctx context.Context, arg sqlc.CreateUserParams) (int64, error) {
	p.logger.Debug("CreateUser called", "login", arg.Login, "status", arg.Status)

	queries := sqlc.New(p.connPool)

	userID, err := queries.CreateUser(ctx, &arg)
	if err != nil {
		p.logger.Error("executing CreateUser", err)
		return 0, fmt.Errorf("executing CreateUser: %w", err)
	}

	return userID, nil
}

func (p *Postgres) GetUserByLogin(ctx context.Context, login string) (*sqlc.User, error) {
	queries := sqlc.New(p.connPool)

	user, err := queries.GetUserByLogin(ctx, login)
	if err != nil {
		p.logger.Error("executing GetUserByLogin", err)
		return nil, fmt.Errorf("executing GetUserByLogin: %w", err)
	}

	return user, nil
}

func (p *Postgres) GetUserByID(ctx context.Context, id int64) (*sqlc.User, error) {
	queries := sqlc.New(p.connPool)

	user, err := queries.GetUserByID(ctx, id)
	if err != nil {
		p.logger.Error("executing GetUserByID", err)
		return nil, fmt.Errorf("executing GetUserByID: %w", err)
	}

	return user, nil
}

func (p *Postgres) UpdateUserStatus(ctx context.Context, arg sqlc.UpdateUserStatusParams) error {
	queries := sqlc.New(p.connPool)

	if err := queries.UpdateUserStatus(ctx, &arg); err != nil {
		p.logger.Error("executing UpdateUserStatus", err)
		return fmt.Errorf("executing UpdateUserStatus: %w", err)
	}

	return nil
}

func (p *Postgres) UpdateUserName(ctx context.Context, arg sqlc.UpdateUserNameParams) error {
	queries := sqlc.New(p.connPool)

	if err := queries.UpdateUserName(ctx, &arg); err != nil {
		p.logger.Error("executing UpdateUserName", err)
		return fmt.Errorf("executing UpdateUserName: %w", err)
	}

	return nil
}

func (p *Postgres) DeleteUser(ctx context.Context, id int64) error {
	queries := sqlc.New(p.connPool)

	if err := queries.DeleteUser(ctx, id); err != nil {
		p.logger.Error("executing DeleteUser", err)
		return fmt.Errorf("executing DeleteUser: %w", err)
	}

	return nil
}

func (p *Postgres) Logout(ctx context.Context, id int64) error {
	queries := sqlc.New(p.connPool)

	if err := queries.Logout(ctx, id); err != nil {
		p.logger.Error("executing Logout", err)
		return fmt.Errorf("executing Logout: %w", err)
	}

	return nil
}
