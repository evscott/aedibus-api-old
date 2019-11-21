CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS Assignments (
    name varchar (255) PRIMARY KEY,
    branch varchar (255) NOT NULL,
    blob_shah varchar(40)
);

CREATE TABLE IF NOT EXISTS Submissions (
    id uuid UNIQUE DEFAULT uuid_generate_v1(),
    assignment varchar (255) REFERENCES Assignments(name),
    branch varchar (255) NOT NULL,
    submitted bool DEFAULT FALSE,
    grade bool DEFAULT FALSE,
    PRIMARY KEY (id, assignment)
);

CREATE TABLE IF NOT EXISTS File (
    submissions_id uuid REFERENCES Submissions(id),
    file_name varchar(255) NOT NULL,
    PRIMARY KEY (submissions_id, file_name)
);