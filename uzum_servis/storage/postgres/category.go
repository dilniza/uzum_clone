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

	var exists int
	err = c.db.QueryRow(ctx, "SELECT COUNT(slug) FROM category WHERE slug = $1", req.Slug).Scan(&exists)
	if err != nil {
		log.Println("error while checking for duplicate slug")
		return nil, err
	}
	if exists > 0 {
		return nil, fmt.Errorf("category with slug '%s' already exists", req.Slug)
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
        parent_id,
        created_at,
        updated_at
    ) VALUES (
        $1, $2, $3, $4, $5, $6, $7, 
        CASE 
            WHEN $8 = '' THEN NULL 
            ELSE $8::uuid 
        END, 
        CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
    )`, id, req.Slug, req.NameUz, req.NameRu, req.NameEn, req.Active, req.OrderNo, req.ParentId)

	if err != nil {
		log.Println("error while creating category")
		return nil, err
	}
	fmt.Println(req)
	category, err := c.GetByID(ctx, &ct.CategoryPrimaryKey{Id: id})
	if err != nil {
		log.Println("error while getting category by id")
		return nil, err
	}

	return category, nil
}

func (c *categoryRepo) Update(ctx context.Context, req *ct.UpdateCategory) (resp *ct.Category, err error) {
	resp = &ct.Category{}
	var parentid sql.NullString
	if req.ParentId != " " || req.ParentId != "" {
		parentid.String = req.ParentId
	}

	_, err = c.db.Exec(ctx, `
		UPDATE category SET 
			slug = $2,
			name_uz = $3,
			name_ru = $4,
			name_en = $5,
			active = $6,
			order_no = $7,
			parent_id = $8,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = $1
	`, req.Id, req.Slug, req.NameUz, req.NameRu, req.NameEn, req.Active, req.OrderNo, parentid)

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

func (c *categoryRepo) GetByID(ctx context.Context, req *ct.CategoryPrimaryKey) (resp *ct.Category, err error) {
	resp = &ct.Category{}
	var (
		parentid  sql.NullString
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
			active,
			order_no,
			COALESCE(parent_id::text, '') AS parent_id,
			created_at,
			updated_at,
			deleted_at
		FROM category
		WHERE id = $1
	`, req.Id)

	err = row.Scan(&resp.Id, &resp.Slug, &resp.NameUz, &resp.NameRu, &resp.NameEn, &resp.Active, &resp.OrderNo, &parentid, &createdat, &updatedat, &resp.DeletedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("category not found with ID %s", req.Id)
		}
		return nil, err
	}

	resp.UpdatedAt = updatedat.String
	resp.CreatedAt = createdat.String
	resp.ParentId = parentid.String

	return resp, nil
}

func (c *categoryRepo) GetAll(ctx context.Context, req *ct.GetAllCategoryRequest) (resp *ct.GetAllCategoryResponse, err error) {

	resp = &ct.GetAllCategoryResponse{}

	log.Printf("GetAll request received: limit=%d, offset=%d, search=%s", req.Limit, req.Offset, req.Search)

	offset := (req.Offset - 1) * req.Limit

	var filter string
	if req.Search != "" {
		filter = fmt.Sprintf(` AND (name_uz ILIKE '%%%v%%' OR name_ru ILIKE '%%%v%%' OR name_en ILIKE '%%%v%%')`, req.Search, req.Search, req.Search)
	}

	query := `
        SELECT
            id,
            slug,
            name_uz,
            name_ru,
            name_en,
            active,
            order_no,
            COALESCE(parent_id::text, '') AS parent_id,
            created_at,
            updated_at,
            deleted_at
        FROM category
        WHERE deleted_at = 0 
    ` + filter +
		` ORDER BY created_at DESC
        LIMIT $1 OFFSET $2`

	rows, err := c.db.Query(ctx, query, req.Limit, offset)

	if err != nil {
		log.Printf("Error while querying categories: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			category  ct.Category
			parentid  sql.NullString
			createdat sql.NullString
			updatedat sql.NullString
		)

		err = rows.Scan(
			&category.Id,
			&category.Slug,
			&category.NameUz,
			&category.NameRu,
			&category.NameEn,
			&category.Active,
			&category.OrderNo,
			&parentid,
			&createdat,
			&updatedat,
			&category.DeletedAt,
		)
		if err != nil {
			log.Printf("Error while scanning categories: %v", err)
			return nil, err
		}

		category.UpdatedAt = updatedat.String
		category.CreatedAt = createdat.String
		category.ParentId = parentid.String

		resp.Categorys = append(resp.Categorys, &category)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Row error: %v", err)
		return nil, err
	}

	log.Printf("Successfully retrieved categories: %+v", resp.Categorys)

	err = c.db.QueryRow(ctx, `SELECT COUNT(id) FROM category WHERE deleted_at = 0 `+filter).Scan(&resp.Count)

	if err != nil {
		log.Printf("Error while counting categories: %v", err)
		return nil, err
	}

	log.Printf("Successfully retrieved categories: %+v", resp)

	return resp, nil
}

func (c *categoryRepo) Delete(ctx context.Context, req *ct.CategoryPrimaryKey) error {
	_, err := c.db.Exec(ctx, `
		UPDATE category SET 
			deleted_at = 1
		WHERE id = $1
	`, req.Id)

	if err != nil {
		log.Println("error while deleting category")
		return err
	}

	return nil
}
