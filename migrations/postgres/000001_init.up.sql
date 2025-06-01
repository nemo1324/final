DROP TABLE IF EXISTS users;

CREATE TABLE IF NOT EXISTS users (
                                     id BIGSERIAL PRIMARY KEY,
                                     name TEXT NOT NULL,
                                     login TEXT NOT NULL UNIQUE,
                                     pass TEXT NOT NULL,
                                     created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
                                     status TEXT NOT NULL CHECK (status IN ('register', 'logout'))
);
