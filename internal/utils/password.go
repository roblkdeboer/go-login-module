package utils

import (
	"regexp"

	"github.com/roblkdeboer/login-module/internal/errors"
)

// ValidatePasswordStrength returns true if the password meets the required strength criteria.
func IsValidPassword(password string) ( error) {
	if len(password) < 8 {
        return &errors.ValidationError{
            Message: "Password should be 8 characters long",
        }
    }

	re, err := regexp.MatchString("([a-z])+", password)
    if err != nil {
        return err
    }
    if !re {
        return &errors.ValidationError{
            Message: "Password should contain atleast one lower case character",
        }
    }

	re, err = regexp.MatchString("([A-Z])+", password)
    if err != nil {
        return err
    }
    if !re {
        return &errors.ValidationError{
            Message: "Password should contain atleast one upper case character",
        }
    }

	re, err = regexp.MatchString("([0-9])+", password)
    if err != nil {
        return err
    }
    if !re {
        return &errors.ValidationError{
            Message: "Password should contain atleast one number",
        }
    }

	re, err = regexp.MatchString("([!@#$%^&*.?-])+", password)
    if err != nil {
        return err
    }
    if !re {
        return &errors.ValidationError{
            Message: "Password should contain atleast one special character",
        }
    }

	return nil
}