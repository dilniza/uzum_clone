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

CREATE TABLE IF NOT EXISTS "product_category" (
    id UUID PRIMARY KEY,
    product_id UUID NOT NULL,
    category_id UUID NOT NULL,
    FOREIGN KEY (product_id) REFERENCES products(id),
    FOREIGN KEY (category_id) REFERENCES categories(id)
);

CREATE TABLE IF NOT EXISTS "product_review" (
    id UUID PRIMARY KEY,
    customer_id UUID,
    product_id UUID,
    text VARCHAR(500),
    rating FLOAT,
    order_id UUID,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS "customer" (
    id UUID PRIMARY KEY,
    firstname VARCHAR(55) NOT NUll,
    lastname VARCHAR(55),
    phone VARCHAR(20) NOT NULL,
    gmail VARCHAR(30) NOT NULL,
    language VARCHAR(2) CHECK (language IN ('uz', 'ru', 'en')) NOT NULL,
    date_of_birth DATE NOT NULL,
    gender VARCHAR(6) CHECK (gender IN ('male', 'female')) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at INTEGER DEFAULT 0
);

CREATE TABLE IF NOT EXISTS "seller" (
    id UUID PRIMARY KEY,
    phone VARCHAR(20) UNIQUE NOT NULL,
    gmail VARCHAR(30) UNIQUE NOT NULL,
    name VARCHAR(255),
    shop_id UUID,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at INTEGER DEFAULT 0,
    FOREIGN KEY (shop_id) REFERENCES shops(id)
);

CREATE TABLE IF NOT EXISTS "system_user" (
    id UUID PRIMARY KEY,
    phone VARCHAR(20) UNIQUE NOT NULL,
    gmail VARCHAR(30) UNIQUE NOT NULL,
    name VARCHAR(255),
    role VARCHAR(7) CHECK (role IN ('admin', 'courier')) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at INTEGER DEFAULT 0
);

CREATE TABLE IF NOT EXISTS "branch" (
    id UUID PRIMARY KEY,
    phone VARCHAR(20) UNIQUE NOT NULL,
    name VARCHAR(20) DEFAULT '',
    location POLYGON NOT NULL,
    address VARCHAR(255),
    open_time TIME NOT NULL,
    close_time TIME NOT NULL,
    active BOOLEAN NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at INTEGER DEFAULT 0
);

CREATE TABLE IF NOT EXISTS "shop" (
    id UUID PRIMARY KEY,
    slug VARCHAR(255) UNIQUE NOT NULL,
    phone VARCHAR(20) UNIQUE NOT NULL,
    name_uz VARCHAR(20) DEFAULT '',
    name_ru VARCHAR(20) DEFAULT '',
    name_en VARCHAR(20) DEFAULT '',
    description_uz VARCHAR(500) DEFAULT '',
    description_ru VARCHAR(500) DEFAULT '',
    description_en VARCHAR(500) DEFAULT '',
    location VARCHAR(255) NOT NULL,
    currency VARCHAR(3) CHECK (currency IN ('USD', 'EUR', 'UZS')) NOT NULL,
    payment_types VARCHAR[] NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at INTEGER DEFAULT 0
);