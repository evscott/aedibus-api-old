CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS Assignments (
    id uuid DEFAULT uuid_generate_v1() UNIQUE,
    name varchar (255) UNIQUE,
    blob_sha varchar(40),
    created_at timestamp DEFAULT now(),
    PRIMARY KEY (id, name)
);

CREATE TABLE IF NOT EXISTS Dropboxes (
    id uuid DEFAULT uuid_generate_v1() UNIQUE,
    name varchar (255) NOT NULL,
    aid uuid REFERENCES Assignments(id) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS Files (
    id uuid DEFAULT uuid_generate_v1() UNIQUE,
    name varchar(255) NOT NULL,
    aid uuid REFERENCES Assignments(id) NOT NULL,
    did uuid REFERENCES Dropboxes(id) NOT NULL,
    commit_id varchar(40) NOT NULL,
    created_at timestamp DEFAULT now(),
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS Submissions (
    id uuid DEFAULT uuid_generate_v1() UNIQUE,
    aid uuid REFERENCES Assignments(id) NOT NULL,
    did uuid REFERENCES Dropboxes(id) NOT NULL,
    pr_number smallint,
    created_at timestamp DEFAULT now(),
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS Submission_Results (
    id uuid DEFAULT uuid_generate_v1() UNIQUE,
    sid uuid REFERENCES Submissions(id) NOT NULL,
    tests_ran smallint,
    tests_passed smallint,
    reviewed boolean,
    created_at timestamp DEFAULT now(),
    PRIMARY KEY (id)
);