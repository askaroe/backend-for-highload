CREATE TABLE IF NOT EXISTS orders (
    id bigserial PRIMARY KEY,
    user_id bigint NOT NULL,
    order_status varchar(50) NOT NULL,
    total_amount numeric(10, 2) NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
    );