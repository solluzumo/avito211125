CREATE TYPE pull_request_status AS ENUM ('MERGED','OPEN');

CREATE TABLE Users (
    user_id VARCHAR(50) PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    team_name VARCHAR(20),
    is_active BOOLEAN NOT NULL,
   
);

CREATE TABLE Team (
    team_name VARCHAR(20) PRIMARY KEY
);

CREATE TABLE pull_requests (
    pull_request_id VARCHAR(50) PRIMARY KEY,
    pull_request_name VARCHAR(100) NOT NULL,
    author_id VARCHAR(50),
    status pull_request_status NOT NULL,
    createdAt TIMESTAMPTZ NOT NULL,
    mergedAt TIMESTAMPTZ NOT NULL,
);

CREATE TABLE pr_reviewers (
    pull_request_id INT NOT NULL REFERENCES PR(pull_request_id) ON DELETE CASCADE,
    user_id INT NOT NULL REFERENCES users(user_id),
    
    PRIMARY KEY (pull_request_id, user_id)
);