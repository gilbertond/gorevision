package contracts

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Contracts", func() {
	BeforeEach(func() {
		res:=setup()
		Expect(res).NotTo(BeNil())
		TestSetupProducer()
	})
   Context("pact broker setup", func() {
	   It("setup works", func() {
	   		
	   })
   })
})
