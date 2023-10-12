insert into usuarios (nome,nick,email,senha)
VALUES
("Usuario 1","Usuario1","usuarios1@gmail.com","$2a$10$t9SggSYMFV27igB/i5e78ebpB.D27Tdnh3q1./nxu0ZuBRD8nmMce"),
("Usuario 2","Usuario2","usuarios2@gmail.com","$2a$10$t9SggSYMFV27igB/i5e78ebpB.D27Tdnh3q1./nxu0ZuBRD8nmMce"),
("Usuario 3","Usuario3","usuarios3@gmail.com","$2a$10$t9SggSYMFV27igB/i5e78ebpB.D27Tdnh3q1./nxu0ZuBRD8nmMce"),
("Usuario 4","Usuario4","usuarios4@gmail.com","$2a$10$t9SggSYMFV27igB/i5e78ebpB.D27Tdnh3q1./nxu0ZuBRD8nmMce")


insert into seguidores(usuario_id,seguidor_id)
VALUES
(1,3),
(3,5),
(3,4),
(3,6),
(3,1),
(4,3),
(5,3),
(1,4),
(1,5),
(4,5)


insert into publicacoes (titulo,conteudo,autor_id)
VALUES
("Publicação teste para mais usuarios", "Essa é a publicação desse usuario",1),
("Publicação teste para mais usuarios", "Essa é a publicação desse usuario",2),
("Publicação teste para mais usuarios", "Essa é a publicação desse usuario",3),
("Publicação teste para mais usuarios", "Essa é a publicação desse usuario",4),
("Publicação teste para mais usuarios", "Essa é a publicação desse usuario",5),
("Publicação teste para mais usuarios", "Essa é a publicação desse usuario",6),
("Publicação teste para mais usuarios", "Essa é a publicação desse usuario",7),
("Publicação teste para mais usuarios", "Essa é a publicação desse usuario",5),
("Publicação teste para mais usuarios", "Essa é a publicação desse usuario",6),
("Publicação teste para mais usuarios", "Essa é a publicação desse usuario",6);