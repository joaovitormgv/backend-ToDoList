-- Descrição: Este arquivo contém os comandos SQL para criar o banco de dados e tabelas.
CREATE DATABASE todolist;

\c todolist;

CREATE TABLE IF NOT EXISTS usuario (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    senha VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS ToDos (
    userId INT NOT NULL,
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    completed BOOLEAN NOT NULL,
    FOREIGN KEY (userId) REFERENCES usuario(id)
);