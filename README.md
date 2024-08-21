# Goingressos Ticket Handling

## Descrição

O **Goingressos Ticket Handling** é um serviço em Go que roda em segundo plano e executa operações automáticas relacionadas à devolução de ingressos. Ele é projetado para funcionar de maneira contínua, processando solicitações de devolução de ingressos de forma automática em intervalos regulares, interagindo diretamente com o banco de dados.

## Funcionalidades

- Processamento automático de devoluções de ingressos.
- Execução em segundo plano com intervalos configuráveis.
- Log de atividades para monitoramento.
- Conexão com banco de dados MySQL para manipulação dos ingressos.

## Requisitos

- Go 1.18+ (https://golang.org/dl/)
- MySQL (versão 5.7 ou superior)
- Git (opcional, para controle de versão)
  
### Bibliotecas utilizadas

- [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql): Driver MySQL para Go.
- [godotenv](https://github.com/joho/godotenv): Para carregar variáveis de ambiente a partir de arquivos `.env`.
- [logrus](https://github.com/sirupsen/logrus): Para log estruturado.

## Configuração

### Clonando o Repositório

Primeiro, clone o repositório para o seu ambiente local:

```bash
git clone https://github.com/DukaSiqueira/goingressos_ticket_handling.git
cd goingressos_ticket_handling
```

### Instalando as Dependências

Depois de clonar o repositório, instale as dependências necessárias para o projeto:

```bash
go mod tidy
```

### Configurando o Banco de Dados

Certifique-se de ter um banco de dados MySQL em execução. Crie um banco de dados para o projeto:

```sql
CREATE DATABASE goingressos;
```

Em seguida, edite o arquivo .env na raiz do projeto com as informações de conexão do banco de dados:

```bash
cp .env.example .env
```

### Edite o arquivo .env com os detalhes da sua configuração de banco de dados:

```plaintext
DB_USERNAME=root
DB_PASSWORD=sua_senha
DB_HOST=127.0.0.1
DB_PORT=3306
DB_DATABASE=goingressos
```
### Executando o Projeto

Para rodar o projeto, basta executar o seguinte comando:

```bash
go run cmd/goingressos_ticket_handling/main.go
```

O serviço começará a rodar em segundo plano, processando devoluções de ingressos a cada intervalo de tempo definido.
Estrutura do Projeto

```plaintext
goingressos_ticket_handling/
├── cmd/
│   └── goingressos_ticket_handling/
│       └── main.go
├── config/
│   └── config.go
├── internal/
│   ├── database/
│   │   └── mysql.go
│   └── tickets/
│       ├── handler.go
│       └── models.go
├── pkg/
│   └── logger/
│       └── logger.go
├── .env.example
├── README.md
├── go.mod
└── go.sum
    cmd/goingressos_ticket_handling/main.go: Ponto de entrada do projeto.
    config/: Carrega e gerencia as configurações do projeto.
    internal/database/: Contém a lógica de conexão com o banco de dados MySQL.
    internal/tickets/: Contém a lógica de devolução de ingressos e os modelos de dados.
    pkg/logger/: Implementa o logger para registrar logs estruturados.
    .env.example: Exemplo de arquivo .env para configuração do banco de dados.
```

### Contribuições

Contribuições são bem-vindas! Para contribuir com este projeto, siga estas etapas:
    Faça um fork do projeto.
    Crie uma nova branch (git checkout -b feature/nome-da-sua-feature).
    Faça suas modificações e adicione commits (git commit -m 'Adicionar nova feature').
    Envie suas mudanças para o seu repositório fork (git push origin feature/nome-da-sua-feature).
    Crie um Pull Request explicando as mudanças feitas.

### Licença
Este projeto está licenciado sob a MIT License.
