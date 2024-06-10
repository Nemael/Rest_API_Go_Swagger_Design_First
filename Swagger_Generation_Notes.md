# Base swagger file
```
swagger: "2.0"
info:
  title: Books API
  description: A simple Book management API
  version: 1.0.0
basePath: /api/v1
schemes:
  - http
paths:
  /books:
    get:
      summary: Lists all books
      operationId: getBooks
      responses:
        200:
          description: An array of books
          schema:
            type: array
            items:
              $ref: '#/definitions/Book'
        default:
          description: generic error response
          schema:
            $ref: "#/definitions/error"
    post:
      summary: Adds a book
      operationId: addBook
      parameters:
        - name: body
          in: body
          schema:
            $ref: "#/definitions/Book"
      responses:
        201:
          description: Created
          schema:
            $ref: "#/definitions/Book"
        default:
          description: error
          schema:
            $ref: "#/definitions/error"
  /book/{id}:
    parameters:
    - type: integer
      format: int64
      name: id
      in: path
      required: true
    get:
      summary: Return a book from ID
      operationId: getBook
      parameters:
        - name: body
          in: body
          schema:
            $ref: "#/definitions/Book"
      responses:
        200:
          description: A book
          schema:
            $ref: '#/definitions/Book'
        default:
          description: generic error response
          schema:
            $ref: "#/definitions/error"
  /checkout:
    patch:
      operationId: checkoutBook
      parameters:
      - type: integer
        format: int64
        name: id
        in: query
      responses:
        200:
          description: OK
          schema:
            $ref: "#/definitions/Book"
        default:
          description: generic error response
          schema:
            $ref: "#/definitions/error"
  /return:
    patch:
      operationId: returnBook
      parameters:
      - type: integer
        format: int64
        name: id
        in: query
      responses:
        200:
          description: OK
          schema:
            $ref: "#/definitions/Book"
        default:
          description: generic error response
          schema:
            $ref: "#/definitions/error"

definitions:
  Book:
    type: object
    required:
      - id
      - title
      - author
      - quantity
    properties:
      id:
        type: integer
        format: int64
      title:
        type: string
      author:
        type: string
      quantity:
        type: integer
        format: int64
  error:
    type: object
    required:
      - message
    properties:
      code:
        type: integer
        format: int64
      message:
        type: string
```

# Commands:
- `go get -u github.com/go-swagger/go-swagger/cmd/swagger`
- `go mod init [NAME_OF_FOLDER]`
- `swagger generate server -A Rest_API_Swagger_Design_First -f ./swagger.yaml`
- `swagger generate client -A Rest_API_Swagger_Design_First -f ./swagger.yaml`
- `go get github.com/go-openapi/runtime`
- `go get github.com/jessevdk/go-flags`
- OR `go mod tidy`
- `go build ./cmd/todo-server/main.go` to build the server
- `./main.go` to run the server
- `go run ./cmd/books-server/main.go`
- `curl -i http://127.0.0.1:34839/api/v1/books` to test the API, verify that /api/v1 is included
- `curl -i http://127.0.0.1:8080/api/v1/books -d '{"author":"Adeline","id":4,"quantity":2,"title":"Colors for adults"}' -H 'Content-Type: application/json'` to add a new book
- `curl -i http://127.0.0.1:8080/api/v1/checkout?id=1 --request PATCH`

# Swagger.yaml setup
- Great example at `https://goswagger.io/go-swagger/tutorial/todo-list`

# Validation
- Validate file using `swagger vaidate swagger.yaml`
- Re-validate it online on `https://editor.swagger.io/`

# After generation
- Edit `restapi/configure_books.go` The file is safe to edit. You can set mysql or data storage in there. There you can access the handlers