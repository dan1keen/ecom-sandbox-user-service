CREATE TABLE users (
   id SERIAL PRIMARY KEY,
   phone varchar(20) UNIQUE NOT NULL,
   email varchar(255) UNIQUE NULL,
   firstname varchar(255) NULL,
   lastname varchar(255) NULL,
   created_at TIMESTAMP DEFAULT NOW(),
   updated_at TIMESTAMP DEFAULT NOW()
);