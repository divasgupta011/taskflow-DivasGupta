-- Insert user
INSERT INTO users (id, name, email, password)
VALUES (
  uuid_generate_v4(),
  'Test User',
  'test@example.com',
  '$2a$12$h7ui5y5v9oR2FS6.RXuJVOZG3DaAVHyrkJeRUrHB/sg4p8JMmkDjm'
);

-- Insert project (linked to user)
INSERT INTO projects (id, name, description, owner_id)
VALUES (
  uuid_generate_v4(),
  'Sample Project',
  'This is a seeded project',
  (SELECT id FROM users WHERE email = 'test@example.com')
);

-- Insert tasks (linked to project)
INSERT INTO tasks (title, description, status, priority, project_id)
VALUES
(
  'Task 1',
  'First seeded task',
  'todo',
  'medium',
  (SELECT id FROM projects WHERE name = 'Sample Project')
),
(
  'Task 2',
  'Second seeded task',
  'in_progress',
  'high',
  (SELECT id FROM projects WHERE name = 'Sample Project')
),
(
  'Task 3',
  'Third seeded task',
  'done',
  'low',
  (SELECT id FROM projects WHERE name = 'Sample Project')
);