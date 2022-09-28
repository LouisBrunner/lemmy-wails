package config

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/LouisBrunner/lemmy/library/backend/internal/common"
	"github.com/LouisBrunner/lemmy/library/backend/internal/contracts"
	"github.com/sirupsen/logrus"
)

const configFilename = "config.json"

type config[User any] struct {
	Internal contracts.ConfigInternal
	User     User
}

type configManager[User any] struct {
	logger    *logrus.Logger
	folder    string
	config    config[User]
	hasLoaded bool
}

func NewManager[User any](logger *logrus.Logger, folder string) contracts.ConfigManager[User] {
	me := &configManager[User]{
		logger: logger,
		folder: folder,
	}
	readConfig, err := me.readConfig()
	if err != nil {
		me.logger.Warningf("failed to fetch config, replace with defaults: %v", err)
		readConfig = &config[User]{}
	} else {
		me.logger.Infof("config read successfully")
		me.logger.Debugf("config: %+v", readConfig.Internal)
		me.hasLoaded = true
	}
	me.config = *readConfig
	return me
}

func (me *configManager[User]) getConfigPath() (string, error) {
	path := filepath.Join(me.folder, configFilename)
	return path, common.MkdirIfDoNotExist(me.folder)
}

func (me *configManager[User]) readConfig() (*config[User], error) {
	path, err := me.getConfigPath()
	if err != nil {
		return nil, err
	}
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	parsed := &config[User]{}
	err = json.Unmarshal(content, parsed)
	if err != nil {
		return nil, err
	}
	return parsed, nil
}

// User
func (me *configManager[User]) Load() (User, bool) {
	return me.config.User, me.hasLoaded
}

func (me *configManager[User]) Save(newConfig User) error {
	return me.save(config[User]{
		Internal: me.config.Internal,
		User:     newConfig,
	})
}

// Internal
func (me *configManager[User]) Internal() contracts.ConfigInternal {
	return me.config.Internal
}

func (me *configManager[User]) SaveInternal(newConfig contracts.ConfigInternal) error {
	return me.save(config[User]{
		Internal: newConfig,
		User:     me.config.User,
	})
}

// Helpers
func (me *configManager[User]) save(newConfig config[User]) error {
	content, err := json.Marshal(newConfig)
	if err != nil {
		return err
	}
	path, err := me.getConfigPath()
	if err != nil {
		return err
	}
	err = os.WriteFile(path, content, 0644)
	if err != nil {
		return err
	}
	me.config = newConfig
	me.hasLoaded = true
	return nil
}
