package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	us "service/genproto/user_service"
	"service/storage"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type userRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) storage.UserRepoI {
	return &userRepo{
		db: db,
	}
}

func (u *userRepo) Create(ctx context.Context, req *us.CreateUs) (*us.Us, error) {
	// resp := &us.Us{}
	id := uuid.NewString()

	_, err := u.db.Exec(ctx, `
		INSERT INTO system_user (
			id,
			phone,
			gmail,
			name,
			role
		) VALUES (
			$1, $2, $3, $4, $5
		) `, id, req.Phone, req.Gmail, req.Name, req.Role)

	if err != nil {
		log.Println("error while creating user in storage", err)
		return nil, err
	}

	user, err := u.GetByID(ctx, &us.UsPrimaryKey{Id: id})
	if err != nil {
		log.Println("error while getting user by id", err)
		return nil, err
	}
	return user, nil
}

func (u *userRepo) GetByID(ctx context.Context, req *us.UsPrimaryKey) (*us.Us, error) {
	resp := &us.Us{}

	var (
		created_at sql.NullString
		updated_at sql.NullString
	)
	err := u.db.QueryRow(ctx, `
	        SELECT id,
	        phone,
			gmail,
			name,
			role,
	        created_at,
	        updated_at
	        FROM system_user
	    WHERE id=$1`, req.Id).Scan(&resp.Id, &resp.Phone, &resp.Gmail, &resp.Name, &resp.Role,
		&created_at, &updated_at)

	if err != nil {
		log.Println("error while getting user by id", err)
		return nil, err
	}

	resp.CreatedAt = created_at.String
	resp.UpdatedAt = updated_at.String

	return resp, nil
}

func (u *userRepo) GetList(ctx context.Context, req *us.GetListUsRequest) (resp *us.GetListUsResponse, err error) {
	resp = &us.GetListUsResponse{}

	var (
		filter     string
		created_at sql.NullString
		updated_at sql.NullString
	)

	offset := (req.Page - 1) * req.Limit

	if req.Search != "" {
		filter = fmt.Sprintf(` AND (name ILIKE '%%%v%%' OR phone ILIKE '%%%v%%' OR email ILIKE '%%%v%%')`, req.Search, req.Search, req.Search)
	}

	filter += fmt.Sprintf(" OFFSET %v LIMIT %v", offset, req.Limit)

	rows, err := u.db.Query(ctx, `
		SELECT
			id,
			phone,
			gmail,
			name,
			role,
	        created_at,
	        updated_at
			FROM system_user WHERE deleted_at=0
			ORDER BY created_at DESC
			`+filter)

	if err != nil {
		log.Println("error while getting all users", err)
		return nil, err
	}
	defer rows.Close()

	var count int64 = 0

	for rows.Next() {
		var user us.Us
		count++
		err = rows.Scan(&user.Id, &user.Phone, &user.Gmail, &user.Name, &user.Role,
			&created_at, &updated_at)

		if err != nil {
			log.Println("error while scanning users", err)
			return nil, err
		}

		user.CreatedAt = created_at.String
		user.UpdatedAt = updated_at.String

		resp.Users = append(resp.Users, &user)
	}

	resp.Count = count

	return resp, nil
}

func (u *userRepo) Update(ctx context.Context, req *us.UpdateUs) (resp *us.Us, err error) {
	_, err = u.db.Exec(ctx, `
        UPDATE system_user SET
            phone = $2,
            gmail = $3,
            name = $4,
            role = $5,
            updated_at = NOW()
        WHERE id = $1`, req.Id, req.Phone, req.Gmail, req.Name, req.Role)

	if err != nil {
		log.Println("error while updating user in storage", err)
		return nil, err
	}

	user, err := u.GetByID(ctx, &us.UsPrimaryKey{Id: req.Id})
	if err != nil {
		log.Println("error while getting updated user by id", err)
		return nil, err
	}

	return user, nil
}

func (u *userRepo) Delete(ctx context.Context, req *us.UsPrimaryKey) error {
	_, err := u.db.Exec(ctx, `
		UPDATE system_user SET 
			deleted_at = 1
		WHERE id = $1
	`, req.Id)

	if err != nil {
		log.Println("error while deleting user", err)
		return err
	}

	return nil
}
