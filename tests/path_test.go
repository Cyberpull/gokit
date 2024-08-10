package tests

import (
	"strings"
	"testing"

	"github.com/Cyberpull/gokit"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type PathTestSuite struct {
	suite.Suite
}

func (x *PathTestSuite) TestJoin() {
	value := gokit.Path.Join("/", "foo", "bar")
	assert.Equal(x.T(), "foo/bar", value)
}

func (x *PathTestSuite) TestJoinWithSlashPrefix() {
	value := gokit.Path.Join("/", "/foo", "/bar")
	assert.Equal(x.T(), "foo/bar", value)
}

func (x *PathTestSuite) TestJoinWithSlashSuffix() {
	value := gokit.Path.Join("/", "foo/", "bar/")
	assert.Equal(x.T(), "foo/bar", value)
}

func (x *PathTestSuite) TestExpand() {
	value := gokit.Path.Expand("~/www/gmail")
	assert.True(x.T(), strings.HasPrefix(value, "/home"))
}

func (x *PathTestSuite) TestExcecutablePath() {
	path, err := gokit.Path.FromExecutable()
	require.NoError(x.T(), err)
	assert.NotEqual(x.T(), "", path)
}

// ===============================

func TestPath(t *testing.T) {
	suite.Run(t, new(PathTestSuite))
}
