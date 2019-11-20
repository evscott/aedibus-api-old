CREATE TABLE IF NOT EXISTS Assignments (
    RepoName varchar (255) PRIMARY KEY,
    MasterBranch varchar (255) NOT NULL,
    LatestCommit varchar (40) NOT NULL
);

CREATE TABLE IF NOT EXISTS Submissions (
    RepoName varchar (255) PRIMARY KEY REFERENCES Assignments(RepoName),
    MasterBranch varchar (255) NOT NULL,
    WorkingBranch varchar (255) NOT NULL,
    LatestCommit varchar (40) NOT NULL,
    Submitted bool DEFAULT FALSE,
    Reviewed bool DEFAULT FALSE
);