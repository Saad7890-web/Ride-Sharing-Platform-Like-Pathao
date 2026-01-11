package repository

import (
	"context"
	"errors"

	"ride/internal/auth/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)


type PostgresUserRepository struct {
	db *pgxpool.Pool
}

func NewPostgresUserRepository(db *pgxpool.Pool) UserRepository {
	return &PostgresUserRepository{db: db}
}


func (r *PostgresUserRepository) Create(user *domain.User) error {
	query := `
		INSERT INTO users (email, password_hash, is_active)
		VALUES ($1, $2, $3)
	`
	_, err := r.db.Exec(context.Background(), query,
		user.Email,
		user.PasswordHash,
		user.IsActive,
	)
	return err
}

func (r *PostgresUserRepository) FindByEmail(email string) (*domain.User, error) {
	query := `
	SELECT id, email, password_hash, is_active
	FROM users
	WHERE email = $1
	`

	row := r.db.QueryRow(context.Background(), query, email)

	var user domain.User
	if err := row.Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.IsActive,
	); err != nil {
		return nil, errors.New("user not found")
	}

	
	user.Roles = r.loadRoles(user.ID)

	return &user, nil
}


func (r *PostgresUserRepository) loadRoles(userID string) []domain.Role {
	query := `
	SELECT r.id, r.name, p.id, p.name
	FROM roles r
	JOIN user_roles ur ON ur.role_id = r.id
	JOIN role_permissions rp ON rp.role_id = r.id
	JOIN permissions p ON p.id = rp.permission_id
	WHERE ur.user_id = $1
	`

	rows, err := r.db.Query(context.Background(), query, userID)
	if err != nil {
		return nil
	}
	defer rows.Close()

	roleMap := make(map[int]*domain.Role)

	for rows.Next() {
		var (
			roleID   int
			roleName string
			permID   int
			permName string
		)

		_ = rows.Scan(&roleID, &roleName, &permID, &permName)

		if _, ok := roleMap[roleID]; !ok {
			roleMap[roleID] = &domain.Role{
				ID:   roleID,
				Name: roleName,
			}
		}

		roleMap[roleID].Permissions = append(
			roleMap[roleID].Permissions,
			domain.Permission{
				ID:   permID,
				Name: permName,
			},
		)
	}

	var roles []domain.Role
	for _, r := range roleMap {
		roles = append(roles, *r)
	}

	return roles
}

func (r *PostgresUserRepository) FindByID(id string) (*domain.User, error) {
		const query = `
	SELECT
		u.id,
		u.email,
		r.id,
		r.name,
		p.id,
		p.name
	FROM users u
	JOIN user_roles ur ON ur.user_id = u.id
	JOIN roles r ON r.id = ur.role_id
	JOIN role_permissions rp ON rp.role_id = r.id
	JOIN permissions p ON p.id = rp.permission_id
	WHERE u.id = $1
	`

	rows, err := r.db.Query(context.Background(), query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	user := &domain.User{
		ID:    id,
		Roles: []domain.Role{},
	}

	roleMap := make(map[int]*domain.Role)

	for rows.Next() {
		var (
			email    string
			roleID   int
			roleName string
			permID   int
			permName string
		)

		if err := rows.Scan(
			&user.ID,
			&email,
			&roleID,
			&roleName,
			&permID,
			&permName,
		); err != nil {
			return nil, err
		}

		user.Email = email

		role, exists := roleMap[roleID]
		if !exists {
			role = &domain.Role{
				ID:          roleID,
				Name:        roleName,
				Permissions: []domain.Permission{},
			}
			roleMap[roleID] = role
		}

		role.Permissions = append(role.Permissions, domain.Permission{
			ID:   permID,
			Name: permName,
		})
	}

	
	for _, role := range roleMap {
		user.Roles = append(user.Roles, *role)
	}

	if len(user.Roles) == 0 {
		return nil, errors.New("user not found or has no roles")
	}

	return user, nil
}

func (r *PostgresUserRepository) AssignRole(userID string, roleID int) error {
	query := `
		INSERT INTO user_roles (user_id, role_id)
		VALUES ($1, $2)
	`
	_, err := r.db.Exec(context.Background(), query, userID, roleID)
	return err
}