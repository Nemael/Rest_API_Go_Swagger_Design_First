// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"database/sql"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	"Rest_API_Swagger_Design_First/models"
	"Rest_API_Swagger_Design_First/restapi/operations"

	_ "github.com/go-sql-driver/mysql"
)

//go:generate swagger generate server --target ../../Rest_API_Swagger_Design_First --name Books --spec ../swagger.yaml --principal interface{}

var (
	ConnectionString = "rest:password@tcp(localhost:3306)/books"
)

func getDB() (*sql.DB, error) {
	return sql.Open("mysql", ConnectionString)
}

func configureFlags(api *operations.BooksAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.BooksAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.GetBooksHandler = operations.GetBooksHandlerFunc(func(params operations.GetBooksParams) middleware.Responder {
		books := []*models.Book{}
		db, err := getDB()
		if err != nil {
			return operations.NewGetBooksDefault(500).WithPayload(&models.Error{Code: 500, Message: swag.String("Connection to database failed")})
		}
		rows, err := db.Query("SELECT id, title, author, quantity FROM books")
		if err != nil {
			return operations.NewGetBooksDefault(500).WithPayload(&models.Error{Code: 500, Message: swag.String("Connection to database failed")})
		}
		for rows.Next() {
			var book models.Book
			rows.Scan(&book.ID, &book.Title, &book.Author, &book.Quantity)
			books = append(books, &book)
		}
		return (operations.NewGetBooksOK().WithPayload(books))
	})

	api.GetBookHandler = operations.GetBookHandlerFunc(func(params operations.GetBookParams) middleware.Responder {
		db, err := getDB()
		if err != nil {
			return operations.NewGetBookDefault(500).WithPayload(&models.Error{Code: 500, Message: swag.String("Connection to database failed")})
		}
		id := params.ID //Path parameter
		var book models.Book
		row := db.QueryRow("SELECT id, title, author, quantity FROM books where id = ?", id)
		err = row.Scan(&book.ID, &book.Title, &book.Author, &book.Quantity)
		if err != nil {
			return operations.NewGetBookDefault(500).WithPayload(&models.Error{Code: 500, Message: swag.String("Book not found")})
		}
		return (operations.NewGetBookOK().WithPayload(&book))
	})

	api.AddBookHandler = operations.AddBookHandlerFunc(func(params operations.AddBookParams) middleware.Responder {
		db, err := getDB()
		if err != nil {
			return operations.NewAddBookDefault(500).WithPayload(&models.Error{Code: 500, Message: swag.String("Connection to database failed")})
		}
		newBook := params.Body
		_, err = db.Exec("INSERT INTO books (id, title, author, quantity) VALUES (?, ?, ?, ?)", newBook.ID, newBook.Title, newBook.Author, newBook.Quantity)
		if err != nil {
			return operations.NewAddBookDefault(500).WithPayload(&models.Error{Code: 500, Message: swag.String("Error inserting into the database")})
		}
		return (operations.NewAddBookCreated().WithPayload(newBook))
	})

	api.CheckoutBookHandler = operations.CheckoutBookHandlerFunc(func(params operations.CheckoutBookParams) middleware.Responder {
		var book models.Book
		db, err := getDB()
		if err != nil {
			return operations.NewCheckoutBookDefault(500).WithPayload(&models.Error{Code: 500, Message: swag.String("Connection to database failed")})
		}
		id := params.ID // Query parameter
		row := db.QueryRow("SELECT id, title, author, quantity FROM books where id = ?", id)
		err = row.Scan(&book.ID, &book.Title, &book.Author, &book.Quantity)
		if err != nil {
			return operations.NewCheckoutBookDefault(500).WithPayload(&models.Error{Code: 500, Message: swag.String("Book not found")})
		}

		if *book.Quantity <= int64(0) {
			return operations.NewCheckoutBookDefault(500).WithPayload(&models.Error{Code: 500, Message: swag.String("Book not available")})
		}
		*book.Quantity -= int64(1)

		db.QueryRow("UPDATE books SET quantity = ? WHERE id = ?", book.Quantity, book.ID)
		return (operations.NewCheckoutBookOK().WithPayload(&book))
	})

	api.ReturnBookHandler = operations.ReturnBookHandlerFunc(func(params operations.ReturnBookParams) middleware.Responder {
		var book models.Book
		db, err := getDB()
		if err != nil {
			return operations.NewReturnBookDefault(500).WithPayload(&models.Error{Code: 500, Message: swag.String("Connection to database failed")})
		}
		id := params.ID // Query parameter
		row := db.QueryRow("SELECT id, title, author, quantity FROM books where id = ?", id)
		err = row.Scan(&book.ID, &book.Title, &book.Author, &book.Quantity)
		if err != nil {
			return operations.NewReturnBookDefault(500).WithPayload(&models.Error{Code: 500, Message: swag.String("Book not found")})
		}

		*book.Quantity += int64(1)

		db.QueryRow("UPDATE books SET quantity = ? WHERE id = ?", book.Quantity, book.ID)
		return (operations.NewReturnBookOK().WithPayload(&book))
	})

	if api.AddBookHandler == nil {
		api.AddBookHandler = operations.AddBookHandlerFunc(func(params operations.AddBookParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.AddBook has not yet been implemented")
		})
	}
	if api.CheckoutBookHandler == nil {
		api.CheckoutBookHandler = operations.CheckoutBookHandlerFunc(func(params operations.CheckoutBookParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.CheckoutBook has not yet been implemented")
		})
	}
	if api.GetBookHandler == nil {
		api.GetBookHandler = operations.GetBookHandlerFunc(func(params operations.GetBookParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetBook has not yet been implemented")
		})
	}
	if api.GetBooksHandler == nil {
		api.GetBooksHandler = operations.GetBooksHandlerFunc(func(params operations.GetBooksParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetBooks has not yet been implemented")
		})
	}
	if api.ReturnBookHandler == nil {
		api.ReturnBookHandler = operations.ReturnBookHandlerFunc(func(params operations.ReturnBookParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.ReturnBook has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
