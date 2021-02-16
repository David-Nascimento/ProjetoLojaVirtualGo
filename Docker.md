# Ambiente de testes

## Baixar as imagens
docker pull postgres
docker pull dpage/pgadmin4

### Criar a rede  Docker
docker network create --driver bridge skynet

### Subir o Banco de Dados
docker run --name pgdb --network=skynet -e "POSTGRES_PASSWORD=123456" -p 5432:5432 -v var/lib/postgresql/data -d postgres

### Subir o Administrador do Banco de Dados (PgAdmin)
docker run --name pgadmin --network=skynet -p 15432:80 -e "PGADMIN_DEFAULT_EMAIL=root@tests.io" -e "PGADMIN_DEFAULT_PASSWORD=123456" -d dpage/pgadmin4

### Importante:​
### Quando você reiniciar o seu computador, certifique-se que o Docker esteja online e suba containers​ novamente:

docker start pgdb
docker start pgadmin

### Se alguma coisa der errado e for necessário refazer a aula, voce poderá remover os containers com os seguintes comandos:

docker rm -f ​pgdb
docker rm -f ​pgadmin