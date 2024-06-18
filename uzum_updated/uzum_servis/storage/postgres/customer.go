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

type customerRepo struct {
	db *pgxpool.Pool
}

func NewCustomerRepo(db *pgxpool.Pool) storage.CustomerRepoI {
	return &customerRepo{
		db: db,
	}
}

// Create implements storage.CustomerRepoI.
func (c *customerRepo) Create(ctx context.Context, req *us.CreateCustomer) (resp *us.Customer, err error) {
	resp = &us.Customer{}
	id := uuid.NewString()

	_, err = c.db.Exec(ctx, `
		INSERT INTO customer (
			id,
			firstname,
			lastname,
			phone,
			email,
			language,
			date_of_birth,
			gender
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8
		) `, id, req.Firstname, req.Lastname, req.Phone, req.Email, req.Languange, req.DateOfBirth, req.Gender)

	if err != nil {
		log.Println("error while creating customer in storage")
		return nil, err
	}

	customer, err := c.GetByID(ctx, &us.CustomerPrimaryKey{Id: id})
	if err != nil {
		log.Println("error while getting customer by id")
		return nil, err
	}
	return customer, nil
}

// GetByID implements storage.CustomerRepoI.
func (c *customerRepo) GetByID(ctx context.Context, req *us.CustomerPrimaryKey) (*us.Customer, error) {
	resp := &us.Customer{}

	var (
		created_at    sql.NullString
		updated_at    sql.NullString
		date_of_birth sql.NullString
	)

	err := c.db.QueryRow(ctx, `
	        SELECT 
			id,
	        firstname,
			lastname,
			phone,
			email,
			language,
			date_of_birth,
			gender,
	        created_at,
	        updated_at
	        FROM customer
	    WHERE id=$1`, req.Id).Scan(&resp.Id, &resp.Firstname, &resp.Lastname, &resp.Phone, &resp.Email,
		&resp.Languange, &date_of_birth, &resp.Gender, &created_at, &updated_at)

	if err != nil {
		log.Println("error while getting customer by id", err)
		return nil, err
	}

	resp.DateOfBirth = date_of_birth.String
	resp.CreatedAt = created_at.String
	resp.UpdatedAt = updated_at.String

	return resp, nil
}

// GetAll implements storage.CustomerRepoI.
func (c *customerRepo) GetList(ctx context.Context, req *us.GetListCustomerRequest) (resp *us.GetListCustomerResponse, err error) {
	resp = &us.GetListCustomerResponse{}
	var (
		filter        string
		created_at    sql.NullString
		updated_at    sql.NullString
		date_of_birth sql.NullString
	)
	offset := (req.Page - 1) * req.Limit

	if req.Search != "" {
		filter = fmt.Sprintf(` AND (firstname ILIKE '%%%v%%' OR lastname ILIKE '%%%v%%' OR email ILIKE '%%%v%%')`, req.Search, req.Search, req.Search)
	}

	filter += fmt.Sprintf(" OFFSET %v LIMIT %v", offset, req.Limit)

	rows, err := c.db.Query(ctx, `
        SELECT
            id,
            firstname,
            lastname,
            phone,
            email,
            language,
            date_of_birth,
            gender,
            created_at,
            updated_at
        FROM customer WHERE deleted_at is null
    `+filter)

	if err != nil {
		log.Println("error while getting all customers:", err)
		return nil, err
	}

	defer rows.Close()

	var count int64

	for rows.Next() {
		var customer us.Customer
		count++
		err = rows.Scan(&customer.Id, &customer.Firstname, &customer.Lastname, &customer.Phone, &customer.Email,
			&customer.Languange, &date_of_birth, &customer.Gender, &created_at, &updated_at)

		if err != nil {
			log.Println("error while scanning customers:", err)
			return nil, err
		}
		customer.CreatedAt = created_at.String
		customer.UpdatedAt = updated_at.String
		customer.DateOfBirth = date_of_birth.String

		resp.Customers = append(resp.Customers, &customer)
	}

	if err = rows.Err(); err != nil {
		log.Println("rows iteration error:", err)
		return nil, err
	}

	resp.Count = count

	return resp, nil
}

// Update implements storage.CustomerRepoI.
func (c *customerRepo) Update(ctx context.Context, req *us.UpdateCustomer) (resp *us.Customer, err error) {

	_, err = c.db.Exec(ctx, `
        UPDATE customer SET
            firstname = $2,
            lastname = $3,
            phone = $4,
            email = $5,
            language = $6,
            date_of_birth = $7,
            updated_at = NOW()
        WHERE id = $1`, req.Id, req.Firstname, req.Lastname, req.Phone, req.Email,
		req.Languange, req.DateOfBirth)

	if err != nil {
		log.Println("error while updating customer in storage", err)
		return nil, err
	}

	customer, err := c.GetByID(ctx, &us.CustomerPrimaryKey{Id: req.Id})
	if err != nil {
		log.Println("error while getting updated customer by id", err)
		return nil, err
	}

	return customer, nil
}

// Delete implements storage.CustomerRepoI.
func (c *customerRepo) Delete(ctx context.Context, req *us.CustomerPrimaryKey) error {
	_, err := c.db.Exec(ctx, `
		UPDATE customer SET 
			deleted_at = 1
		WHERE id = $1
	`, req.Id)

	if err != nil {
		log.Println("error while deleting customer")
		return err
	}

	return nil
}
