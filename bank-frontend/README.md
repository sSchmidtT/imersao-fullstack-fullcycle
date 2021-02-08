# Imersão Full Stack && Full Cycle

Participe: https://imersao.fullcycle.com.br

## Aplicação Front End do banco

Este é o front-end do banco utilizado para cadastrar e consultar novas chaves pix e realizar novas transações.

## Como executar

Utilizamos Docker para que todos os serviços que utilizaremos fiquem disponíveis.

- Faça o clone do projeto
- Tendo o docker instalado em sua máquina será necessário configurar um docker-compose para cada banco que quiser ter rodando, uma vez configurado rode o comando a seguir para cada arquivo:
`docker-compose -f {nome do arquivo .yaml} up -d

### Como executar a aplicação
- Acesse o container da aplicação executando: `docker-compose -f {nome do arquivo yaml} exec {nome do serviço definido no arquivo yaml} bash`
- Rode `npm run dev

**Importante:** Cada nova configuração feita através do docker-compose para subir uma nova instância da aplicação deverá ter uma Bank-api correnspondente rodando.

### Serviços utilizados ao executar o docker-compose

- Aplicação principal integrado ao Bank-api
