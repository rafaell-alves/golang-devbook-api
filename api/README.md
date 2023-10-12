
# Devbook Golang - API

Api de uma rede social para devs.



## Stack utilizada

**Back-end feito em:** Golang


## Variáveis de Ambiente

Para rodar esse projeto, você vai precisar adicionar as seguintes variáveis de ambiente no seu .env

`DB_USUARIO`

`DB_SENHA`

`DB_NOME`

`API_PORT`

`SECRET_KEY`


## Instalação

Rode o codigo abaixo para criar o banco e as tabelas ou copie da pasta de sql o arquivo com nome de sql.sql para criar a base de dados e as tabelas:

```sql
CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS publicacoes;
DROP TABLE IF EXISTS seguidores;
DROP TABLE IF EXISTS usuarios;

CREATE TABLE usuarios(
    id int auto_increment primary key,
    nome varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    senha varchar(100) not null ,
    criadoEm timestamp default current_timestamp()
) ENGINE=INNODB;


DROP TABLE IF EXISTS seguidores;
-- Create the table in the specified schema
CREATE TABLE seguidores(
    usuario_id int not NULL, -- primary key column
    FOREIGN KEY (usuario_id) REFERENCES usuarios(id) ON DELETE CASCADE,
    seguidor_id int not null,
    FOREIGN KEY (seguidor_id) REFERENCES usuarios(id) ON DELETE CASCADE,
    PRIMARY KEY(usuario_id,seguidor_id)

) ENGINE=INNODB;


CREATE TABLE publicacoes(
    id int auto_increment primary key,
    titulo varchar(50) not null,
    conteudo varchar(300) not null,
    autor_id int not null ,
    FOREIGN KEY (autor_id) REFERENCES usuarios(id) on DELETE CASCADE,
    curtidas int DEFAULT 0 , 
    criadaEm timestamp default current_timestamp()
) ENGINE=INNODB;

```

Rode o projeto com go run main.go ou use o api.exe para rodar

```bash
  go run main.go
```
    
## Rodando localmente

Clone o projeto

```bash
  git clone https://link-para-o-projeto
```

Entre no diretório do projeto

```bash
  cd api
```

Inicie o servidor

```bash
  go run main.go
```

