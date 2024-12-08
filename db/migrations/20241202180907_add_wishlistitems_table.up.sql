CREATE TABLE IF NOT EXISTS wishListItem (
    id bigserial PRIMARY KEY,
    wishlist_id bigint NOT NULL REFERENCES wishlists(id),
    product_id bigint NOT NULL REFERENCES product(id),
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
    );
