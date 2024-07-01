// configs/config.go

package configs

import (
    "os"
)

// Config estructura de configuración de la aplicación
type Config struct {
    Port string
    // Otros campos de configuración según sea necesario
}

// LoadConfig carga la configuración de variables de entorno o valores predeterminados
func LoadConfig() Config {
    port := os.Getenv("PORT")
    if port == "" {
        port = "3100" // Puerto por defecto si no se especifica
    }

    return Config{
        Port: port,
    }
}
