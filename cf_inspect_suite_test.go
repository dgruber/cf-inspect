package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCfInspect(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CfInspect Suite")
}
