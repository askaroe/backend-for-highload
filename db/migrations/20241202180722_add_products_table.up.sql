CREATE TABLE IF NOT EXISTS product (
    id bigserial PRIMARY KEY,
    name varchar(100) NOT NULL,
    description text,
    price numeric(10, 2) NOT NULL,
    stock_quantity int NOT NULL DEFAULT 0,
    category_id bigint NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
    );
