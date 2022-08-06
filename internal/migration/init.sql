CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS books (id UUID DEFAULT uuid_generate_v1() PRIMARY KEY, title varchar, author varchar, rating int, publish_date date);
