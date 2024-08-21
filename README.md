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
