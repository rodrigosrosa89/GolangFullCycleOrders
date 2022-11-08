package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenAndEMptyID_WhenCreateANewOrder_ThemShouldReceiveAnError(t *testing.T) {
	order := Order{}
	assert.Error(t, order.IsValid(), "invalid id")

}
