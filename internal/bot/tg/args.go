package tg

import (
	"errors"
	"strings"
)

func parseAddRoleArgs(args []string) (string, error) {
	if len(args) < 1 {
		return "", errors.New("*role name* is required")
	}
	roleName := strings.TrimSpace(args[0])
	if roleName == "" {
		return "", errors.New("*role name* is required")
	}
	return roleName, nil
}

func parseRemoveRoleArgs(args []string) (string, error) {
	if len(args) < 1 {
		return "", errors.New("*role name* is required")
	}
	return args[0], nil
}

func parseRoleUserArgs(args []string) (string, string, error) {
	if len(args) < 2 {
		return "", "", errors.New("*role name* and *user tag* are required")
	}
	roleName := strings.TrimSpace(args[0])
	if roleName == "" {
		return "", "", errors.New("*role name* is required")
	}
	userTgId := strings.TrimSpace(args[1])
	if len(userTgId) < 2 {
		return "", "", errors.New("*user tag* is required")
	}
	return roleName, userTgId, nil
}

func parseAddSuperuserArgs(args []string) (string, error) {
	if len(args) < 1 {
		return "", errors.New("*token* is required")
	}
	token := strings.TrimSpace(args[0])
	if token == "" {
		return "", errors.New("*token* is required")
	}
	return token, nil
}
