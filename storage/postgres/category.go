package postgres

import (
	"context"
	"database/sql"
	"log"
	ct "uzum_clone/genproto/category_service"
	"uzum_clone/storage"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type categoryRepo struct {
	db *pgxpool.Pool
}

func NewCategoryRepo(db *pgxpool.Pool) storage.CategoryRepoI {
	return &categoryRepo{
		db: db,
	}
}

func (c *categoryRepo) Create(ctx context.Context, req *ct.CreateCategory) (resp *ct.Category, err error) {

	resp = &ct.Category{}
	id := uuid.NewString()

	if req.ParentId == "" {
		req.ParentId = id
	}

	_, err = c.db.Exec(ctx, `
		INSERT INTO category (
			id,
			slug,
			name_uz,
			name_ru,
			name_en,
			active,
			order_no,
			parent_id
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8
		) `, id, req.Slug, req.NameUz, req.NameRu, req.NameEn, req.Active, req.OrderNo, req.ParentId)

	if err != nil {
		log.Println("error while creating category")
		return nil, err
	}

	category, err := c.GetByID(ctx, &ct.CategoryPrimaryKey{Id: id})
	if err != nil {
		log.Println("error while getting category by id")
		return nil, err
	}

	return category, nil
}

func (c *categoryRepo) GetByID(ctx context.Context, req *ct.CategoryPrimaryKey) (resp *ct.Category, err error) {

	resp = &ct.Category{}

	var ParentId sql.NullString

	err = c.db.QueryRow(ctx, `
		SELECT
			id,
			slug,
			name_uz,
			name_ru,
			name_en,
			active,
			order_no,
			parent_id,
			created_at,
			updated_at,
			deleted_at
		FROM category
		WHERE id = $1
	`, req.Id).Scan(&resp.Id, &resp.Slug, &resp.NameUz, &resp.NameRu, &resp.NameEn, &resp.Active, &resp.OrderNo, &ParentId, &resp.CreatedAt, &resp.UpdatedAt, &resp.DeletedAt)

	if err != nil {
		log.Println("error while getting category by id")
		return nil, err
	}

	resp.ParentId = ParentId.String

	return resp, nil
}

func (c *categoryRepo) Update(ctx context.Context, req *ct.UpdateCategory) (resp *ct.Category, err error) {
	resp = &ct.Category{}

	_, err = c.db.Exec(ctx, `
		UPDATE category SET 
			slug = $2,
			name_uz = $3,
			name_ru = $4,
			name_en = $5,
			active = $6,
			order_no = $7,
			parent_id = $8,
			updated_at = NOW()
		WHERE id = $1
	`, req.Id, req.Slug, req.NameUz, req.NameRu, req.NameEn, req.Active, req.OrderNo, req.ParentId)

	if err != nil {
		log.Println("error while updating category")
		return nil, err
	}

	category, err := c.GetByID(ctx, &ct.CategoryPrimaryKey{Id: req.Id})
	if err != nil {
		log.Println("error while getting category by id")
		return nil, err
	}

	return category, nil
}

func (c *categoryRepo) GetAll(ctx context.Context, req *ct.GetListCategoryRequest) (resp *ct.GetListCategoryResponse, err error) {

	resp = &ct.GetListCategoryResponse{}

	rows, err := c.db.Query(ctx, `
		SELECT
			id,
			slug,
			name_uz,
			name_ru,
			name_en,
			active,
			order_no,
			parent_id,
			created_at,
			updated_at,
			deleted_at
		FROM category
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`, req.Limit, req.Offset)

	if err != nil {
		log.Println("error while getting all categories")
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var category ct.Category
		var ParentId sql.NullString

		err = rows.Scan(&category.Id, &category.Slug, &category.NameUz, &category.NameRu, &category.NameEn, &category.Active, &category.OrderNo, &ParentId, &category.CreatedAt, &category.UpdatedAt, &category.DeletedAt)

		if err != nil {
			log.Println("error while scanning categories")
			return nil, err
		}

		category.ParentId = ParentId.String
		resp.Categorys = append(resp.Categorys, &category)
	}

	var count int64
	err = c.db.QueryRow(ctx, `SELECT COUNT(*) FROM category`).Scan(&count)
	if err != nil {
		log.Println("error while counting categories")
		return nil, err
	}

	resp.Count = count

	return resp, nil
}

func (c *categoryRepo) Delete(ctx context.Context, req *ct.CategoryPrimaryKey) error {

	_, err := c.db.Exec(ctx, `
		UPDATE category SET 
			deleted_at = NOW()
		WHERE id = $1
	`, req.Id)

	if err != nil {
		log.Println("error while deleting category")
		return err
	}

	return nil
}
