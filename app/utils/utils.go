// app/utils/utils.go

package utils

import (
    "log"
)

// ErrorHandler maneja los errores y registra un mensaje de error
func ErrorHandler(err error, message string) {
    if err != nil {
        log.Printf("%s: %v", message, err)
    }
}
