-- +goose Up
-- +goose StatementBegin
INSERT INTO roles (name) VALUES
('Admin'),
('Doctor'),
('Patient'),
('Pharmacist');

INSERT INTO permissions (groups, name) VALUES
('USER', 'CREATE'),
('USER', 'READ'),
('USER', 'UPDATE'),
('USER', 'DELETE');


-- Admin Role Permissions
INSERT INTO role_permissions (role_id, permission_id, status) VALUES
((SELECT id FROM roles WHERE name='Admin'), (SELECT id FROM permissions WHERE name='CREATE'), TRUE),
((SELECT id FROM roles WHERE name='Admin'), (SELECT id FROM permissions WHERE name='READ'), TRUE),
((SELECT id FROM roles WHERE name='Admin'), (SELECT id FROM permissions WHERE name='UPDATE'), TRUE),
((SELECT id FROM roles WHERE name='Admin'), (SELECT id FROM permissions WHERE name='DELETE'), TRUE);

-- Doctor Role Permissions
INSERT INTO role_permissions (role_id, permission_id, status) VALUES
((SELECT id FROM roles WHERE name='Doctor'), (SELECT id FROM permissions WHERE name='CONSULT_PATIENT'), TRUE),
((SELECT id FROM roles WHERE name='Doctor'), (SELECT id FROM permissions WHERE name='WRITE_PRESCRIPTION'), TRUE),
((SELECT id FROM roles WHERE name='Doctor'), (SELECT id FROM permissions WHERE name='VIEW_MEDICAL_RECORDS'), TRUE),
((SELECT id FROM roles WHERE name='Doctor'), (SELECT id FROM permissions WHERE name='UPDATE_MEDICAL_RECORDS'), TRUE),
((SELECT id FROM roles WHERE name='Doctor'), (SELECT id FROM permissions WHERE name='GIVE_HEALTH_ADVICE'), TRUE);

-- Patient Role Permissions
INSERT INTO role_permissions (role_id, permission_id, status) VALUES
((SELECT id FROM roles WHERE name='Patient'), (SELECT id FROM permissions WHERE name='SCHEDULE_APPOINTMENT'), TRUE),
((SELECT id FROM roles WHERE name='Patient'), (SELECT id FROM permissions WHERE name='CONSULT_DOCTOR'), TRUE),
((SELECT id FROM roles WHERE name='Patient'), (SELECT id FROM permissions WHERE name='VIEW_MEDICAL_RECORDS'), TRUE),
((SELECT id FROM roles WHERE name='Patient'), (SELECT id FROM permissions WHERE name='DOWNLOAD_MEDICAL_RECORDS'), TRUE),
((SELECT id FROM roles WHERE name='Patient'), (SELECT id FROM permissions WHERE name='ORDER_MEDICATION'), TRUE),
((SELECT id FROM roles WHERE name='Patient'), (SELECT id FROM permissions WHERE name='TRACK_HEALTH'), TRUE),
((SELECT id FROM roles WHERE name='Patient'), (SELECT id FROM permissions WHERE name='ACCESS_EDUCATIONAL_CONTENT'), TRUE);

-- Pharmacist Role Permissions
INSERT INTO role_permissions (role_id, permission_id, status) VALUES
((SELECT id FROM roles WHERE name='Pharmacist'), (SELECT id FROM permissions WHERE name='PROCESS_PRESCRIPTIONS'), TRUE),
((SELECT id FROM roles WHERE name='Pharmacist'), (SELECT id FROM permissions WHERE name='MANAGE_MEDICATION_STOCK'), TRUE),
((SELECT id FROM roles WHERE name='Pharmacist'), (SELECT id FROM permissions WHERE name='GIVE_MEDICATION_ADVICE'), TRUE),
((SELECT id FROM roles WHERE name='Pharmacist'), (SELECT id FROM permissions WHERE name='ARRANGE_MEDICATION_DELIVERY'), TRUE);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
