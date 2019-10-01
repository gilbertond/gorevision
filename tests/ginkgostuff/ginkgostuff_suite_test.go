package ginkgostuff_test

import (
	"fmt"
	. "github.com/gorevision/channels"
	. "github.com/gorevision/tests/ginkgostuff"
	"github.com/stretchr/testify/assert"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGinkgoStuff(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ginkgo stuff Suite")
}

var _ = Describe("my test suite", func() {

	BeforeEach(func() {
		fmt.Println("Before each test runs")
	})

	AfterEach(func() {
		fmt.Println("After each test runs")
	})

	Context("context1", func() {
		fmt.Println("CONTEXT 1")
		BeforeEach(func() {
			fmt.Println("before context1 tests")
		})

		AfterEach(func() {
			fmt.Println("after context1 tests")
		})

		It("context1 test1", func() {
			Expect(true)
			fmt.Println("context1 test 1")
		})
	})

	Context("context2", func() {
		fmt.Println("CONTEXT 2")
		BeforeEach(func() {
			fmt.Println("before context2 tests run")
		})
		AfterEach(func() {
			fmt.Println("after context2 tests run")
		})
		It("context2 test 1", func() {
			fmt.Println("context2 test 1")
			Expect(true)
		}, 2000)

		It("context2 test 2", func() {
			fmt.Println("context2 test 2")
			Expect(true)
		}, 2000)
	})

	It("sum test", func() {
		fmt.Println("testing sum")

		//Different assertions
		Expect(Sum(1, 2)).To(BeEquivalentTo(3))
		Expect(Sum(2, 2)).Should(Equal(4))

		sum, error := Sum(1, 1)
		_, customError := SumWithError(1, 1)

		Expect(error).Should(Succeed()) //Assert no error occurred
		Expect(customError).ShouldNot(Succeed())//😡😡😡😡😡😡😡😡😡😡😡😡😡😡😡😡😡

		assert.True(nil, sum == 2)

	})

	It("assertions on channels", func() {
		channel := make(chan int, 5)
		go SomeFunc(cap(channel), channel)

		//close(channel)
		Eventually(channel).Should(Receive(Equal(1)))
	})
})
