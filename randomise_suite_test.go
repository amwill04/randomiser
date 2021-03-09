package randomise_test

import (
	"testing"

	"github.com/onsi/gomega/format"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRandomise(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Randomise Suite")
}

var _ = BeforeSuite(func() {
	format.UseStringerRepresentation = true
})
