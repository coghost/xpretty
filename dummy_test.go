package xpretty_test

import (
	"testing"

	"github.com/coghost/xpretty"
	"github.com/stretchr/testify/suite"
)

type DummySuite struct {
	suite.Suite
}

func TestDummy(t *testing.T) {
	suite.Run(t, new(DummySuite))
}

func (s *DummySuite) SetupSuite() {
	xpretty.Initialize(xpretty.WithNoColor(false), xpretty.WithDummyLog(true))
}

func (s *DummySuite) TearDownSuite() {
}

func (s *DummySuite) TestDumbLog() {
	type args struct {
		msg []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{"empty", args{msg: []interface{}{}}},
		{"empty str", args{msg: []interface{}{""}}},
		{"str int", args{msg: []interface{}{"abc", 123}}},
	}

	for _, tt := range tests {
		xpretty.DummyLog(tt.args.msg...)
		s.True(true)
	}

	for _, tt := range tests {
		xpretty.DummyErrorLog(tt.args.msg...)
		s.True(true)
	}
}
