INSERT INTO guilds (name, description) 
VALUES 
('Guild_1', 'Description of Guild 1'),
('Guild_2', 'Description of Guild 2'),
('Guild_3', 'Description of Guild 3');

INSERT INTO members (guild_id, name, role)
VALUES 
(1, 'Member_1', 'Role_1'),
(1, 'Member_2', 'Role_2'),
(2, 'Member_3', 'Role_3');
