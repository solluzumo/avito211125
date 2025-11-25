CREATE TYPE pull_request_status AS ENUM ('MERGED','OPEN');

CREATE TABLE users (
    id VARCHAR(50) PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    team_name VARCHAR(20),
    is_active BOOLEAN NOT NULL   
);

CREATE TABLE teams (
    team_name VARCHAR(20) PRIMARY KEY
);

CREATE TABLE pull_requests (
    id VARCHAR(50) PRIMARY KEY,
    pull_request_name VARCHAR(100) NOT NULL,
    author_id VARCHAR(50),
    status pull_request_status NOT NULL,
    created_at TIMESTAMPTZ,
    merged_at TIMESTAMPTZ
);


CREATE TABLE pr_reviewers (
    pull_request_id VARCHAR(50) NOT NULL REFERENCES pull_requests(id) ON DELETE CASCADE,
    user_id VARCHAR(50) NOT NULL REFERENCES users(id),
    
    PRIMARY KEY (pull_request_id, user_id)
);