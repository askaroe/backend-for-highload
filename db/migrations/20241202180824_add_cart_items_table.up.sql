CREATE TABLE IF NOT EXISTS cartItems (
    id bigserial PRIMARY KEY,
    cart_id bigint NOT NULL REFERENCES ShoppingCart(id),
    product_id bigint NOT NULL REFERENCES Product(id),
    quantity int NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
    );
