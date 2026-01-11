INSERT INTO roles (name) VALUES ('ADMIN');
INSERT INTO permissions (name) VALUES ('ride:create');

INSERT INTO role_permissions (role_id, permission_id)
VALUES (1, 1);

INSERT INTO user_roles (user_id, role_id)
VALUES (
  (SELECT id FROM users WHERE email = 'admin@pathao.com'),
  1
);
