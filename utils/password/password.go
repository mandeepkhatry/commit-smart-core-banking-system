package password

import "golang.org/x/crypto/bcrypt"

type HashPasswordResp struct {
	HashedPassword string
	Error          error
}

func HashPassword(password string) HashPasswordResp {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return HashPasswordResp{Error: err}
	}
	return HashPasswordResp{HashedPassword: string(hashedPassword), Error: nil}
}

type ValidatePasswordParams struct {
	Password       string
	HashedPassword string
}

func ValidatePassword(validatePasswordParams ValidatePasswordParams) error {
	return bcrypt.CompareHashAndPassword([]byte(validatePasswordParams.HashedPassword), []byte(validatePasswordParams.Password))
}
