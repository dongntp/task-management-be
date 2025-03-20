CREATE TYPE role_enum AS ENUM (
  'Employee',
  'Employer'
);

CREATE TYPE status_enum AS ENUM (
  'Pending',
  'InProgress',
  'Complete'
);

CREATE TABLE IF NOT EXISTS account (
  username TEXT NOT NULL PRIMARY KEY,
  password TEXT NOT NULL,
  role role_enum NOT NULL,
  active BOOLEAN NOT NULL DEFAULT true,

  last_updated TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS task (
  id TEXT NOT NULL PRIMARY KEY,
  title TEXT NOT NULL,
  assignee TEXT,
  description TEXT,
  status status_enum NOT NULL DEFAULT 'Pending',

  last_updated TIMESTAMP NOT NULL DEFAULT NOW(),
  CONSTRAINT task_account_fk FOREIGN KEY (assignee) REFERENCES account (username) ON DELETE SET NULL
);
