-- +goose Up
-- +goose StatementBegin
CREATE TABLE schedules (
    id SERIAL PRIMARY KEY,
    doctor_id INTEGER NOT NULL,
    date VARCHAR(255) NOT NULL,
    start_time VARCHAR(255) NOT NULL,
    end_time VARCHAR(255) NOT NULL,
    is_available BOOLEAN DEFAULT TRUE,
    created_at VARCHAR(255) NOT NULL,
    updated_at VARCHAR(255) NOT NULL,
    deleted_at VARCHAR(255),
    FOREIGN KEY (doctor_id) REFERENCES users(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS schedules;
-- +goose StatementEnd
