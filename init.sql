CREATE SCHEMA api;

CREATE TABLE applications
(
    id           serial  CONSTRAINT applications_pkey PRIMARY KEY,
    name         VARCHAR(80) NOT NULL,
    client_id    VARCHAR(80) NOT NULL,
    secret       VARCHAR(80) NOT NULL,
    redirect_url VARCHAR(100),
    created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

