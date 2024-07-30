-- +goose Up
-- +goose StatementBegin
INSERT INTO roles (name) VALUES
('Admin', true, '2024-12-12 10.00', ''),
('Doctor', true, '2024-12-12 10.00', ''),
('Patient', true, '2024-12-12 10.00', ''),
('Pharmacist', true, '2024-12-12 10.00', '');

INSERT INTO permissions (groups, name, created_at, updated_at) VALUES
('PERMISSION', 'CREATE', '2024-12-12 10.00', '')

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM permissions WHERE groups = 'PERMISSION' AND name = 'CREATE' AND created_at = '2024-12-12 10.00';

DELETE FROM roles WHERE name = 'Admin' AND created_at = '2024-12-12 10.00';
DELETE FROM roles WHERE name = 'Doctor' AND created_at = '2024-12-12 10.00';
DELETE FROM roles WHERE name = 'Patient' AND created_at = '2024-12-12 10.00';
DELETE FROM roles WHERE name = 'Pharmacist' AND created_at = '2024-12-12 10.00';
-- +goose StatementEnd
