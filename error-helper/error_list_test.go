package error_helper

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestErrorList(t *testing.T) {
	errList := NewErrorList()
	errList.Add(errors.New("error 1"), errors.New("error 2"), errors.New("error 3"))

	assert.Equal(t, true, errList.HasErrors())
	assert.Equal(t, errList, errList.ErrOrNil())
	assert.NotEmpty(t, errList.Error())
}
