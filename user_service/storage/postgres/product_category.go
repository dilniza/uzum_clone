package postgres

import (
	"context"
	"log"
	ct "service/genproto/catalog_service"
	"service/storage"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type productCategoryRepo struct {
	db *pgxpool.Pool
}

func NewProductCategoryRepo(db *pgxpool.Pool) storage.ProductCategoryRepoI {
	return &productCategoryRepo{
		db: db,
	}
}

func (c *productCategoryRepo) Create(ctx context.Context, req *ct.CreateProductCategory) (resp *ct.ProductCategory, err error) {
	resp = &ct.ProductCategory{}
	id := uuid.NewString()

	_, err = c.db.Exec(ctx, `
		INSERT INTO product_category (
			id,
			product_id,
			category_id
		) VALUES (
			$1, $2::uuid, $3::uuid
		)
	`, id, req.ProductId, req.CategoryId)

	if err != nil {
		log.Println("error while creating product category")
		return nil, err
	}

	productCategory, err := c.GetByID(ctx, &ct.ProductCategoryPrimaryKey{Id: id})
	if err != nil {
		log.Println("error while getting product category by id")
		return nil, err
	}

	return productCategory, nil
}

func (c *productCategoryRepo) GetByID(ctx context.Context, req *ct.ProductCategoryPrimaryKey) (resp *ct.ProductCategory, err error) {
	resp = &ct.ProductCategory{}

	err = c.db.QueryRow(ctx, `
		SELECT
			id,
			product_id,
			category_id
		FROM product_category
		WHERE id = $1
	`, req.Id).Scan(&resp.Id, &resp.ProductId, &resp.CategoryId)

	if err != nil {
		log.Println("error while getting product category by id")
		return nil, err
	}

	return resp, nil
}

func (c *productCategoryRepo) Update(ctx context.Context, req *ct.UpdateProductCategory) (resp *ct.ProductCategory, err error) {
	resp = &ct.ProductCategory{}

	_, err = c.db.Exec(ctx, `
		UPDATE product_category SET
			product_id = $2::uuid,
			category_id = $3::uuid
		WHERE id = $1
	`, req.Id, req.ProductId, req.CategoryId)

	if err != nil {
		log.Println("error while updating product category")
		return nil, err
	}

	productCategory, err := c.GetByID(ctx, &ct.ProductCategoryPrimaryKey{Id: req.Id})
	if err != nil {
		log.Println("error while getting product category by id")
		return nil, err
	}

	return productCategory, nil
}

func (c *productCategoryRepo) GetAll(ctx context.Context, req *ct.GetAllProductCategoryRequest) (resp *ct.GetAllProductCategoryResponse, err error) {
	resp = &ct.GetAllProductCategoryResponse{}

	rows, err := c.db.Query(ctx, `
		SELECT
			id,
			product_id,
			category_id
		FROM product_category
		ORDER BY id DESC
		LIMIT $1 OFFSET $2
	`, req.Limit, req.Offset)

	if err != nil {
		log.Println("error while getting all product categories")
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var productCategory ct.ProductCategory
		err = rows.Scan(&productCategory.Id, &productCategory.ProductId, &productCategory.CategoryId)
		if err != nil {
			log.Println("error while scanning product categories")
			return nil, err
		}

		resp.ProductCategorys = append(resp.ProductCategorys, &productCategory)
	}

	var count int64
	err = c.db.QueryRow(ctx, `SELECT COUNT(*) FROM product_category`).Scan(&count)
	if err != nil {
		log.Println("error while counting product categories")
		return nil, err
	}

	resp.Count = count
	return resp, nil
}

func (c *productCategoryRepo) Delete(ctx context.Context, req *ct.ProductCategoryPrimaryKey) error {
	_, err := c.db.Exec(ctx, `
		DELETE FROM product_category
		WHERE id = $1
	`, req.Id)

	if err != nil {
		log.Println("error while deleting product category")
		return err
	}

	return nil
}
