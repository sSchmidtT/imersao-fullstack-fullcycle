# Imersão Full Stack && Full Cycle

Participe: https://imersao.fullcycle.com.br

Este projeto simula as operações de transferências através de uma chave única (pix) onde é possível realizar transações entre contas e bancos diversos, para isso foi desenvolvido a aplicação front-end do banco, a api para integração do banco e um microsserviço utilizando gRPC responsável por armazenar, validar e distribuir as requisições entre bancos diversos..

## Como executar

Utilizamos Docker para que todos os serviços que utilizaremos fiquem disponíveis.

- Faça o clone do projeto
- Tendo o docker instalado execute os containers de cada aplicação.`

### Como executar a aplicação
- Inicialize os containers do Apache Kafka.
- Inicialize os containers do CodePix (aplicação em GO e Postgres).
- Inicialize cada Bank-api desejada (é possível configurar várias para simular bancos diferente).
- Inicialize cada Bank-frontend desejado (para cada Bank-api rodando deverá ter uma frontend configurado).

