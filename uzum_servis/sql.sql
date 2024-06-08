CREATE TABLE category (
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


CREATE TABLE product (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
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
    image TEXT ARRAY,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
    deleted_at INTEGER DEFAULT 0
);

CREATE TABLE product_categories (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    product_id UUID NOT NULL,
    category_id UUID NOT NULL,
    FOREIGN KEY (product_id) REFERENCES products(id),
    FOREIGN KEY (category_id) REFERENCES categories(id)
);

CREATE TABLE product_reviews (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    customer_id UUID,
    product_id UUID,
    text VARCHAR(500),
    rating FLOAT,
    order_id UUID,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW()
);

-- UserService
CREATE TABLE customers (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    phone VARCHAR(20) UNIQUE NOT NULL,
    gmail VARCHAR(30) UNIQUE NOT NULL,
    language VARCHAR(2) DEFAULT 'uz',
    date_of_birth DATE,
    gender VARCHAR(10),
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
    deleted_at INTEGER DEFAULT 0
);

CREATE TABLE sellers (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    phone VARCHAR(20) UNIQUE NOT NULL,
    gmail VARCHAR(30) UNIQUE NOT NULL,
    name VARCHAR(255),
    shop_id UUID,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
    deleted_at INTEGER DEFAULT 0,
    FOREIGN KEY (shop_id) REFERENCES shops(id)
);

CREATE TABLE system_users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    phone VARCHAR(20) UNIQUE NOT NULL,
    gmail VARCHAR(30) UNIQUE NOT NULL,
    name VARCHAR(255),
    role VARCHAR(10) NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
    deleted_at INTEGER DEFAULT 0
);

CREATE TABLE branches (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    phone VARCHAR(20) UNIQUE NOT NULL,
    name VARCHAR(20) DEFAULT '',
    location GEOGRAPHY(POLYGON),
    address VARCHAR(255),
    open_time TIME WITHOUT TIME ZONE,
    close_time TIME WITHOUT TIME ZONE,
    active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
    deleted_at INTEGER DEFAULT 0
);

CREATE TABLE shops (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    slug VARCHAR(255) UNIQUE NOT NULL,
    phone VARCHAR(20) UNIQUE NOT NULL,
    name_uz VARCHAR(20) DEFAULT '',
    name_ru VARCHAR(20) DEFAULT '',
    name_en VARCHAR(20) DEFAULT '',
    description_uz VARCHAR(500) DEFAULT '',
    description_ru VARCHAR(500) DEFAULT '',
    description_en VARCHAR(500) DEFAULT '',
    location VARCHAR(255),
    currency VARCHAR(10),
    payment_types TEXT ARRAY,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
    deleted_at INTEGER DEFAULT 0
);

-- OrderService
CREATE TABLE orders (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    external_id VARCHAR(255),
    type VARCHAR(20) NOT NULL,
    customer_phone VARCHAR(20),
    customer_name VARCHAR(20),
    customer_id UUID,
    payment_type VARCHAR(20),
    status VARCHAR(20) NOT NULL,
    to_address VARCHAR(255),
    to_location GEOGRAPHY(POLYGON),
    discount_amount FLOAT,
    amount FLOAT,
    delivery_price FLOAT,
    paid BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
    deleted_at INTEGER DEFAULT 0,
    FOREIGN KEY (customer_id) REFERENCES customers(id)
);

CREATE TABLE order_products (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    product_id UUID,
    count INTEGER,
    discount_price FLOAT,
    price FLOAT,
    order_id UUID,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
    deleted_at INTEGER DEFAULT 0,
    FOREIGN KEY (product_id) REFERENCES products(id),
    FOREIGN KEY (order_id) REFERENCES orders(id)
);

CREATE TABLE order_status_notes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    order_id UUID,
    status VARCHAR(20) NOT NULL,
    user_id UUID,
    reason VARCHAR(100),
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
    FOREIGN KEY (order_id) REFERENCES orders(id),
    FOREIGN KEY (user_id) REFERENCES system_users(id)
);
