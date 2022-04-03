-- Start ;
INSERT INTO roles(name) VALUES ('admin');
INSERT INTO roles(name) VALUES ('premium');
INSERT INTO roles(name) VALUES ('staff');

INSERT INTO users(username,password,created_at,updated_at,last_login) VALUES ('evai','123321',(select now()),(select now()),(select now()));
INSERT INTO users_roles(user_id,role_id) VALUES ((select u.id from users u where u.username = 'evai'),(select r.id from roles r where r.name = 'admin'));

INSERT INTO languages(name, short_name) VALUES('Russia','rus'), ('English','eng');

INSERT INTO initials(name)   VALUES('b'),('p'),('m'),('f'),('d'),('t'),('n'),('l'),('g'),('k'),('h'),('j'),('q'),('x'),('z'),('c'),('s'),('zh'),('ch'),('sh'),('r');
INSERT INTO finals(name)     VALUES('a'),('o'),('e'),('i'),('u'),('ü'),('ai'),('ao'),('an'),('ang'),('ou'),('ong'),('ei'),('en'),('eng'),('er'),('ia'),('iao'),('ian'),('iang'),('ie'),('iu'),('in'),('ing'),('iong'),('ua'),('uai'),('uan'),('uang'),('uo'),('ui'),('un'),('üe'),('üan'),('ün'),('ueng');
-- End ;

