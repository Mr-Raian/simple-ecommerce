CREATE TYPE order_status_enum AS ENUM ('PAID', 'UNPAID', 'EXPIRED');
CREATE TABLE orders (
    id UUID PRIMARY KEY,
    item_id UUID NOT NULL,
    price BIGINT NOT NULL,
    email VARCHAR NOT NULL,
    payment_method VARCHAR NOT NULL,
    payment_url VARCHAR NOT NULL,
    payment_id VARCHAR NOT NULL,
    order_status order_status_enum NOT NULL,
    FOREIGN KEY (item_id) REFERENCES items(id)
);
