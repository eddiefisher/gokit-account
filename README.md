# Account service

## install database

initialize uuid extension
- CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table users
```sql
-- Not all database features are supported. Do not use for backup.
-- Table Definition ----------------------------------------------

CREATE TABLE users (
    uuid uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    email character varying,
    password character varying
);

-- Indices -------------------------------------------------------

CREATE UNIQUE INDEX users_pkey ON users(id uuid_ops);
```
