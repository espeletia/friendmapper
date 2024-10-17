-- +goose Up
-- +goose StatementBegin
ALTER TABLE places
ADD COLUMN lat DOUBLE PRECISION,
ADD COLUMN lon DOUBLE PRECISION;

CREATE OR REPLACE FUNCTION update_point_column()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.lat IS NOT NULL AND NEW.lon IS NOT NULL THEN
        NEW.point := ST_SetSRID(ST_MakePoint(NEW.lon, NEW.lat), 4326);
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_point_before_insert_or_update
BEFORE INSERT OR UPDATE ON places
FOR EACH ROW
EXECUTE FUNCTION update_point_column();


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
