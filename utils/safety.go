package utils

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
)

// Define salt size
const saltSize int = 64

// Generate _saltSize_ bytes using the Cryptographically secure
// pseudorandom number generator (CSPRNG) in the crypto.rand package
func _GenerateRandomSalt(saltSize int) []byte {
	var salt = make([]byte, saltSize)
	_, err := rand.Read(salt[:])

	if err != nil {
		panic(err)
	} else {
		return salt
	}
}

// Combine password and salt then hash them using the SHA-512
// hashing algorithm and then return the hashed password
// as a base64 encoded string
func _HashPassword(password string, salt []byte) string {
	// Convert password string to byte slice
	var passwordBytes = []byte(password)

	// Create sha-512 hasher
	// and hash pwd + salt
	var sha512Hasher = sha512.New()
	sha512Hasher.Write(append(passwordBytes, salt...))
	var hashedPasswordBytes = sha512Hasher.Sum(nil)

	// Convert the hashed password to a base64 encoded string
	var base64EncodedPasswordHash = base64.URLEncoding.EncodeToString(hashedPasswordBytes)

	return base64EncodedPasswordHash
}

// calculates both the hash and the salt
func HashAndSaltPassword(password string) (string, string) {
	var salt = _GenerateRandomSalt(saltSize)
	var passwordHash = _HashPassword(password, salt)

	return passwordHash, base64.URLEncoding.EncodeToString(salt)
}

// Check if two passwords match
func DoPasswordsMatch(plainPassword string, hashedPassword string, salt string) bool {
	var currPasswordHash = _HashPassword(plainPassword, []byte(salt))
	return hashedPassword == currPasswordHash
}

func ConvertBytesToString(data []byte) string {
	return base64.URLEncoding.EncodeToString(data)
}
