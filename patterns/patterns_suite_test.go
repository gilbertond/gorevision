package test

import (
	"testing"
	
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"
	
	"github.com/gorevision/patterns/singleton"
)

func TestPatterns(t *testing.T)  {
	    RegisterFailHandler(Fail)
	    RunSpecs(t, "design patterns test suite")
}

var _ = Describe("", func() {
	Context("singleton", func() {
		It("can create a singleton", func() {
		    singleton:=singleton.NewSingleton()
		    assert.NotNil(singleton)
		    Expect(singleton).NotTo(BeNil())
		})
	})
	
	Context("proxy", func() {
		It("", func() {
		
		})
	})
})