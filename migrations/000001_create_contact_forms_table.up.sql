-- Create the contact_forms table
CREATE TABLE IF NOT EXISTS contact_forms (
    id UUID PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    timestamp TIMESTAMP,
    message TEXT
);
