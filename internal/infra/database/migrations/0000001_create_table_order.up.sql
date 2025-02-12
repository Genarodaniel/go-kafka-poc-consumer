BEGIN;


CREATE TYPE status AS ENUM ('created', 'shipping', 'delivered', 'done');

CREATE TABLE IF NOT EXISTS tab_order (
    id uuid DEFAULT gen_random_uuid(),
    store_id uuid NOT NULL,
    client_id uuid NOT NULL,
    active SMALLINT NOT NULL DEFAULT 1,
    current_status status,
    notification_email VARCHAR(250) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);


COMMIT;