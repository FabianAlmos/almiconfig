package almierrors_test

import (
	almierrors "github.com/FabianAlmos/almiconfig/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	accessSecretFieldName               = "AccessSecret"
	fieldRequiredAccessSecretErr        = "AlmiConfigError: Field: 'AccessSecret', is required"
	fieldRequiredAccessSecretInvalidErr = "AlmiConfigError: This is an invalid error, please report it to the developer of this module, err: Field: '%!s(MISSING)', is required"
)

func TestAlmiErrorMsg_Build_Successful(t *testing.T) {
	err := almierrors.FieldRequiredErr.Build(accessSecretFieldName)
	assert.Equal(t, fieldRequiredAccessSecretErr, err.String())
}

func TestAlmiErrorMsg_Build_Fail(t *testing.T) {
	err := almierrors.FieldRequiredErr.Build()
	assert.Equal(t, fieldRequiredAccessSecretInvalidErr, err.String())
}
