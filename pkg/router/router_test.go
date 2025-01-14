package router

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T) {
	tt := &treeNode{}
	num1 := []int{55}
	num2 := 66
	tt.registerRouter(num1, 0, 1, 2, 3)
	ret, accurate, err := tt.getVal(1)
	assert.Nil(t, ret)
	assert.Equal(t, false, accurate)
	assert.NotNil(t, err)

	tt.registerRouter(num2)
	ret, accurate, err = tt.getVal()
	assert.Equal(t, num2, ret)
	assert.Equal(t, true, accurate)
	assert.Nil(t, err)

	ret, accurate, err = tt.getVal(1)
	assert.Equal(t, num2, ret)
	assert.Equal(t, false, accurate)
	assert.Nil(t, err)

	ret, accurate, err = tt.getVal(0, 1, 2, 3)
	assert.Equal(t, num1, ret)
	assert.Equal(t, true, accurate)
	assert.Nil(t, err)

	ret, accurate, err = tt.getVal(0, 1, 2)
	assert.Equal(t, num2, ret)
	assert.Equal(t, false, accurate)
	assert.Nil(t, err)
}
