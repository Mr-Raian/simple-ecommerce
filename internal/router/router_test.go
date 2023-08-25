package router_test

import (
	"api/internal/handler"
	"api/internal/router"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	r := router.New(handler.Handler{})
	require.NotNil(t, r)
	assert.NotZero(t, r)
}
