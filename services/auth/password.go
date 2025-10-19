package auth

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(passwordHash), nil
}

//
func ComparePassword(hashedpassword string, plain []byte) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedpassword), plain)
	return err == nil
}
