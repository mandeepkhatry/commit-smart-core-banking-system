package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoundFloat(t *testing.T) {
	asserts := assert.New(t)
	float1 := 300.54678

	serializedAmount1 := SerializeAmount(float1)
	asserts.Equal(int64(300546), serializedAmount1)

	deserializedAmount1 := DeserializeAmount(serializedAmount1)
	asserts.Equal(300.546, deserializedAmount1)

	float2 := 300.5

	serializedAmount2 := SerializeAmount(float2)
	asserts.Equal(int64(300500), serializedAmount2)

	deserializedAmount2 := DeserializeAmount(serializedAmount2)
	asserts.Equal(300.5, deserializedAmount2)

	offset1 := GetOffset(1, 100)
	asserts.Equal(int64(0), offset1)

	offset2 := GetOffset(2, 100)
	asserts.Equal(int64(100), offset2)

}

// func Test(t *testing.T) {

// 	asserts := assert.New(t)
// 	password := "HelloPassword@123"
// 	hashedPasswordResp := HashPassword(password)

// 	asserts.Nil(hashedPasswordResp.Error)

// 	err := ValidatePassword(ValidatePasswordParams{Password: password, HashedPassword: hashedPasswordResp.HashedPassword})

// 	asserts.Nil(err)

// 	wrongPassword := "HelloPassword@123Wrong"

// 	err = ValidatePassword(ValidatePasswordParams{Password: wrongPassword, HashedPassword: hashedPasswordResp.HashedPassword})

// 	asserts.NotNil(err)

// }
