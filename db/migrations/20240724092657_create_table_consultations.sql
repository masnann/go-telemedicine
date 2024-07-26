-- +goose Up
-- +goose StatementBegin
CREATE TABLE consultations (
    id SERIAL PRIMARY KEY,
    patient_id INTEGER NOT NULL,
    patient_email VARCHAR(255) NOT NULL,
    patient_name VARCHAR(255) NOT NULL,
    schedule_id INTEGER NOT NULL,
    doctor_id INTEGER NOT NULL,
    doctor_name VARCHAR(255) NOT NULL,
    start_time VARCHAR(255) NOT NULL,
    end_time VARCHAR(255) NOT NULL,
    consultation_type VARCHAR(255) NOT NULL,
    status VARCHAR(50) NOT NULL,
    notes TEXT,
    created_at VARCHAR(255) NOT NULL,
    updated_at VARCHAR(255) NOT NULL,
    deleted_at VARCHAR(255)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS consultations;
-- +goose StatementEnd
