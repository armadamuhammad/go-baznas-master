package routes

import (
	"api/app/controller"
	"api/app/controller/account"
	"api/app/controller/balance"
	"api/app/controller/balancerecord"
	"api/app/controller/category"
	"api/app/controller/login"
	"api/app/controller/group"
	"api/app/controller/input"
	"api/app/controller/payment"
	"api/app/controller/role"
	"api/app/controller/transaction"
	"api/app/controller/user"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/spf13/viper"
)

// Handle all request to route to controller
func Handle(app *fiber.App) {
	app.Use(cors.New())
	services.InitDatabase()
	// services.InitRedis()

	api := app.Group(viper.GetString("ENDPOINT"))

	api.Get("/", controller.GetAPIIndex)
	api.Get("/info.json", controller.GetAPIInfo)

	// Account
	api.Post("/accounts", account.PostAccount)
	api.Get("/accounts", account.GetAccount)
	api.Put("/accounts/:id", account.PutAccount)
	api.Get("/accounts/:id", account.GetAccountID)
	api.Delete("/accounts/:id", account.DeleteAccount)

	// balance
	api.Post("/balances", balance.PostBalance)
	api.Get("/balances", balance.GetBalance)
	api.Put("/balances/:id", balance.PutBalance)
	api.Get("/balances/:id", balance.GetBalanceID)
	api.Delete("/balances/:id", balance.DeleteBalance)

	// Balance Record
	// api.Post("/balance-records", balancerecord.PostBalanceRecord)
	api.Get("/balance-records", balancerecord.GetBalanceRecord)
	// api.Put("/balance-records/:id", balancerecord.PutBalanceRecord)
	api.Get("/balance-records/:id", balancerecord.GetBalanceRecordID)
	api.Delete("/balance-records/:id", balancerecord.DeleteBalanceRecord)

	// Category
	api.Post("/categories", category.PostCategory)
	api.Get("/categories", category.GetCategory)
	api.Put("/categories/:id", category.PutCategory)
	api.Get("/categories/:id", category.GetCategoryID)
	api.Delete("/categories/:id", category.DeleteCategory)
	api.Get("/categories/:id/group", category.GetCategoryGroup)

	// Group
	api.Post("/groups", group.PostGroup)
	api.Get("/groups", group.GetGroup)
	api.Put("/groups/:id", group.PutGroup)
	api.Get("/groups/:id", group.GetGroupID)
	api.Delete("/groups/:id", group.DeleteGroup)

	// Login 
	api.Post("/logins", login.PostLogin)

	// Input
	api.Post("/inputs", input.PostInput)
	api.Get("/inputs", input.GetInput)
	api.Put("/inputs/:id", input.PutInput)
	api.Get("/inputs/:id", input.GetInputID)
	api.Delete("/inputs/:id", input.DeleteInput)

	// payment
	api.Post("/payments", payment.PostPayment)
	api.Get("/payments", payment.GetPayment)
	api.Put("/payments/:id", payment.PutPayment)
	api.Get("/payments/:id", payment.GetPaymentID)
	api.Delete("/payments/:id", payment.DeletePayment)

	// Role
	api.Post("/roles", role.PostRole)
	api.Get("/roles", role.GetRole)
	api.Put("/roles/:id", role.PutRole)
	api.Get("/roles/:id", role.GetRoleID)
	api.Delete("/roles/:id", role.DeleteRole)

	// Transaction
	api.Post("/transactions", transaction.PostTransaction)
	api.Get("/transactions", transaction.GetTransaction)
	api.Put("/transactions/:id", transaction.PutTransaction)
	api.Get("/transactions/:id", transaction.GetTransactionID)
	api.Delete("/transactions/:id", transaction.DeleteTransaction)

	// User
	api.Post("/users", user.PostUser)
	api.Get("/users", user.GetUser)
	api.Put("/users/:id", user.PutUser)
	api.Get("/users/:id", user.GetUserID)
	api.Delete("/users/:id", user.DeleteUser)
	api.Put("/users/:id/verify", user.PutUserVerify)

}
