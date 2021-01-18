# Account service

## install database

create table users
```sql
-- Iinitialize uuid extension ------------------------------------

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Not all database features are supported. Do not use for backup.
-- Table Definition ----------------------------------------------

CREATE TABLE users (
    uuid uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    email character varying UNIQUE,
    password character varying
);

-- Indices -------------------------------------------------------

CREATE UNIQUE INDEX users_pkey ON users(uuidid uuid_ops);
```
