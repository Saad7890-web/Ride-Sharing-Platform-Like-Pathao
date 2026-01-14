CREATE TABLE driver_documents (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    driver_id UUID NOT NULL,

    document_type VARCHAR(50) NOT NULL,
   

    document_url TEXT NOT NULL,

    verified BOOLEAN NOT NULL DEFAULT FALSE,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_document_driver
        FOREIGN KEY (driver_id) REFERENCES drivers(id)
);
