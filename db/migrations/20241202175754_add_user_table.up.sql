CREATE TABLE IF NOT EXISTS users(
    id bigserial PRIMARY KEY,
    createdAt    timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updatedAt    timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    username     varchar(50),
    email        varchar(50),
    password     varchar(50),
    first_name   varchar(50),
    last_name    varchar(50)
);