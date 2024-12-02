CREATE TABLE IF NOT EXISTS category (
    id bigserial PRIMARY KEY,
    name varchar(100) NOT NULL,
    parent_id bigint REFERENCES Category(id),
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
    );
