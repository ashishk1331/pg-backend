package config

import "os"

// Declare the variables with proper types.
var JwtSecret []byte
var ResetSecret []byte

// Initialize the variables in this function.
func InitJWTSecret() {
	JwtSecret = []byte(os.Getenv("JWT_SECRET"))
	ResetSecret = []byte(os.Getenv("RESET_JWT_SECRET"))
}
