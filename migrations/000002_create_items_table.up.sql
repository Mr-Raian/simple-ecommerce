-- Create the "items" table
CREATE TABLE items (
    id UUID PRIMARY KEY,
    price BIGINT NOT NULL,
    metadata JSON
);
