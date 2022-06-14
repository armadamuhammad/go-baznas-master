package lib

import (
	"testing"

	"github.com/gofiber/fiber/v2/utils"
)

func TestPasswordEncrypt(t *testing.T) {
	raw := "password"
	salt := "salt"
	key := "CIPHER_SECRETKEY_MUST_HAVE_32BIT"

	utils.AssertEqual(t, false, PasswordEncrypt(raw, salt, key) == "")
}

func TestPasswordCompare(t *testing.T) {
	raw := "password"
	salt := "salt"
	key := "CIPHER_SECRETKEY_MUST_HAVE_32BIT"
	hashed := PasswordEncrypt(raw, salt, key)
	utils.AssertEqual(t, true, PasswordCompare(hashed, raw, salt, key))
}
