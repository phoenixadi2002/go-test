package config

import (
    "github.com/joho/godotenv"
    "github.com/spf13/viper"
    "log"
)

type Config struct {
    Port       string
    DBUser     string
    DBPassword string
    DBName     string
    DBHost     string
    DBPort     string
}

func LoadConfig() *Config {
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found")
    }

    viper.AutomaticEnv() // Bind env vars

    return &Config{
        Port:       viper.GetString("PORT"),
        DBUser:     viper.GetString("DB_USER"),
        DBPassword: viper.GetString("DB_PASSWORD"),
        DBName:     viper.GetString("DB_NAME"),
        DBHost:     viper.GetString("DB_HOST"),
        DBPort:     viper.GetString("DB_PORT"),
    }
}