CREATE TABLE vehicles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    driver_id UUID NOT NULL UNIQUE,

    type VARCHAR(20) NOT NULL,
    

    plate_number VARCHAR(50) NOT NULL UNIQUE,
    model VARCHAR(100),

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_vehicle_driver
        FOREIGN KEY (driver_id) REFERENCES drivers(id)
);
