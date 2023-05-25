# Importador CSV para MySQL em Golang

Este projeto é um simples importador de CSV para MySQL escrito em Go. Ele lê um arquivo CSV e importa os dados para uma tabela MySQL. 

## Características

- Lê dados de um arquivo CSV.
- Cria a tabela se ela não existir.
- Importa dados do CSV para a tabela MySQL.
- Conexão com MySQL usando Docker em localhost.
- O usuário pode escolher a pasta do arquivo, nome da tabela e outros campos pertinentes.

## Pré-requisitos

Para executar este projeto, você precisa ter:

- Go 1.16 ou superior
- Docker
- MySQL

## Como usar

1. Clone este repositório em sua máquina local.
2. Execute o MySQL em Docker no localhost.
3. A partir do diretório do projeto, execute o comando `go run cmd/main.go`.
4. Siga as instruções na tela para fornecer as informações necessárias, como host do banco de dados, porta, usuário, senha, nome do banco de dados, nome da tabela e caminho do arquivo CSV.
5. O programa criará a tabela se ela não existir e importará os dados do CSV para a tabela MySQL.

## Estrutura do Projeto

- `pkg/database.go`: Contém a configuração do banco de dados e a função para inicializar o banco de dados.
- `pkg/tables.go`: Contém as funções para criar a tabela e importar dados do CSV para a tabela.
- `cmd/main.go`: O ponto de entrada do programa. Solicita ao usuário as informações necessárias e chama as funções apropriadas para inicializar o banco de dados, criar a tabela e importar os dados.

## Contribuição

Contribuições são sempre bem-vindas. Sinta-se à vontade para abrir uma issue ou enviar um pull request.

## Licença

Este projeto está licenciado sob a licença MIT - veja o arquivo `LICENSE.md` para mais detalhes.

