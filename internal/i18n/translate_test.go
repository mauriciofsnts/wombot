package i18n_test

import (
	"testing"

	"code.db.cafe/wombot/internal/i18n"
	"github.com/stretchr/testify/assert"
)

//meu primeiro teste em go uWu
func TestReplace(t *testing.T) {
	text := "Salve irmãozin, ${0} da ${1}"
	replaced := i18n.Replace(text, "Wombot", 17)
	assert.Equal(t, "Salve irmãozin, Wombot da 17", replaced)
}
