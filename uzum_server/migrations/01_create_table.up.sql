CREATE TABLE IF NOT EXISTS "category" (
    id UUID PRIMARY KEY,
    slug VARCHAR(255) UNIQUE NOT NULL,
    name_uz VARCHAR(20) DEFAULT '',
    name_ru VARCHAR(20) DEFAULT '',
    name_en VARCHAR(20) DEFAULT '',
    order_no INTEGER DEFAULT 0,
    active BOOLEAN DEFAULT TRUE,
    parent_id UUID,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at INTEGER DEFAULT 0
);