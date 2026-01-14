CREATE TABLE driver_locations (
    driver_id UUID PRIMARY KEY,

    latitude DOUBLE PRECISION NOT NULL,
    longitude DOUBLE PRECISION NOT NULL,

    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_location_driver
        FOREIGN KEY (driver_id) REFERENCES drivers(id)
);
CREATE INDEX idx_driver_locations_lat_lng
ON driver_locations(latitude, longitude);
