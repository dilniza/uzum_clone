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

type sellerRepo struct {
	db *pgxpool.Pool
}

func NewSellerRepo(db *pgxpool.Pool) storage.SellerRepoI {
	return &sellerRepo{
		db: db,
	}
}

// Create implements storage.UserRepoI.
func (u *sellerRepo) Create(ctx context.Context, req *us.CreateSeller) (*us.Seller, error) {
	// resp := &us.Seller{}
	id := uuid.NewString()

	_, err := u.db.Exec(ctx, `
		INSERT INTO seller (
			id,
			phone,
			email,
			name,
			shop_id
		) VALUES (
			$1, $2, $3, $4, $5
		) `, id, req.Phone, req.Email, req.Name, req.ShopId)

	if err != nil {
		log.Println("error while creating seller in storage", err)
		return nil, err
	}

	user, err := u.GetByID(ctx, &us.SellerPrimaryKey{Id: id})
	if err != nil {
		log.Println("error while getting seller by id", err)
		return nil, err
	}
	return user, nil
}

// GetByID implements storage.UserRepoI.
func (u *sellerRepo) GetByID(ctx context.Context, req *us.SellerPrimaryKey) (*us.Seller, error) {
	resp := &us.Seller{}

	var (
		created_at sql.NullString
		updated_at sql.NullString
	)

	err := u.db.QueryRow(ctx, `
	        SELECT id,
	        phone,
			email,
			name,
			shop_id,
	        created_at,
	        updated_at
	        FROM seller
	    WHERE id=$1`, req.Id).Scan(&resp.Id, &resp.Phone, &resp.Email, &resp.Name, &resp.ShopId,
		&created_at, &updated_at)

	if err != nil {
		log.Println("error while getting seller by id", err)
		return nil, err
	}

	resp.CreatedAt = created_at.String
	resp.UpdatedAt = updated_at.String

	return resp, nil
}

// GetAll implements storage.UserRepoI.
func (u *sellerRepo) GetList(ctx context.Context, req *us.GetListSellerRequest) (*us.GetListSellerResponse, error) {
	resp := &us.GetListSellerResponse{}

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
			email,
			name,
			shop_id,
	        created_at,
	        updated_at
		FROM seller WHERE deleted_at=0
		ORDER BY created_at DESC `+filter)

	if err != nil {
		log.Println("error while getting all sellers", err)
		return nil, err
	}
	defer rows.Close()
	var count int64

	for rows.Next() {
		var seller us.Seller
		count++
		err = rows.Scan(&seller.Id, &seller.Phone, &seller.Email, &seller.Name, &seller.ShopId,
			&created_at, &updated_at)

		if err != nil {
			log.Println("error while scanning sellers", err)
			return nil, err
		}
		seller.CreatedAt = created_at.String
		seller.UpdatedAt = updated_at.String

		resp.Sellers = append(resp.Sellers, &seller)
	}

	resp.Count = count

	return resp, nil
}

// Update implements storage.UserRepoI.
func (u *sellerRepo) Update(ctx context.Context, req *us.UpdateSeller) (resp *us.Seller, err error) {
	_, err = u.db.Exec(ctx, `
        UPDATE seller SET
            phone = $2,
            email = $3,
            name = $4,
            shop_id = $5,
            updated_at = NOW()
        WHERE id = $1`, req.Id, req.Phone, req.Email, req.Name, req.ShopId)

	if err != nil {
		log.Println("error while updating seller in storage", err)
		return nil, err
	}

	seller, err := u.GetByID(ctx, &us.SellerPrimaryKey{Id: req.Id})
	if err != nil {
		log.Println("error while getting updated seller by id", err)
		return nil, err
	}

	return seller, nil
}

// Delete implements storage.UserRepoI.
func (u *sellerRepo) Delete(ctx context.Context, req *us.SellerPrimaryKey) error {
	_, err := u.db.Exec(ctx, `
		UPDATE seller SET 
			deleted_at = 1
		WHERE id = $1
	`, req.Id)

	if err != nil {
		log.Println("error while deleting seller", err)
		return err
	}

	return nil
}
