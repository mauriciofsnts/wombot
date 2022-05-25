package i18n_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReplaceSuccess(t *testing.T) {
	fmt.Println("batata mi like")
}

func TestReplaceFail(t *testing.T) {
	require.Equal(t, 1, 2)
	assert.Equal(t, 1, 2)
}
