# Cadastro de super


## Tabela de conteúdos
  * [Sobre](#Sobre)
  * [Como usar](#como-usar)
  * [Tecnologias](#tecnologias)

## Sobre
Essa api permite adicionar em um banco de dados dados de um super herói/vilão quando fornecido o nome do super. Para obter os dados do super a api https://superheroapi.com é consultada. Além disso, é possível encontrar um super pelo nome e pelo uuid, visualizar todos os dados cadastrados e deletar um super. O banco de dados utilizado foi o PostgreSQL.

## Como usar

Na raiz do projeto está disponível o arquivo super.sql que tem o código para gerar o banco de dados utilizado. Após gerar o banco de dados, use o código disponível dentro do arquivo superhero.sql para gerar a tabela. Para que a conexão com o banco de dados seja feita acesse o arquivo connection.go na pasta model e modifique os seguintes dados de acordo com seus parâmetros: host, port, user, password.

Os endpoints disponíveis são:
- "/new" - POST - adiciona um super. O nome do super deve ser fornecido no body da requisição como Name.
- "/" - GET - retorna todos os supers cadastrados;
- "/getbyname/:name" - GET - retorna o super que tem o nome fornecido;
- "/getbyuuid/:uuid" - GET - retorna o super que tem o uuid fornecido;
- "/delete/:uuid" - DELETE - deleta o super que tem o uuid fornecido.

## Tecnologias
- Go versão 1.16.2;
- [Fiber](https://github.com/gofiber/fiber) versão 2.6.0 por ser uma dependência parecida com o Express do Node que tenho maior familiaridade;
- [Uuid](https://github.com/google/uuid) versão 1.2.0 para gerar um uuid para cada super;
- [Pq](https://github.com/lib/pq) versão 1.10.0 para fazer a conexão com o banco de dados.