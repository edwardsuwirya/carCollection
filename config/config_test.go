package config

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

type ConfigTestSuite struct {
	suite.Suite
	configFile string
}

func (suite *ConfigTestSuite) SetupSuite() {
	suite.configFile = "config-test.json"
	d1 := []byte(`{"log_level":"debug"}`)
	err := ioutil.WriteFile(suite.configFile, d1, 0644)
	if err != nil {
		panic(err)
	}
}

func (suite *ConfigTestSuite) TearDownSuite() {
	err := os.Remove(suite.configFile)
	if err != nil {
		panic(err)
	}
}

func (suite *ConfigTestSuite) TestBuildConfig() {
	c := NewConfig(suite.configFile)
	assert.Equal(suite.T(), suite.configFile, c.configFilePath)
}

func (suite *ConfigTestSuite) TestConfig() {
	c := NewConfig(suite.configFile)
	vp := c.init()
	assert.Equal(suite.T(), vp.ConfigFileUsed(), suite.configFile)
}

func (suite *ConfigTestSuite) TestConfigNotFound() {
	assert.Panics(suite.T(), func() {
		c := NewConfig("dummy.json")
		c.init()
	})
}

func (suite *ConfigTestSuite) TestInit() {
	c := NewConfig(suite.configFile)
	c.init()
	assert.NotPanics(suite.T(), func() {
		log.Print("Pass")
	})
}
func (suite *ConfigTestSuite) TestGetConfigValue() {
	c := NewConfig(suite.configFile)
	c.init()
	assert.Equal(suite.T(), "debug", c.GetConfigValue("log_level"))
}
func (suite *ConfigTestSuite) TestLogLevelNotFound() {
	assert.Panics(suite.T(), func() {
		c := NewConfig("dummy.json")
		c.init()
	})
}
func TestConfigTestSuite(t *testing.T) {
	suite.Run(t, new(ConfigTestSuite))
}
