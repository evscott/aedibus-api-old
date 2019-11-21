CREATE TABLE IF NOT EXISTS Assignments (
    name varchar (255) PRIMARY KEY,
    branch varchar (255) NOT NULL,
    blob_shah varchar(40)
);

CREATE TABLE IF NOT EXISTS Submissions (
    assignment varchar (255) PRIMARY KEY REFERENCES Assignments(name),
    branch varchar (255) NOT NULL,
    submitted bool DEFAULT FALSE,
    grade bool DEFAULT FALSE
);