package main

import (
	"errors"
	"os"
)

type PgConfig struct {
	Host     string
	Username string
	Password string
}

func GetConfig() (PgConfig, error) {
	// turn all these into env vars maybe
	var pgPassword, err = os.LookupEnv("pgPassword")
	if !err {
		return PgConfig{}, errors.New("pgPassword is not in environment")
	}
	var c PgConfig = PgConfig{"0.0.0.0:111", "", pgPassword}
	return c, nil
}
