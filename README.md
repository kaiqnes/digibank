# Digibank

## Tecnologias
- Golang 1.18
- Gin Gonic v1.7.7
- GORM v1.23.5

## Descrição

A Digibank-app tem como função simular uma rotina de transações onde cada portador de cartão (cliente) possui uma conta com seus dados.
A cada operação realizada pelo cliente, uma transação é criada e associada à respectiva conta.
Cada transação possui um tipo (compra à vista, compra parcelada, saque ou pagamento), um valor e uma data de criação.
Transações de tipo **compra e saque** são registradas com **valor negativo**, enquanto transações de **pagamento** são registradas com **valor positivo**.

## Iniciando a aplicação
Para iniciar a aplicação, execute algum dos comandos abaixo:


```
make run
```
> O comando 'make run' iniciará os containers necessários para que a aplicação fique no ar.



```
make stop
```
> O comando 'make stop' encerrará a atividade dos containers.


## Testes e documentação
Para verificar se a aplicação está no ar, existe uma URL de [HEALTH_CHECK](http://localhost:8080/health).

Para os testes, o projeto conta com uma [POSTMAN_COLLECTION](https://github.com/kaiqnes/digibank/blob/master/Digibank.postman_collection.json) para auxiliar nas requisições;

Também existe um [SWAGGER](http://localhost:8080/swagger/index.html) para que possa ser observado em maiores detalhes os retornos possiveis de cada rota.
