DROP TABLE IF EXISTS users;

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    surname VARCHAR NOT NULL
);

DROP TABLE IF EXISTS projects;

CREATE TABLE projects (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    description VARCHAR
);

DROP TABLE IF EXISTS bugs;

CREATE TABLE bugs (
    id SERIAL PRIMARY KEY,
    description VARCHAR(100) NOT NULL,
    creation_date TIMESTAMP NOT NULL DEFAULT NOW(),
    user_id INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id),
    project_id INT NOT NULL,
    FOREIGN KEY (project_id) REFERENCES projects(id)
)
