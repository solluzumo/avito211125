
---ДЛЯ ТЕСТА ФУНКИОНАЛА

INSERT INTO teams (team_name) VALUES
('backend'),
('frontend'),
('devops'),
('mobile');

INSERT INTO users (id, username, team_name, is_active) VALUES
('u1', 'alice',   'backend',  TRUE),
('u2', 'bob',     'backend',  TRUE),
('u3', 'charlie', 'backend',  FALSE),
('u4', 'david',   'frontend', TRUE),
('u5', 'eva',     'frontend', TRUE),
('u6', 'frank',   'frontend', TRUE),
('u7', 'helen',   'devops',   TRUE),
('u8', 'igor',    'devops',   FALSE),
('u9', 'jane',    'mobile',   TRUE),
('u10','kevin',   'mobile',   TRUE);


INSERT INTO pull_requests (
    id, pull_request_name, author_id, status, created_at, merged_at
) VALUES
('pr1',  'Add authentication module', 'u1',  'MERGED', '2025-01-10 12:30:00+00', '2025-01-11 15:00:00+00'),
('pr2',  'Fix backend cache bug',      'u2',  'OPEN',   '2025-02-01 09:00:00+00', NULL),
('pr3',  'Refactor backend services',  'u3',  'MERGED', '2025-02-15 11:45:00+00', '2025-02-16 13:20:00+00'),

('pr4',  'Improve frontend layout',    'u4',  'OPEN',   '2025-03-02 08:20:00+00', NULL),
('pr5',  'Add new frontend widgets',   'u5',  'MERGED', '2025-03-10 14:10:00+00', '2025-03-11 10:50:00+00'),
('pr6',  'Cleanup FE legacy code',     'u6',  'OPEN',   '2025-04-01 10:00:00+00', NULL),

('pr7',  'Implement CI pipelines',     'u7',  'MERGED', '2025-05-04 07:30:00+00', '2025-05-05 12:00:00+00'),
('pr8',  'Refactor deployment jobs',   'u8',  'OPEN',   '2025-05-12 15:00:00+00', NULL),

('pr9',  'Add push notifications',     'u9',  'MERGED', '2025-06-01 11:15:00+00', '2025-06-02 09:40:00+00'),
('pr10', 'App navigation redesign',    'u10', 'OPEN',   '2025-06-05 12:50:00+00', NULL);


INSERT INTO pr_reviewers (pull_request_id, user_id) VALUES
('pr1','u2'),
('pr1','u3'),
('pr2','u1'),
('pr2','u3'),
('pr3','u1'),
('pr3','u2');

INSERT INTO pr_reviewers (pull_request_id, user_id) VALUES
('pr4','u5'),
('pr4','u6'),
('pr5','u4'),
('pr5','u6'),
('pr6','u4'),
('pr6','u5');

INSERT INTO pr_reviewers (pull_request_id, user_id) VALUES
('pr7','u8'),
('pr8','u7');

INSERT INTO pr_reviewers (pull_request_id, user_id) VALUES
('pr9','u10'),
('pr10','u9');


-- ДЛЯ НАГРУЗКИ (PRCREATE)

-- --20 пользователей
-- INSERT INTO teams (team_name)
-- SELECT 'team' || g
-- FROM generate_series(1,20) AS g;

-- -- Создаём 200 пользователей, равномерно распределённых по 20 командам
-- INSERT INTO users (id, username, team_name, is_active)
-- SELECT 
--     'u' || i AS id,
--     'user' || i AS username,
--     'team' || ((i-1) % 20 + 1) AS team_name,
--     TRUE AS is_active
-- FROM generate_series(1,200) AS i;


-------------------------------------------------------


--- ДЛЯ НАГРУЗКИ (deactivateTEAM)

-- -- 200 команд
-- INSERT INTO teams (team_name)
-- SELECT 'team' || g
-- FROM generate_series(1,200) AS g;

-- -- 200 пользователей, равномерно распределённых по 200 командам
-- INSERT INTO users (id, username, team_name, is_active)
-- SELECT 
--     'u' || i AS id,
--     'user' || i AS username,
--     'team' || ((i-1) % 200 + 1) AS team_name,
--     TRUE AS is_active
-- FROM generate_series(1,200) AS i;
