package tests

import (
	"testing"

	"github.com/Cyberpull/gokit"

	"github.com/stretchr/testify/assert"
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

// ===============================

func TestPath(t *testing.T) {
	suite.Run(t, new(PathTestSuite))
}
