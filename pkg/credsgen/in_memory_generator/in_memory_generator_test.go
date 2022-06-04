package inmemorygenerator_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"code.cloudfoundry.org/quarks-secret/pkg/credsgen"
	inmemorygenerator "code.cloudfoundry.org/quarks-secret/pkg/credsgen/in_memory_generator"
	helper "code.cloudfoundry.org/quarks-utils/testing/testhelper"
	"os"
	"strconv"
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
			EXP,err := strconv.ParseUint(os.Getenv("CERT_EXPIRY_DAYS"), 10, 32)
			_ = err
			KEY_BITS,err2 := strconv.ParseUint(os.Getenv("CERT_KEY_BITS"), 10, 32)
			_ = err2
			ALG := os.Getenv("CERT_ALGORITHM")

			It("succeeds if the default generator is "+strconv.Itoa(KEY_BITS)+" bits", func() {
				Expect(defaultGenerator.(*inmemorygenerator.InMemoryGenerator).Bits).To(Equal(int(KEY_BITS)))
			})

			It("succeeds if the default generator is "+ALG, func() {
				Expect(defaultGenerator.(*inmemorygenerator.InMemoryGenerator).Algorithm).To(Equal(ALG))
			})

			It("succeeds if the default generator certs expires in "+strconv.Itoa(EXP)+" days", func() {
				Expect(defaultGenerator.(*inmemorygenerator.InMemoryGenerator).Expiry).To(Equal(int(EXP)))
			})
		})
	})
})
