package postgres

import (
	"context"
	"fmt"
	"log"
	st "uzum_orderclone/genproto/order_status_notes"
	"uzum_orderclone/storage"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type orderStatusRepo struct {
	db *pgxpool.Pool
}

func NewOrderStatuRepo(db *pgxpool.Pool) storage.OrderStatusI {

	return &orderStatusRepo{
		db: db,
	}
}

func (c *orderStatusRepo) Create(ctx context.Context, req *st.CreateStatusRequest) (*st.GetOrderStatusResponse, error) {

	resp := &st.GetOrderStatusResponse{}

	id := uuid.New().String()

	query := `INSERT INTO order_status_notes(
	id,
	order_id,
	status,
	user_id,
	reason,
	created_at)
	VALUES($1,$2,$3,$4,$5,NOW())
`
	_, err := c.db.Exec(ctx, query,
		id,
		req.OrderId,
		req.Status,
		req.UserId,
		req.Reason)

	if err != nil {
		log.Println("error while creating ordersatatus")
		return resp, err
	}

	status, err := c.GetByID(ctx, &st.OrderPrimaryKeyRequest{Id: id})
	if err != nil {
		log.Println("error while getting category by id")
		return nil, err
	}

	return status, nil
}

func (c *orderStatusRepo) GetByID(ctx context.Context, req *st.OrderPrimaryKeyRequest) (*st.GetOrderStatusResponse, error) {

	resp := &st.GetOrderStatusResponse{}

	err := c.db.QueryRow(ctx, `SELECT 
	id,
	order_id,
	status,
	user_id,
	reason
	FROM order_status_notes WHERE id=$1`, req.Id).Scan(&resp.Id,
		&resp.OrderId, &resp.Status, &resp.UserId, &resp.Reason)

	if err != nil {
		log.Println("error whilr gettting order by id :", err)
		return nil, err
	}

	return resp, nil
}
func (c *orderStatusRepo) PUTCH(ctx context.Context, req *st.OrderPrimaryStatusKeyRequest) (*st.GetOrderStatusResponse, error) {
	resp := &st.GetOrderStatusResponse{}
	var query string

	fmt.Println(req.Status)
	// Determine the query based on the status
	if req.Status == "cancelled" || req.Status == "finished" {
		query = `
			UPDATE order_status_notes
			SET status = $2, updated_at = NOW(), deleted_at = 1
			WHERE order_id = $1 AND deleted_at = 0;
		`
	} else {
		query = `
			UPDATE order_status_notes
			SET status = $2, updated_at = NOW()
			WHERE order_id = $1 AND deleted_at = 0;
		`
	}

	// Execute the update query
	_, err := c.db.Exec(ctx, query, req.OrderId, req.Status)
	if err != nil {
		log.Printf("error while updating order status for order_id %s: %v", req.OrderId, err)
		return resp, err
	}

	// Fetch the updated status
	status, err := c.GetStatusByID(ctx, &st.OrderPrimaryStatusKeyRequest{OrderId: req.OrderId})
	if err != nil {
		log.Printf("error while getting updated status for order_id %s: %v", req.OrderId, err)
		return nil, err
	}

	return status, nil
}

func (c *orderStatusRepo) GetStatusByID(ctx context.Context, req *st.OrderPrimaryStatusKeyRequest) (*st.GetOrderStatusResponse, error) {

	resp := &st.GetOrderStatusResponse{}

	err := c.db.QueryRow(ctx, `SELECT 
	id,
	order_id,
	status,
	user_id,
	reason
	FROM order_status_notes WHERE order_id=$1`, req.OrderId).Scan(&resp.Id,
		&resp.OrderId, &resp.Status, &resp.UserId, &resp.Reason)

	if err != nil {
		log.Println("error whilr gettting order by id :", err)
		return nil, err
	}

	return resp, nil
}
