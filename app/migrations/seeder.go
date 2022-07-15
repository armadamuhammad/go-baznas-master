package migrations

import "api/app/model"

var (
	user model.User
	role model.Role
	payment model.Payment
	balance model.Balance
)

// DataSeeds data to seeds
func DataSeeds() []interface{} {
	return []interface{}{
		user.Seed(),
		role.Seed(),
		payment.Seed(),
		balance.Seed(),
	}
}
