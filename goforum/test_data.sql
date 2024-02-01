INSERT INTO forum(title, description, is_public) VALUES ('forum_test', 'this is a test for the forum', True);

/*passwords and emails are encrypted in prod database */
INSERT INTO goforum_user(id, username, password, email) VALUES (1, 'user_test', 'strongpassword', 'user_test@example.com');
