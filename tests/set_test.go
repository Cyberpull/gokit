package tests

import (
	"strings"
	"testing"

	"github.com/Cyberpull/gokit"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SetTestSuite struct {
	suite.Suite
}

func (x *SetTestSuite) TestJoin() {
	value := gokit.Join("/", "foo", "bar")
	assert.Equal(x.T(), "foo/bar", value)
}

func (x *SetTestSuite) TestJoinFunc() {
	entries := []any{"/foo/", "/bar/"}

	value := gokit.JoinFunc("/", entries, func(v string) string {
		v = strings.TrimPrefix(v, "/")
		v = strings.TrimSuffix(v, "/")
		return v
	})

	assert.Equal(x.T(), "foo/bar", value)
}

// ===============================

func TestSet(t *testing.T) {
	suite.Run(t, new(SetTestSuite))
}
