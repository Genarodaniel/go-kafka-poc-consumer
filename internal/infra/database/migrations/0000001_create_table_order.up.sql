BEGIN;

CREATE TABLE IF NOT EXISTS tab_order (
    id uuid DEFAULT gen_random_uuid(),
    store_id uuid NOT NULL,
    client_id uuid NOT NULL,
    active VARCHAR NOT NULL DEFAULT 1,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);


COMMIT;