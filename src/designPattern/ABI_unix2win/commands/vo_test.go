package commands

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestCommandVo1(t *testing.T) {
	c := NewCommandVO("")
	assert.Equal(t, c.CommandName(), "")
}

func TestCommandVo2(t *testing.T) {
	c := NewCommandVO("ls")
	assert.Equal(t, c.CommandName(), "ls")
}

func TestCommandVo3(t *testing.T) {
	c := NewCommandVO("ls -a haha")
	assert.Equal(t, c.CommandName(), "ls")
	assert.Equal(t, c.ParamList()[0], "a")
	assert.Equal(t, c.FormatData(), "haha")
}



