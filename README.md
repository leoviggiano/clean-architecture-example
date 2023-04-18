# Exemplo de Clean Architecture - GO

1. [Objetivo](#objetivo)
2. [Endpoints](#endpoints)
3. [Configurando o projeto](#configurando-o-projeto)

## Objetivo

Mostrar na prática uma forma de arquitetura limpa seguindo os padrões da linguagem Golang.

Esse projeto possui:

- Conexão com o banco de dados - `PostgreSQL`
- Conexão com o cache - `Redis`
- Clean Architecture seguindo as camadas Entity/Service/Repository
- Injeção de dependência nos `handlers HTTP` (services)

---

## Endpoints

Existem 5 endpoints que representam o mais básico CRUD (Create, Read, Update, Delete) para a demonstração do projeto.

O repositório conta com uma collection no Insomnia/Postman com os endpoints: [collection.json](/docs/collection.json)

É possível identificá-los no código:

> [cmd/app/app.go](/cmd/app/app.go#L43-L47)

### Create

```go
r.HandleFunc("/create-user", handler.CreateUser)`
```

### Read

```go
r.HandleFunc("/get-user", handler.GetUser)
r.HandleFunc("/get-users", handler.GetUsers)
```

### Update

```go
r.HandleFunc("/update-user", handler.UpdateUser)
```

### Delete

```go
r.HandleFunc("/delete-user", handler.DeleteUser)
```

## Configurando o projeto

### Dependências:

- Go >1.19
- [Docker Compose](https://furydocs.io/container-platform/1.0.9/guide/#/install-colima)

### Rodando o projeto:

1. É necessário que tenha criado a tabela `users`

   ```sh
   make migrate
   ```

2. Se já foi rodado a migração, rode o projeto com o comando:
   ```sh
   make run
   ```
