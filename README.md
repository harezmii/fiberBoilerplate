#<font color="cyan">Fiber Boilerplate</font>
###Scalable rest app fiber

##<font color="yellow">Tech Stack</font>
```

Go 1.17
Fiber 2.35
Docker
Mysql
Ent
Validation
```

##<font color="yellow"> Directory Structure</font>

```
+---build
|       Dockerfile
|
+---cmd
|       main.go
|
+---pkg
|   +---app
|   |       app.go
|   |
|   +---docs
|   |       docs.go
|   |       swagger.json
|   |       swagger.yaml
|   |
|   +---ent
|   |   |   client.go
|   |   |   config.go
|   |   |   connection.go
|   |   |   context.go
|   |   |   ent.go
|   |   |   generate.go
|   |   |   mutation.go
|   |   |   runtime.go
|   |   |   tx.go
|   |   |   user.go
|   |   |   user_create.go
|   |   |   user_delete.go
|   |   |   user_query.go
|   |   |   user_update.go
|   |   |
|   |   +---enttest
|   |   |       enttest.go
|   |   |
|   |   +---hook
|   |   |       hook.go
|   |   |
|   |   +---migrate
|   |   |       migrate.go
|   |   |       schema.go
|   |   |
|   |   +---predicate
|   |   |       predicate.go
|   |   |       
|   |   +---runtime
|   |   |       runtime.go
|   |   |
|   |   +---schema
|   |   |       user.go
|   |   |
|   |   \---user
|   |           user.go
|   |           where.go
|   |
|   +---model
|   |       model.go
|   |
|   +---repository
|   |   |   repo.go
|   |   |
|   |   \---user
|   |           user.go
|   |
|   +---route
|   |       route.go
|   |
|   \---validation
|           validation.go
```

##<font color="yellow">Build</font>
#### docker container build -t server . 

##<font color="yellow">Usage</font>
```
go mod download 

OR

go run ./pkg/cmd/main.go

OR

air
```