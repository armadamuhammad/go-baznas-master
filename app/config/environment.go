package config

// Environment variables
var Environment = map[string]interface{}{
	"app_id":          "6A6EF734-F557-4F7E-9C2A-32CD28E43420",
	"app_version":     "v1.0.0",
	"app_name":        "unknown",
	"app_description": "",
	"port":            8000,
	"endpoint":        "/api/v1",
	"environment":     "development",
	"db_host":         "postgres",
	"db_port":         5432,
	"db_user":         "postgres",
	"db_pass":         "postgres",
	"db_name":         "postgres",
	"db_table_prefix": "",
	"redis_host":      "redis",
	"redis_port":      6379,
	"redis_pass":      "",
	"redis_index":     0,
	"prefork":         false,
}
