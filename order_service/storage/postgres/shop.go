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

type shopRepo struct {
	db *pgxpool.Pool
}

func NewShopRepo(db *pgxpool.Pool) storage.ShopRepoI {
	return &shopRepo{
		db: db,
	}
}

// Create implements storage.ShopRepoI.
func (s *shopRepo) Create(ctx context.Context, req *us.CreateShop) (resp *us.Shop, err error) {
	resp = &us.Shop{}
	id := uuid.NewString()

	_, err = s.db.Exec(ctx, `
		INSERT INTO shop (
			id,
			slug,
			phone,
			name_uz,
			name_ru,
			name_en,
			description_uz,
			description_ru,
			description_en,
			location,
			currency,
			payment_types,
			created_at,
			updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, NOW(), NOW()
		) `, id, req.Slug, req.Phone, req.NameUz, req.NameRu, req.NameEn, req.DescriptionUz, req.DescriptionRu, req.DescriptionEn, req.Location, req.Currency, req.PaymentTypes)

	if err != nil {
		log.Println("error while creating shop in storage", err)
		return nil, err
	}

	shop, err := s.GetById(ctx, &us.ShopPrimaryKey{Id: id})
	if err != nil {
		log.Println("error while getting shop by id", err)
		return nil, err
	}
	return shop, nil
}

// GetByID implements storage.ShopRepoI.
func (s *shopRepo) GetById(ctx context.Context, req *us.ShopPrimaryKey) (*us.Shop, error) {
	resp := &us.Shop{}

	var (
		created_at sql.NullString
		updated_at sql.NullString
	)

	err := s.db.QueryRow(ctx, `
	        SELECT id,
	        slug,
	        phone,
			name_uz,
			name_ru,
			name_en,
			description_uz,
			description_ru,
			description_en,
	        location,
	        currency,
	        payment_types,
	        created_at,
	        updated_at
	        FROM shop
	    WHERE id=$1`, req.Id).Scan(&resp.Id, &resp.Slug, &resp.Phone, &resp.NameUz, &resp.NameRu, &resp.NameEn, &resp.DescriptionUz, &resp.DescriptionRu, &resp.DescriptionEn, &resp.Location, &resp.Currency, &resp.PaymentTypes, &created_at, &updated_at)

	if err != nil {
		log.Println("error while getting shop by id", err)
		return nil, err
	}

	resp.CreatedAt = created_at.String
	resp.UpdatedAt = updated_at.String

	return resp, nil
}

// GetList implements storage.ShopRepoI.
func (s *shopRepo) GetList(ctx context.Context, req *us.GetListShopRequest) (*us.GetListShopResponse, error) {
	resp := &us.GetListShopResponse{}

	var (
		filter         string
		created_at     sql.NullString
		updated_at     sql.NullString
		slug           sql.NullString
		location       sql.NullString
		description_en sql.NullString
		description_ru sql.NullString
		description_uz sql.NullString
		currency       sql.NullString
	)

	offset := (req.Page - 1) * req.Limit

	if req.Search != "" {
		filter = fmt.Sprintf(` AND (phone ILIKE '%%%v%%' OR location ILIKE '%%%v%%' OR currency ILIKE '%%%v%%')`, req.Search, req.Search, req.Search)
	}

	filter += fmt.Sprintf(" OFFSET %v LIMIT %v", offset, req.Limit)

	rows, err := s.db.Query(ctx, `
		SELECT
			id,
			slug,
			phone,
			name_uz,
			name_ru,
			name_en,
			description_uz,
			description_ru,
			description_en,
	        location,
	        currency,
	        payment_types,
	        created_at,
	        updated_at
		FROM shop WHERE deleted_at=0
		ORDER BY created_at DESC`+filter)

	if err != nil {
		log.Println("error while getting all shops", err)
		return nil, err
	}
	defer rows.Close()
	var count int64
	for rows.Next() {
		var shop us.Shop
		count++
		err = rows.Scan(&shop.Id, &slug, &shop.Phone, &shop.NameUz,
			&shop.NameRu, &shop.NameEn, &description_uz, &description_ru,
			&description_en, &location, &currency,
			&shop.PaymentTypes, &created_at, &updated_at)

		if err != nil {
			log.Println("error while scanning shops", err)
			return nil, err
		}
		shop.DescriptionEn = description_en.String
		shop.Slug = slug.String
		shop.Location = location.String
		shop.CreatedAt = created_at.String
		shop.UpdatedAt = updated_at.String
		shop.Currency = currency.String

		resp.Shops = append(resp.Shops, &shop)
	}

	// var count int64
	// err = s.db.QueryRow(ctx, `SELECT COUNT(*) FROM shop`).Scan(&count)
	// if err != nil {
	// 	log.Println("error while counting shops", err)
	// 	return nil, err
	// }

	resp.Count = count

	return resp, nil
}

// Update implements storage.ShopRepoI.
func (s *shopRepo) Update(ctx context.Context, req *us.UpdateShop) (resp *us.Shop, err error) {
	_, err = s.db.Exec(ctx, `
        UPDATE shop SET
            slug = $2,
            phone = $3,
            name_uz = $4,
            name_ru = $5,
            name_en = $6,
            description_uz = $7,
            description_ru = $8,
            description_en = $9,
            location = $10,
            currency = $11,
            payment_types = $12,
            updated_at = NOW()
        WHERE id = $1`, req.Id, req.Slug, req.Phone, req.NameUz, req.NameRu, req.NameEn, req.DescriptionUz, req.DescriptionRu, req.DescriptionEn, req.Location, req.Currency, req.PaymentTypes)

	if err != nil {
		log.Println("error while updating shop in storage", err)
		return nil, err
	}

	shop, err := s.GetById(ctx, &us.ShopPrimaryKey{Id: req.Id})
	if err != nil {
		log.Println("error while getting updated shop by id", err)
		return nil, err
	}

	return shop, nil
}

// Delete implements storage.ShopRepoI.
func (s *shopRepo) Delete(ctx context.Context, req *us.ShopPrimaryKey) error {
	_, err := s.db.Exec(ctx, `
		UPDATE shop SET 
			deleted_at = 1
		WHERE id = $1
	`, req.Id)

	if err != nil {
		log.Println("error while deleting shop", err)
		return err
	}

	return nil
}
