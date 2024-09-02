package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	ct "service/genproto/catalog_service"
	"service/storage"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type productRepo struct {
	db *pgxpool.Pool
}

func NewProductRepo(db *pgxpool.Pool) storage.ProductRepoI {
	return &productRepo{
		db: db,
	}
}

func (c *productRepo) Create(ctx context.Context, req *ct.CreateProduct) (resp *ct.Product, err error) {
	resp = &ct.Product{}
	id := uuid.NewString()
	slug := ""

	_, err = c.db.Exec(ctx, `
    INSERT INTO product (
        id,
        slug,
        name_uz,
        name_ru,
        name_en,
        description_uz,
        description_ru,
        description_en,
        active,
        order_no,
        in_price,
        out_price,
        left_count,
        discount_percent,
        image,
        created_at,
        updated_at
    ) VALUES (
        $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
    )`, id, slug, req.NameUz, req.NameRu, req.NameEn, req.DescriptionUz, req.DescriptionRu, req.DescriptionEn, req.Active, req.OrderNo, req.InPrice, req.OutPrice, req.LeftCount, req.DiscountPercent, req.Image)

	if err != nil {
		log.Println("error while creating product")
		return nil, err
	}

	product, err := c.GetByID(ctx, &ct.ProductPrimaryKey{Id: id})
	if err != nil {
		log.Println("error while getting product by id")
		return nil, err
	}

	return product, nil
}

func (c *productRepo) Update(ctx context.Context, req *ct.UpdateProduct) (resp *ct.Product, err error) {
	resp = &ct.Product{}
	slug := ""

	_, err = c.db.Exec(ctx, `
		UPDATE product SET
			slug = $2
			name_uz = $3,
			name_ru = $4,
			name_en = $5,
			description_uz = $6,
			description_ru = $7,
			description_en = $8,
			active = $9,
			order_no = $10,
			in_price = $11,
			out_price = $12,
			left_count = $13,
			discount_percent = $14,
			image = $15,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = $1
	`, req.Id, slug, req.NameUz, req.NameRu, req.NameEn, req.DescriptionUz, req.DescriptionRu, req.DescriptionEn, req.Active, req.OrderNo, req.InPrice, req.OutPrice, req.LeftCount, req.DiscountPercent, req.Image)

	if err != nil {
		log.Println("error while updating product")
		return nil, err
	}

	product, err := c.GetByID(ctx, &ct.ProductPrimaryKey{Id: req.Id})
	if err != nil {
		log.Println("error while getting product by id")
		return nil, err
	}

	return product, nil
}

func (c *productRepo) GetByID(ctx context.Context, req *ct.ProductPrimaryKey) (resp *ct.Product, err error) {
	resp = &ct.Product{}
	var (
		createdat sql.NullString
		updatedat sql.NullString
	)

	row := c.db.QueryRow(ctx, `
		SELECT
			id,
			slug,
			name_uz,
			name_ru,
			name_en,
			description_uz,
			description_ru,
			description_en,
			active,
			order_no,
			in_price,
			out_price,
			left_count,
			discount_percent,
			image,
			created_at,
			updated_at,
			deleted_at
		FROM product
		WHERE id = $1
	`, req.Id)

	err = row.Scan(&resp.Id, &resp.Slug, &resp.NameUz, &resp.NameRu, &resp.NameEn, &resp.DescriptionUz, &resp.DescriptionRu, &resp.DescriptionEn, &resp.Active, &resp.OrderNo, &resp.InPrice, &resp.OutPrice, &resp.LeftCount, &resp.DiscountPercent, &resp.Image, &createdat, &updatedat, &resp.DeletedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("product not found with ID %s", req.Id)
		}
		return nil, err
	}

	resp.UpdatedAt = updatedat.String
	resp.CreatedAt = createdat.String

	return resp, nil
}

func (c *productRepo) GetAll(ctx context.Context, req *ct.GetAllProductRequest) (resp *ct.GetAllProductResponse, err error) {
	resp = &ct.GetAllProductResponse{}

	log.Printf("GetAll request received: limit=%d, offset=%d, search=%s", req.Limit, req.Offset, req.Search)

	offset := (req.Offset - 1) * req.Limit

	var filter string
	if req.Search != "" {
		filter = fmt.Sprintf(` AND (name_uz ILIKE '%%%v%%' OR name_ru ILIKE '%%%v%%' OR name_en ILIKE '%%%v%%' OR description_uz ILIKE '%%%v%%' OR description_ru ILIKE '%%%v%%' OR description_en ILIKE '%%%v%%')`, req.Search, req.Search, req.Search, req.Search, req.Search, req.Search)
	}

	query := `
        SELECT
            id,
            slug,
            name_uz,
            name_ru,
            name_en,
            description_uz,
            description_ru,
            description_en,
            active,
            order_no,
            in_price,
            out_price,
            left_count,
            discount_percent,
            image,
            created_at,
            updated_at,
            deleted_at
        FROM product
        WHERE deleted_at = 0 
    ` + filter +
		` ORDER BY created_at DESC
        LIMIT $1 OFFSET $2`

	rows, err := c.db.Query(ctx, query, req.Limit, offset)

	if err != nil {
		log.Printf("Error while querying products: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			product   ct.Product
			createdat sql.NullString
			updatedat sql.NullString
		)

		err = rows.Scan(
			&product.Id,
			&product.Slug,
			&product.NameUz,
			&product.NameRu,
			&product.NameEn,
			&product.DescriptionUz,
			&product.DescriptionRu,
			&product.DescriptionEn,
			&product.Active,
			&product.OrderNo,
			&product.InPrice,
			&product.OutPrice,
			&product.LeftCount,
			&product.DiscountPercent,
			&product.Image,
			&createdat,
			&updatedat,
			&product.DeletedAt,
		)
		if err != nil {
			log.Printf("Error while scanning products: %v", err)
			return nil, err
		}

		product.UpdatedAt = updatedat.String
		product.CreatedAt = createdat.String

		resp.Products = append(resp.Products, &product)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Row error: %v", err)
		return nil, err
	}

	log.Printf("Successfully retrieved products: %+v", resp.Products)

	err = c.db.QueryRow(ctx, `SELECT COUNT(id) FROM product WHERE deleted_at = 0 `+filter).Scan(&resp.Count)

	if err != nil {
		log.Printf("Error while counting products: %v", err)
		return nil, err
	}

	log.Printf("Successfully retrieved products: %+v", resp)

	return resp, nil
}

func (c *productRepo) Delete(ctx context.Context, req *ct.ProductPrimaryKey) error {
	_, err := c.db.Exec(ctx, `
		UPDATE product SET 
			deleted_at = 1
		WHERE id = $1
	`, req.Id)

	if err != nil {
		log.Println("error while deleting product")
		return err
	}

	return nil
}
