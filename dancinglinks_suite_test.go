package dancinglinks_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDancinglinks(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Dancinglinks Suite")
}
