CREATE TABLE projects (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  description TEXT,
  owner_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE
);
