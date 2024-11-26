--drop table users;
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(100) UNIQUE,
    is_active BOOLEAN,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    uuid UUID DEFAULT gen_random_uuid(),
    age INT,
    salary NUMERIC(10, 2),
    joining_date DATE,
    gender VARCHAR(20)
);
--DELETE FROM users