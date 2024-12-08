CREATE TABLE IF NOT EXISTS payments (
    id bigserial PRIMARY KEY,
    order_id bigint NOT NULL REFERENCES orders(id),
    payment_method varchar(50) NOT NULL,
    amount numeric(10, 2) NOT NULL,
    status varchar(50) NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
    );
