package sts

import (
	"errors"

	"oiam.io/pkg/access"
	"oiam.io/pkg/storage"
)

func New(config *Config, storageInterface storage.Interface, authorizer access.Interface) (Interface, error) {
	if config.Redis != nil {
		return NewRedis(config.Redis.Addr, config.Redis.Password, config.Redis.DB, storageInterface, authorizer)
	}
	return nil, errors.New("no storage provider is configured")
}
