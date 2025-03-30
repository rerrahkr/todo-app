CREATE TABLE IF NOT EXISTS todolist (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO todolist (title, description) VALUES
('Sample Task 1', 'This is a sample task description.'),
('Sample Task 2', 'This is another sample task description.');
