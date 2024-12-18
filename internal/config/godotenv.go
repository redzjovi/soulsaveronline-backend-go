package config

import "github.com/joho/godotenv"

func NewGodotenv() {
	godotenv.Load(".env")
}
