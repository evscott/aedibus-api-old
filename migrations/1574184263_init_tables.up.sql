CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS Assignments (
    name varchar (255) UNIQUE,
    blob_sha varchar(40),
    PRIMARY KEY (name)
);

CREATE TABLE IF NOT EXISTS Dropboxes (
    name varchar (255) UNIQUE NOT NULL,
    assignment_name varchar (255) REFERENCES Assignments(name) NOT NULL,
    submitted bool DEFAULT FALSE,
    grade bool DEFAULT FALSE,
    PRIMARY KEY (name, assignment_name)
);

CREATE TABLE IF NOT EXISTS Files (
    name varchar(255) NOT NULL,
    assignment_name varchar (255) REFERENCES Assignments(name) NOT NULL,
    dropbox_name varchar (255) REFERENCES Dropboxes(name) NOT NULL,
    PRIMARY KEY (name, assignment_name, dropbox_name)
);