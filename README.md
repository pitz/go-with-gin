### Learning GO
This project is a simple API written in Go using the Gin framework. 
It serves as learning material to help understand the basics of Go and how to build a web service with the Gin framework.

### Design:
This project is inspired by the Diplomat Architects concept from Nubank, which simplifies the Hexagonal Architecture. 

Basically, the service is divided into three main parts:
- **In**: Handles all external interactions with us.
- **Internal**: Domain and business logic.
- **Models**: Models and internal schemas.
- **Out**: Handles all calls to the external world (databases are also part of the external world).

### Caveats:
This is a sandbox project intended to learn Golang. 
Therefore, it may not follow best practices, as I am still in the process of learning it :)

### Resources:
- https://go.dev/doc/tutorial/web-service-gin
- https://building.nubank.com.br/working-with-clojure-at-nubank/
