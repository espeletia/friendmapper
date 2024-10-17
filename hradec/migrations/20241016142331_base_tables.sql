-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS postgis;

CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TABLE IF NOT EXISTS places (
    id varchar(64) NOT NULL PRIMARY KEY,
    type VARCHAR(255) NOT NULL,
    sub_type VARCHAR(255),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    accessibility INTEGER NOT NULL,
    accessibility_note TEXT,
    capacity INTEGER,
    capacity_note TEXT,
    phones VARCHAR(1024),
    web VARCHAR(2048) NOT NULL,
    okres VARCHAR(128) NOT NULL,
    obce VARCHAR(128) NOT NULL,
    address VARCHAR(512) NOT NULL,
    point GEOMETRY(Point, 4326),

    like_count INT NOT NULL DEFAULT 0
);

CREATE SEQUENCE users_id_seq
START WITH 30000
INCREMENT BY 1
NO MINVALUE
NO MAXVALUE
CACHE 1;

CREATE TABLE IF NOT EXISTS users (
    id INT NOT NULL PRIMARY KEY DEFAULT nextval('users_id_seq'),
    password_hash VARCHAR(255) NOT NULL,
    username VARCHAR(64) NOT NULL,
    profile_picture_url VARCHAR(255),
    display_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER update_users_updated_at
BEFORE UPDATE ON users
FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

CREATE SEQUENCE collection_id_seq
START WITH 40000
INCREMENT BY 1
NO MINVALUE
NO MAXVALUE
CACHE 1;

CREATE TABLE IF NOT EXISTS collections (
    id INT NOT NULL PRIMARY KEY DEFAULT nextval('collection_id_seq'),
    name VARCHAR(256) NOT NULL,
    description TEXT,
    thumbnail VARCHAR(255),
    user_id INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    like_count INTEGER NOT NULL DEFAULT 0,

    CONSTRAINT fk_collection_user_id
    FOREIGN KEY (user_id)
    REFERENCES users (id)
    ON DELETE CASCADE
    ON UPDATE CASCADE

);

CREATE TRIGGER update_collections_updated_at
BEFORE UPDATE ON collections
FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();


CREATE TABLE IF NOT EXISTS collections_places (
    collection_id INT NOT NULL,
    place_id VARCHAR(64) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (collection_id, place_id),

    CONSTRAINT fk_collection_places_collection_id
    FOREIGN KEY (collection_id)
    REFERENCES collections (id)
    ON DELETE CASCADE
    ON UPDATE CASCADE,

    CONSTRAINT fk_collection_places_place_id
    FOREIGN KEY (place_id)
    REFERENCES places (id)
    ON DELETE CASCADE
    ON UPDATE CASCADE

);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
