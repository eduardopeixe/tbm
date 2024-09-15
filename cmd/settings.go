package main

import "errors"

type Settings struct {
	Port int `json:"port"`
	DB   DB  `json:"db"`
}

func (s *Settings) valid() error {
	err := s.DB.valid()
	if err != nil {
		return err
	}

	return nil

}

type DB struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	DBName   string `json:"dbName"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (d *DB) valid() error {
	if d.Host == "" {
		return errors.New("error 'Host' not present")
	}
	if d.Port == "" {
		return errors.New("error 'Port' not present")
	}
	if d.DBName == "" {
		return errors.New("error 'DBName' not present")
	}
	if d.Username == "" {
		return errors.New("error 'Username' not present")
	}
	if d.Password == "" {
		return errors.New("error 'Password' not present")
	}

	return nil
}

// helpers

func loadSettings() (Settings, error) {
	return Settings{
		Port: 8081,
	}, nil
}
