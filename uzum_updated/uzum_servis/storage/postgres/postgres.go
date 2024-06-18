package postgres

import (
	"context"
	"fmt"
	"log"
	"service/config"
	"service/storage"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Store struct {
	db              *pgxpool.Pool
	category        storage.CategoryRepoI
	product         storage.ProductRepoI
	productCategory storage.ProductCategoryRepoI
	productReview   storage.ProductReviewRepoI

	customer storage.CustomerRepoI
	user     storage.UserRepoI
	seller   storage.SellerRepoI
	branch   storage.BranchRepoI
	shop     storage.ShopRepoI
}

func NewPostgres(ctx context.Context, cfg config.Config) (storage.StorageI, error) {
	config, err := pgxpool.ParseConfig(fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase,
	))
	if err != nil {
		return nil, err
	}

	config.MaxConns = cfg.PostgresMaxConnections

	pool, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	return &Store{
		db: pool,
	}, err
}

func (s *Store) CloseDB() {
	s.db.Close()
}

func (l *Store) Log(ctx context.Context, level pgx.LogLevel, msg string, data map[string]interface{}) {
	args := make([]interface{}, 0, len(data)+2) // making space for arguments + level + msg
	args = append(args, level, msg)
	for k, v := range data {
		args = append(args, fmt.Sprintf("%s=%v", k, v))
	}
	log.Println(args...)
}

func (s *Store) Category() storage.CategoryRepoI {
	if s.category == nil {
		s.category = NewCategoryRepo(s.db)
	}

	return s.category
}

func (s *Store) Product() storage.ProductRepoI {
	if s.product == nil {
		s.product = NewProductRepo(s.db)
	}

	return s.product
}

func (s *Store) ProductCategory() storage.ProductCategoryRepoI {
	if s.productCategory == nil {
		s.productCategory = NewProductCategoryRepo(s.db)
	}

	return s.productCategory
}

func (s *Store) ProductReview() storage.ProductReviewRepoI {
	if s.productReview == nil {
		s.productReview = NewProductReviewRepo(s.db)
	}

	return s.productReview
}

func (s *Store) Customer() storage.CustomerRepoI {
	if s.customer == nil {
		s.customer = NewCustomerRepo(s.db)
	}

	return s.customer
}

func (s *Store) User() storage.UserRepoI {
	if s.user == nil {
		s.user = NewUserRepo(s.db)
	}

	return s.user
}

func (s *Store) Seller() storage.SellerRepoI {
	if s.seller == nil {
		s.seller = NewSellerRepo(s.db)
	}
	return s.seller
}

func (s *Store) Branch() storage.BranchRepoI {
	if s.branch == nil {
		s.branch = NewBranchRepo(s.db)
	}
	return s.branch
}

func (s *Store) Shop() storage.ShopRepoI {
	if s.shop == nil {
		s.shop = NewShopRepo(s.db)
	}
	return s.shop
}
