package config

import (
    "errors"
    "fmt"
    "os"

    "gopkg.in/yaml.v3"
)

func LoadTestConfig() (map[string]any, error) {
    return LoadConfig("config/test.yml")
}

func LoadProdConfig() (map[string]any, error) {
    return LoadConfig("config/prod.yml")
}

func LoadEnvConfig() (map[string]any, error) {
    env := os.Getenv("API_ENV")
    switch env {
        case "test":
            return LoadTestConfig()
        case "prod":
            return LoadProdConfig()
    }
    return nil, errors.New("Invalid or missing API_ENV found.  Expected 'test', or 'prod'.")
}

func LoadConfig(path string) (map[string]any, error) {
    fmt.Println("[Config] Loading config file:", path)

    file, err := os.Open(path)
    if err != nil {
        fmt.Println("[Config] Unable to open config file.")
        return nil, err
    }
    defer file.Close()

    var config_data map[string]any
    decoder := yaml.NewDecoder(file)
    if err = decoder.Decode(&config_data); err != nil {
        fmt.Println("[Config] Unable to decode config file")
        return nil, err
    }

    return config_data, nil
}
