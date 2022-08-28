package password

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {

	asserts := assert.New(t)
	password := "HelloPassword@123"
	hashedPasswordResp := HashPassword(password)

	asserts.Nil(hashedPasswordResp.Error)

	err := ValidatePassword(ValidatePasswordParams{Password: password, HashedPassword: hashedPasswordResp.HashedPassword})

	asserts.Nil(err)

	wrongPassword := "HelloPassword@123Wrong"

	err = ValidatePassword(ValidatePasswordParams{Password: wrongPassword, HashedPassword: hashedPasswordResp.HashedPassword})

	asserts.NotNil(err)

}
