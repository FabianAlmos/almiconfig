package almiconfig

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

const (
	accessSecretEnv  = "ACCESS_SECRET"
	refreshSecretEnv = "REFRESH_SECRET"
	kafkaBrokersEnv  = "KAFKA_BROKERS"

	accessSecret  = "access_secret"
	refreshSecret = "refresh_secret"
	kafkaBrokers  = "broker1,broker2,broker3"
)

type testConfig struct {
	AccessSecret  string   `almi:"required,env=ACCESS_SECRET"`
	RefreshSecret string   `almi:"env=REFRESH_SECRET"`
	KafkaBrokers  []string `almi:"required,env=KAFKA_BROKERS,type=[,]string"`
}

type testConfigNoEnvForAccessSecret struct {
	AccessSecret  string   `almi:"required"`
	RefreshSecret string   `almi:"env=REFRESH_SECRET"`
	KafkaBrokers  []string `almi:"required,env=KAFKA_BROKERS,type=[,]string"`
}

type testConfigTypeConstraint struct {
	AccessSecret string `almi:"required,env=ACCESS_SECRET,type=string"`
}

type testConfigInvalidConstraint struct {
	AccessSecret string `almi:"required,env=ACCESS_SECRET,invalid_struct_tag"`
}

type testConfigBadTypeConversion struct {
	AccessLifetime int `almi:"required,env=KAFKA_BROKERS,type=int"`
}

type testConfigTypeMismatch struct {
	AccessLifetime int `almi:"required,env=KAFKA_BROKERS,type=string"`
}

func TestValidateConfig_Successful(t *testing.T) {
	os.Clearenv()

	if err := os.Setenv(accessSecretEnv, accessSecret); err != nil {
		t.Fail()
	}

	if err := os.Setenv(refreshSecretEnv, refreshSecret); err != nil {
		t.Fail()
	}

	if err := os.Setenv(kafkaBrokersEnv, kafkaBrokers); err != nil {
		t.Fail()
	}

	cfg, err := ValidateConfig(testConfig{})
	if err != nil {
		t.Fail()
	}

	assert.NotNil(t, cfg)
}

func TestValidateConfig_Fail_NoEnvConstraint(t *testing.T) {
	os.Clearenv()

	if err := os.Setenv(accessSecretEnv, accessSecret); err != nil {
		t.Fail()
	}

	if err := os.Setenv(refreshSecretEnv, refreshSecret); err != nil {
		t.Fail()
	}

	if err := os.Setenv(kafkaBrokersEnv, kafkaBrokers); err != nil {
		t.Fail()
	}

	cfg, err := ValidateConfig(testConfigNoEnvForAccessSecret{})
	assert.Nil(t, cfg)
	assert.NotNil(t, err)
}

func TestValidateConfig_Fail_RequiredAndNotSet(t *testing.T) {
	os.Clearenv()

	if err := os.Setenv(refreshSecretEnv, refreshSecret); err != nil {
		t.Fail()
	}

	if err := os.Setenv(kafkaBrokersEnv, kafkaBrokers); err != nil {
		t.Fail()
	}

	cfg, err := ValidateConfig(testConfig{})
	assert.Nil(t, cfg)
	assert.NotNil(t, err)
}

func TestValidateConfig_CheckTypeConstraint(t *testing.T) {
	os.Clearenv()

	if err := os.Setenv(accessSecretEnv, accessSecret); err != nil {
		t.Fail()
	}

	cfg, err := ValidateConfig(testConfigTypeConstraint{})
	if err != nil {
		t.Fail()
	}

	assert.NotNil(t, cfg)
}

func TestValidateConfig_Fail_InvalidStructTag(t *testing.T) {
	os.Clearenv()

	if err := os.Setenv(accessSecretEnv, accessSecret); err != nil {
		t.Fail()
	}

	cfg, err := ValidateConfig(testConfigInvalidConstraint{})
	assert.Nil(t, cfg)
	assert.NotNil(t, err)
}

func TestValidateConfig_Fail_BadTypeConversion(t *testing.T) {
	os.Clearenv()

	if err := os.Setenv(kafkaBrokers, kafkaBrokers); err != nil {
		t.Fail()
	}

	cfg, err := ValidateConfig(testConfigBadTypeConversion{})
	assert.Nil(t, cfg)
	assert.NotNil(t, err)
}

func TestValidateConfig_Fail_TypeMismatch(t *testing.T) {
	os.Clearenv()

	if err := os.Setenv(kafkaBrokers, kafkaBrokers); err != nil {
		t.Fail()
	}

	cfg, err := ValidateConfig(testConfigTypeMismatch{})
	assert.Nil(t, cfg)
	assert.NotNil(t, err)
}
