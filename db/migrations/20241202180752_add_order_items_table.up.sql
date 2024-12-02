CREATE TABLE IF NOT EXISTS orderItems (
    id bigserial PRIMARY KEY,
    order_id bigint NOT NULL REFERENCES "Order"(id),
    product_id bigint NOT NULL REFERENCES Product(id),
    quantity int NOT NULL,
    price numeric(10, 2) NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
    );
