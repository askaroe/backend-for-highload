CREATE TABLE IF NOT EXISTS reviews (
    id bigserial PRIMARY KEY,
    product_id bigint NOT NULL REFERENCES product(id),
    user_id bigint NOT NULL,
    rating int NOT NULL CHECK (rating >= 1 AND rating <= 5),
    comment text,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
    );
