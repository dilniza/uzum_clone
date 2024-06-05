package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	ct "server/genproto/category_service"
	"server/storage"

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

// func (c *categoryRepo) Create(ctx context.Context, req *ct.CreateCategory) (resp *ct.Category, err error) {
// 	fmt.Println("----postgres", req)
// 	resp = &ct.Category{}
// 	id := uuid.NewString()
// 	var parentID uuid.UUID

// 	// Check for duplicate slug (using exists with count)
// 	var exists int
// 	err = c.db.QueryRow(ctx, "SELECT COUNT(slug) FROM category WHERE slug = $1", req.Slug).Scan(&exists)
// 	if err != nil {
// 		log.Println("error while checking for duplicate slug")
// 		return nil, err
// 	}
// 	if exists > 0 { // If count is greater than 0, a slug exists
// 		return nil, fmt.Errorf("category with slug '%s' already exists", req.Slug)
// 	}

// 	if req.ParentId != "" && req.ParentId != " " {
// 		if parentID, err = uuid.Parse(req.ParentId); err != nil {
// 			return nil, fmt.Errorf("invalid parent ID: %w", err)
// 		}
// 	}

// 	_, err = c.db.Exec(ctx, `
// 		INSERT INTO category (
// 			id,
// 			slug,
// 			name_uz,
// 			name_ru,
// 			name_en,
// 			active,
// 			order_no,
// 			parent_id
// 		) VALUES (
// 			$1, $2, $3, $4, $5, $6, $7, $8
// 		) `, id, req.Slug, req.NameUz, req.NameRu, req.NameEn, req.Active, req.OrderNo, parentID)

// 	if err != nil {
// 		log.Println("error while creating category")
// 		return nil, err
// 	}

// 	category, err := c.GetByID(ctx, &ct.CategoryPrimaryKey{Id: id})
// 	if err != nil {
// 		log.Println("error while getting category by id")
// 		return nil, err
// 	}

// 	return category, nil
// }

func (c *categoryRepo) Create(ctx context.Context, req *ct.CreateCategory) (resp *ct.Category, err error) {
	resp = &ct.Category{}
	id := uuid.NewString()

	// Check for duplicate slug
	var exists int
	err = c.db.QueryRow(ctx, "SELECT COUNT(slug) FROM category WHERE slug = $1", req.Slug).Scan(&exists)
	if err != nil {
		log.Println("error while checking for duplicate slug")
		return nil, err
	}
	if exists > 0 {
		return nil, fmt.Errorf("category with slug '%s' already exists", req.Slug)
	}
	//create exec
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
	//get it by id
	fmt.Println(req)
	category, err := c.GetByID(ctx, &ct.CategoryPrimaryKey{Id: id})
	if err != nil {
		log.Println("error while getting category by id")
		return nil, err
	}

	return category, nil
}

// func (c *categoryRepo) GetByID(ctx context.Context, req *ct.CategoryPrimaryKey) (resp *ct.Category, err error) {
// 	resp = &ct.Category{}

// 	var parentID sql.NullString // Use sql.NullString to handle null values

// 	err = c.db.QueryRow(ctx, `
// 		SELECT
// 			id,
// 			slug,
// 			name_uz,
// 			name_ru,
// 			name_en,
// 			active,
// 			order_no,
// 			COALESCE(parent_id),
// 			created_at,
// 			updated_at,
// 			deleted_at
// 		FROM category
// 		WHERE id = $1
// 	`, req.Id).Scan(&resp.Id, &resp.Slug, &resp.NameUz, &resp.NameRu, &resp.NameEn, &resp.Active, &resp.OrderNo, &parentID, &resp.CreatedAt, &resp.UpdatedAt, &resp.DeletedAt)

// 	if err != nil {
// 		log.Println("error while getting category by id")
// 		return nil, err
// 	}

// 	if parentID.Valid {
// 		resp.ParentId = parentID.String
// 	} else {
// 		resp.ParentId = ""
// 	}

// 	return resp, nil
// }

func (c *categoryRepo) GetByID(ctx context.Context, req *ct.CategoryPrimaryKey) (resp *ct.Category, err error) {
	resp = &ct.Category{}
	var (
		//deletedat sql.NullInt64
		parentid  sql.NullString
		createdat sql.NullString
		updatedat sql.NullString
	)
	//getbyid query
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

func (c *categoryRepo) Update(ctx context.Context, req *ct.UpdateCategory) (resp *ct.Category, err error) {
	resp = &ct.Category{}

	var parentID uuid.UUID
	if req.ParentId != "" && req.ParentId != " " {
		if parentID, err = uuid.Parse(req.ParentId); err != nil {
			return nil, fmt.Errorf("invalid parent ID: %w", err)
		}
	} else {
		parentID = uuid.Nil
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
			updated_at = NOW()
		WHERE id = $1
	`, req.Id, req.Slug, req.NameUz, req.NameRu, req.NameEn, req.Active, req.OrderNo, parentID)

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
		var parentID uuid.UUID // Declare as UUID

		err = rows.Scan(&category.Id, &category.Slug, &category.NameUz, &category.NameRu, &category.NameEn, &category.Active, &category.OrderNo, &parentID, &category.CreatedAt, &category.UpdatedAt, &category.DeletedAt)

		if err != nil {
			log.Println("error while scanning categories")
			return nil, err
		}

		category.ParentId = parentID.String()
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
			deleted_at = 1
		WHERE id = $1
	`, req.Id)

	if err != nil {
		log.Println("error while deleting category")
		return err
	}

	return nil
}
