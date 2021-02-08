# Imersão Full Stack && Full Cycle

Participe: https://imersao.fullcycle.com.br

## Api Bank-Api

Essa api é responsável por fazer a comunição o front-end dos bancos ao codepix recebendo requisições via Rest e realizando interações via gRPC e Apache Kafka com o codepix.

## Como executar

Utilizamos Docker para que todos os serviços que utilizaremos fiquem disponíveis.

- Faça o clone do projeto
- Tendo o docker instalado em sua máquina será necessário configurar um docker-compose para cada banco que quiser ter rodando, uma vez configurado rode o comando a seguir para cada arquivo
`docker-compose -f {nome do arquivo yaml} up `

### Como executar a aplicação
- Esta aplicação roda ao subir (up) o container.

**Importante:** Lembre-se de verificar se o Apache Kafka está inicializado.

### Serviços utilizados ao executar o docker-compose

- Aplicação principal
- Postgres
