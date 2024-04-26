## Sobre o Projeto
Essa aplicação é um esboço de uma ToDo list ou lista de tarefas, feita para gerenciar as atividades diárias de uma pessoa, ter idéia do que precisa ser feito e do que ainda não foi feito durante o dia.

## Tecnologias utilizadas
- **Go**: Linguagem de backend utilizada para desenvolver o servidor.
- **Fiber**: Framework de Go utilizado no desenvolvimento da aplicação do lado do backend, possibilitando a comunicação com o banco de dados.
- **Postgres**: Banco de dados SQL ideal para aplicações leves.
- **Docker**: Facilita o desenvolvimento e implantação de aplicações em ambientes isolados.
- **HTML e CSS**: Linguagens de marcação usada na construção de páginas estáticas.
- **Javascript e fetch API**: Linguagem e API usadas para dar dinamicidade às páginas e integrar backend com frontend.
   

## Instruções para Executar a Aplicação

Certifique-se de ter o Docker Desktop baixado no seu dispositivo e deixe o aplicativo aberto em segundo plano. Você pode baixá-lo [aqui](https://www.docker.com/products/docker-desktop/).

### Rodando a Aplicação

Para rodar a aplicação juntamente com o backend, siga estes passos:

1. No terminal, execute o comando:
   ```bash
   docker compose up
  
2. Em seguida, em outro terminal, execute o comando:
```bash
go run.main.go
```

### Acessando a aplicação
Se tudo ocorrer como esperado, a aplicação estará rodando e você poderá acessar as telas escrevendo a URL a seguir no navegador de sua preferência:
-Tela inicial:

`http://localhost:3000/Tela_inicial.html`

Você também pode substituir Tela_inicial pelo nome da tela que deseja visualizar.

