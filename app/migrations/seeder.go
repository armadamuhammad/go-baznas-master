package migrations

import "api/app/model"

var (
	user model.User
	role model.Role
)

// DataSeeds data to seeds
func DataSeeds() []interface{} {
	return []interface{}{
		user.Seed(),
		role.Seed(),
	}
}
