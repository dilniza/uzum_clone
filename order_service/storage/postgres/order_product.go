package postgres

import (
	"context"
	"fmt"
	"log"
	op "uzum_orderclone/genproto/order_product"
	"uzum_orderclone/storage"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type orderProductRepo struct {
	db *pgxpool.Pool
}

func NewOrderProductRepo(db *pgxpool.Pool) storage.OrderProductepoI {
	return &orderProductRepo{
		db: db,
	}
}

func (c *orderProductRepo) Create(ctx context.Context, req *op.CreateOrderProducts) (*op.GetOrderProducts, error) {
	resp := &op.GetOrderProducts{}
	id := uuid.New().String()

	query := `INSERT INTO order_products(
		id,
		product_id,
		count,
		discount_price,
		price,
		order_id,
		created_at)
		VALUES($1,$2,$3,$4,$5,$6,NOW()) `

	_, err := c.db.Exec(ctx, query,
		id,
		req.ProductId,
		req.Count,
		req.DiscountPrice,
		req.Price,
		req.OrderId)

	if err != nil {
		log.Println("error while creating catrgory")
		return resp, err
	}
	category, err := c.GetByID(ctx, &op.OrderProductsPrimaryKey{Id: id})
	if err != nil {
		log.Println("error while getting category by id")
		return nil, err
	}

	return category, nil
}

func (c *orderProductRepo) GetByID(ctx context.Context, req *op.OrderProductsPrimaryKey) (*op.GetOrderProducts, error) {

	resp := &op.GetOrderProducts{}

	fmt.Println(req.Id, "id")

	err := c.db.QueryRow(ctx, `
	SELECT 
	id,
	product_id,
	count,
	discount_price,
	price,
	order_id
	FROM order_products
	WHERE id =$1
	`, req.Id).Scan(&resp.Id, &resp.ProductId, &resp.Count, &resp.DiscountPrice, &resp.Price, &resp.OrderId)

	if err != nil {
		log.Println("error while getting order by id:", err)
		return nil, err
	}

	return resp, nil
}
