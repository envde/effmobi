CREATE TABLE subscriptions (
    id           BIGSERIAL PRIMARY KEY,
    service_name VARCHAR(255) NOT NULL,
    price        INTEGER      NOT NULL,
    user_id      UUID         NOT NULL,
    start_date   DATE         NOT NULL,
    end_date     DATE,
    created_at   TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);