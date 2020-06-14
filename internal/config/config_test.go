
package config

import (
        "os"
        "testing"

        "github.com/stretchr/testify/assert"
)

func TestLoadEnv(t *testing.T) {
        assert := assert.New(t)

        os.Setenv(EnvConfigPath, "env.test")
        expectedConfig := Config{
                SomeUserName:    "some-user-name",
                SomeUserPass:    "some-default-pass",
        }
        actualConfig, err := LoadEnv()
        assert.Nil(err, "Error in loading the config file")
        assert.Equal(expectedConfig, actualConfig, "Expected config is not equal to loaded config")
}

func TestLoadEnvFailure(t *testing.T) {
        assert := assert.New(t)
        os.Setenv(EnvConfigPath, "non.existing.env.test")
        _, err := LoadEnv()
        assert.NotNil(err, "Should not have read the configuration")
}
