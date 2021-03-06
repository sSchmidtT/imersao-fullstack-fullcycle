# Imersão Full Stack && Full Cycle

Participe: https://imersao.fullcycle.com.br

## Microsserviço CodePix

Esse microsserviço tem o objetivo de ser um hub de transações entre os bancos que simularemos durante o projeto.

## Como executar

Utilizamos Docker para que todos os serviços que utilizaremos fiquem disponíveis.

- Faça o clone do projeto
- Tendo o docker instalado em sua máquina apenas execute:
`docker-compose up -d`

### Como executar a aplicação
- Acesse o container da aplicação executando: `docker exec -it codepix_app bash`
- Rode `go run main.go fixtures` - este comando irá realizar o cadastro de alguns bancos e algumas contas ao inicializar a aplicação
- Rode `go run main.go all` - este comando irá disponibilizar o **gRPC** e o **Process e Consumer** do **Apache Kafka** deixando a aplicação pronta para receber e processar requisições.

### Serviços utilizados ao executar o docker-compose

- Aplicação principal
- Postgres
- PgAdmin

## Comando para gerar PB do gRPC
protoc --go_out=application/grpc/pb --go_opt=paths=source_relative --go-grpc_out=application/grpc/pb --go-grpc_opt=paths=source_relative --proto_path=application/grpc/protofiles application/grpc/protofiles/*.proto

## Comand para abrir client do gRPC
evans -r repl
**importante:** para que este client esteja disponível a imagem do docker tem que ter o evans instalado e no StartGrpcServer tem que definir o reflection do gRPC  - - reflection.Register(grpcServer)
