package utils

import "strings"

func IsEmail(identifier string) bool {
    return strings.Contains(identifier, "@")
}
