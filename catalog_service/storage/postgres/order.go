package postgres

import (
	"context"
	"fmt"
	"log"
	ct "uzum_orderclone/genproto/orders_service"
	"uzum_orderclone/storage"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type orderRepo struct {
	db *pgxpool.Pool
}

func NewOrderRepo(db *pgxpool.Pool) storage.OrderyRepoI {
	return &orderRepo{
		db: db,
	}
}

func (c *orderRepo) Create(ctx context.Context, req *ct.CreateOrder) (resp *ct.GetOrder, err error) {
	resp = &ct.GetOrder{}

	id := uuid.NewString()
	_, err = c.db.Exec(ctx, `
        INSERT INTO orders(
            id, 
            external_id,
            type,
            customer_phone,
            customer_name,
            customer_id,
            payment_type, 
            status,
            to_address,
            to_location, -- Assuming this is a latitude and longitude pair
            discount_amount,
            amount,
            delivery_price,
            paid,
            courier_id,
            courier_phone,
            courier_name,
            created_at) 
            VALUES (
            $1,
            $2,
            $3,
            $4,
            $5,
            $6,
            $7,
            $8,
            $9,
            ST_MakePoint($10,$11), -- Corrected syntax for to_location
            $12,
            $13,
            $14,
            $15,
            $16,
            $17,
            $18,
            NOW()
        )`, id, req.ExternalId, req.Type, req.CustomerPhone, req.CustomerName, req.CustomerId, req.PaymentType,
		req.Status, req.ToAddress,
		req.ToLocation.Longitude, req.ToLocation.Latitude, req.DiscountAmount, req.Amount, req.DeliveryPrice,
		req.Paid, req.CourierId, req.CourierPhone, req.CourierName)

	// Handle error if any
	if err != nil {
		log.Println("error while creating order")
		return resp, err
	}
	category, err := c.GetByID(ctx, &ct.OrderPrimaryKey{Id: id})
	if err != nil {
		log.Println("error while getting category by id")
		return nil, err
	}

	return category, nil
}
func (c *orderRepo) GetByID(ctx context.Context, req *ct.OrderPrimaryKey) (resp *ct.GetOrder, err error) {
	var long, lat float64

	resp = &ct.GetOrder{}

	err = c.db.QueryRow(ctx, `
        SELECT
            id,
            external_id,
            type,
            customer_phone,
            customer_name,
            customer_id,
            payment_type,
            status,
            to_address,
            ST_X(to_location) as longitude,  -- Extract longitude from the point
            ST_Y(to_location) as latitude,   -- Extract latitude from the point
            discount_amount,
            amount,
            delivery_price,
            paid,
            courier_id,
            courier_phone,
            courier_name
            FROM orders
        WHERE id = $1
    `, req.Id).Scan(&resp.Id, &resp.ExternalId, &resp.Type, &resp.CustomerPhone, &resp.CustomerName, &resp.CustomerId, &resp.PaymentType,
		&resp.Status, &resp.ToAddress, &long, &lat, &resp.DiscountAmount, &resp.Amount, &resp.DeliveryPrice,
		&resp.Paid, &resp.CourierId, &resp.CourierPhone, &resp.CourierName)

	if err != nil {
		log.Println("error while getting order by id:", err)
		return nil, err
	}

	resp.ToLocation = []*ct.Location{{
		Longitude: long,
		Latitude:  lat,
	}}

	fmt.Println(resp)
	return resp, nil
}
