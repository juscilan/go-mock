package main

import (
	"testing"

	"github.com/juscilan/godeep/deep"
	"github.com/stretchr/testify/mock"
)

func TestMainExec(t *testing.T) {
	t.Run("should return a string", func(t *testing.T) {
		d := new(deep.Deep)
		sut := CallDeep(d)
		if sut == "" {
			t.Errorf("Erro")
		}
	})
}

type DeepMock struct {
	mock.Mock
	spyOn string
}

func (d *DeepMock) Exec() (string, error) {
	args := d.Called()
	return d.spyOn, args.Error(0)
}

func TestDependencyWithMock(t *testing.T) {

	t.Run("should return a string like a mock", func(t *testing.T) {
		md := new(DeepMock)
		md.spyOn = "Testing...123"
		md.On("Exec", mock.Anything).Return(nil)

		rm := CallDeep(md)
		re := "Testing...123"

		if rm != re {
			t.Errorf("Erro, expected %s but got %s", re, rm)
		}
	})

}
