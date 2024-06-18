package postgres

import (
	"context"
	"log"
	ct "service/genproto/catalog_service"
	"service/storage"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type productReviewRepo struct {
	db *pgxpool.Pool
}

func NewProductReviewRepo(db *pgxpool.Pool) storage.ProductReviewRepoI {
	return &productReviewRepo{
		db: db,
	}
}

func (c *productReviewRepo) Create(ctx context.Context, req *ct.CreateProductReview) (resp *ct.ProductReview, err error) {
	resp = &ct.ProductReview{}
	id := uuid.NewString()

	_, err = c.db.Exec(ctx, `
		INSERT INTO product_review (
			id,
			customer_id,
			product_id,
			text,
			created_at
		) VALUES (
			$1, $2::uuid, $3::uuid, $4, CURRENT_TIMESTAMP
		)
	`, id, req.CustomerId, req.ProductId, req.Text)

	if err != nil {
		log.Println("error while creating product review")
		return nil, err
	}

	productReview, err := c.GetByID(ctx, &ct.ProductReviewPrimaryKey{Id: id})
	if err != nil {
		log.Println("error while getting product review by id")
		return nil, err
	}

	return productReview, nil
}

func (c *productReviewRepo) GetByID(ctx context.Context, req *ct.ProductReviewPrimaryKey) (resp *ct.ProductReview, err error) {
	resp = &ct.ProductReview{}

	err = c.db.QueryRow(ctx, `
		SELECT
			id,
			customer_id,
			product_id,
			text,
			rating,
			created_at
		FROM product_review
		WHERE id = $1
	`, req.Id).Scan(&resp.Id, &resp.CustomerId, &resp.ProductId, &resp.Text, &resp.Rating, &resp.CreatedAt)

	if err != nil {
		log.Println("error while getting product review by id")
		return nil, err
	}

	return resp, nil
}

func (c *productReviewRepo) Update(ctx context.Context, req *ct.UpdateProductReview) (resp *ct.ProductReview, err error) {
	resp = &ct.ProductReview{}

	_, err = c.db.Exec(ctx, `
		UPDATE product_review SET
			customer_id = $2::uuid,
			product_id = $3::uuid,
			text = $4,
			rating = $5
		WHERE id = $1
	`, req.Id, req.CustomerId, req.ProductId, req.Text, req.Rating)

	if err != nil {
		log.Println("error while updating product review")
		return nil, err
	}

	productReview, err := c.GetByID(ctx, &ct.ProductReviewPrimaryKey{Id: req.Id})
	if err != nil {
		log.Println("error while getting product review by id")
		return nil, err
	}

	return productReview, nil
}

func (c *productReviewRepo) GetAll(ctx context.Context, req *ct.GetAllProductReviewRequest) (resp *ct.GetAllProductReviewResponse, err error) {
	resp = &ct.GetAllProductReviewResponse{}

	rows, err := c.db.Query(ctx, `
		SELECT
			id,
			customer_id,
			product_id,
			text,
			rating,
			created_at
		FROM product_review
		ORDER BY id DESC
		LIMIT $1 OFFSET $2
	`, req.Limit, req.Offset)

	if err != nil {
		log.Println("error while getting all product reviews")
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var productReview ct.ProductReview
		err = rows.Scan(&productReview.Id, &productReview.CustomerId, &productReview.ProductId, &productReview.Text, &productReview.Rating, &productReview.CreatedAt)
		if err != nil {
			log.Println("error while scanning product review")
			return nil, err
		}

		resp.ProductReviews = append(resp.ProductReviews, &productReview)
	}

	var count int64
	err = c.db.QueryRow(ctx, `SELECT COUNT(*) FROM product_review`).Scan(&count)
	if err != nil {
		log.Println("error while counting product reviews")
		return nil, err
	}

	resp.Count = count
	return resp, nil
}

func (c *productReviewRepo) Delete(ctx context.Context, req *ct.ProductReviewPrimaryKey) error {
	_, err := c.db.Exec(ctx, `
		DELETE FROM product_review
		WHERE id = $1
	`, req.Id)

	if err != nil {
		log.Println("error while deleting product review")
		return err
	}

	return nil
}
