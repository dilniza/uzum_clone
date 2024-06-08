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

CREATE TABLE IF NOT EXISTS "product" (
    id UUID PRIMARY KEY DEFAULT,
    slug VARCHAR(255) UNIQUE NOT NULL,
    name_uz VARCHAR(20) DEFAULT '',
    name_ru VARCHAR(20) DEFAULT '',
    name_en VARCHAR(20) DEFAULT '',
    description_uz VARCHAR(500) DEFAULT '',
    description_ru VARCHAR(500) DEFAULT '',
    description_en VARCHAR(500) DEFAULT '',
    active BOOLEAN DEFAULT TRUE,
    order_no INTEGER DEFAULT 0,
    in_price FLOAT,
    out_price FLOAT,
    left_count INTEGER,
    discount_percent FLOAT DEFAULT 0,
    image TEXT[],
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at INTEGER DEFAULT 0
);