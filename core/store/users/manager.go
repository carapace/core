package user

import (
	"context"
	"database/sql"
	"github.com/carapace/core/api/v0/proto"
	"time"
)

type Manager struct {
}

func New() *Manager {
	return &Manager{}
}

// Creates a new user.
//
// Will return an error if the user already exists
func (m *Manager) Create(ctx context.Context, tx *sql.Tx, user v0.User) error {
	_, err := tx.ExecContext(ctx, `
				INSERT INTO users (
					name, 
					email, 
					primary_public_key, 
					recovery_public_key, 
					super_user, 
					auth_level, 
					weight,
					user_set
					) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		user.Name,
		user.Email,
		user.PrimaryPublicKey,
		user.RecoveryPublicKey,
		user.SuperUser,
		user.AuthLevel,
		user.Weight,
		user.Set,
	)
	return err
}

// Alter a user by deleting the field.
//
// Will return an ErrUserDoesNotExists if the user does not exist.
func (m *Manager) Alter(ctx context.Context, tx *sql.Tx, user v0.User) error {
	// This is quite a hacky delete, but NULL fields are not evaluated in indexes; thus there can be only one field
	// with deleted=FALSE, and many fields with deleted=NULL (thus true).
	err := m.Delete(ctx, tx, user)
	if err != nil {
		return err
	}
	return m.Create(ctx, tx, user)
}

// Get a user by publicKey (either primary or recovery will work)
//
// Will return an ErrUserDoesNotExists if the user does not exist.
func (m *Manager) Get(ctx context.Context, tx *sql.Tx, publicKey []byte) (*v0.User, error) {
	row := tx.QueryRowContext(ctx, `
				SELECT  
					name, 
					email, 
					primary_public_key, 
					recovery_public_key, 
					super_user, 
					auth_level, 
					weight,
					user_set 
					FROM users
				WHERE primary_public_key = ? OR recovery_public_key = ?
				AND deleted = FALSE`, publicKey, publicKey,
	)
	user := v0.User{}
	err := row.Scan(
		&user.Name,
		&user.Email,
		&user.PrimaryPublicKey,
		&user.RecoveryPublicKey,
		&user.SuperUser,
		&user.AuthLevel,
		&user.Weight,
		&user.Set,
	)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, ErrUserDoesNotExists
		}
		return nil, err
	}
	return &user, nil
}

// Delete a user (soft delete)
//
// Will return an ErrUserDoesNotExists if the user does not exist.
func (m *Manager) Delete(ctx context.Context, tx *sql.Tx, user v0.User) error {
	result, err := tx.ExecContext(ctx, `
				UPDATE users 
				SET deleted_at = ?, 
				deleted = TRUE 
				WHERE primary_public_key = ?;`,
		time.Now(),
		user.PrimaryPublicKey,
	)
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return ErrUserDoesNotExists
	}
	return nil
}

func (m *Manager) BySet(ctx context.Context, tx *sql.Tx, set string) ([]*v0.User, error) {
	rows, err := tx.QueryContext(ctx, `
				SELECT  
					name, 
					email, 
					primary_public_key, 
					recovery_public_key, 
					super_user, 
					auth_level, 
					weight,
					user_set 
					FROM users
				WHERE user_set = (SELECT ID FROM user_sets WHERE name = ?)
				AND deleted = FALSE`, set)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := []*v0.User{}
	for rows.Next() {
		user := v0.User{}
		err := rows.Scan(
			&user.Name,
			&user.Email,
			&user.PrimaryPublicKey,
			&user.RecoveryPublicKey,
			&user.SuperUser,
			&user.AuthLevel,
			&user.Weight,
			&user.Set,
		)
		if err != nil {
			return nil, err
		}
		res = append(res, &user)
	}
	return res, nil
}
