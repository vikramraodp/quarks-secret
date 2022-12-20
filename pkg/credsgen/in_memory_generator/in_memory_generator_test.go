package inmemorygenerator_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	helper "code.cloudfoundry.org/quarks-utils/testing/testhelper"

	"code.cloudfoundry.org/quarks-secret/pkg/credsgen"
	inmemorygenerator "code.cloudfoundry.org/quarks-secret/pkg/credsgen/in_memory_generator"
)

var _ = Describe("InMemoryGenerator", func() {
	var (
		defaultGenerator credsgen.Generator
	)

	BeforeEach(func() {
		_, log := helper.NewTestLogger()
		defaultGenerator = inmemorygenerator.NewInMemoryGenerator(log)
	})

	Describe("NewInMemoryGenerator", func() {
		Context("object defaults", func() {
			It("succeeds if the default type is inmemorygenerator.InMemoryGenerator", func() {
				t, ok := defaultGenerator.(*inmemorygenerator.InMemoryGenerator)
				Expect(ok).To(BeTrue())
				Expect(t).To(Equal(defaultGenerator))
			})

			It("succeeds if the default generator is 4096 bits", func() {
				Expect(defaultGenerator.(*inmemorygenerator.InMemoryGenerator).Bits).To(Equal(4096))
			})

			It("succeeds if the default generator is rsa", func() {
				Expect(defaultGenerator.(*inmemorygenerator.InMemoryGenerator).Algorithm).To(Equal("rsa"))
			})

			It("succeeds if the default generator certs expires in 365 days", func() {
				Expect(defaultGenerator.(*inmemorygenerator.InMemoryGenerator).Expiry).To(Equal(1095))
			})
		})
	})
})
