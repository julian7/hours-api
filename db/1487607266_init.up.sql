CREATE TABLE clients (
  id   SERIAL PRIMARY KEY,
  name VARCHAR(32) NOT NULL,
  UNIQUE (name)
);

CREATE TABLE projects (
  id        SERIAL PRIMARY KEY,
  client_id INTEGER NOT NULL,
  name      VARCHAR(32) NOT NULL,
  start     DATE NOT NULL,
  finish    DATE NOT NULL,
  UNIQUE (name, client_id),
  FOREIGN KEY (client_id) REFERENCES clients (id)
);
