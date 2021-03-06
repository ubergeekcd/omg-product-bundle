package prabbitmq_test

import (
	"github.com/enaml-ops/enaml"
	bd "github.com/enaml-ops/omg-product-bundle/products/p-rabbitmq/enaml-gen/broker-deregistrar"
	prabbitmq "github.com/enaml-ops/omg-product-bundle/products/p-rabbitmq/plugin"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("rabbitmq broker registrar", func() {
	Context("when initialized with a complete configuration", func() {
		var ig *enaml.InstanceGroup

		const (
			controlNetworkName            = "foundry-net"
			controlSystemServicesPassword = "systemservicespassword"
		)

		BeforeEach(func() {
			p := new(prabbitmq.Plugin)
			c := &prabbitmq.Config{
				Network:                controlNetworkName,
				SystemServicesPassword: controlSystemServicesPassword,
				SystemDomain:           "sys.example.com",
				SkipSSLVerify:          true,
			}
			ig = p.NewRabbitMQBrokerDeregistrar(c)
			Ω(ig).ShouldNot(BeNil())
		})

		It("should configure the instance group parameters", func() {
			Ω(ig.Name).Should(Equal("broker-deregistrar"))
			Ω(ig.Lifecycle).Should(Equal("errand"))
			Ω(ig.Instances).Should(Equal(1))
			Ω(ig.Networks).Should(HaveLen(1))
			Ω(ig.Networks[0].Name).Should(Equal(controlNetworkName))
			Ω(ig.Networks[0].Default).Should(ConsistOf("dns", "gateway"))
		})

		It("should configure the broker-deregistrar job", func() {
			Ω(ig.Jobs).Should(HaveLen(1))
			Ω(ig.Jobs[0].Properties).ShouldNot(BeNil())
			Ω(ig.Jobs[0].Name).Should(Equal("broker-deregistrar"))
			Ω(ig.Jobs[0].Release).Should(Equal(prabbitmq.CFRabbitMQReleaseName))

			props := ig.Jobs[0].Properties.(*bd.BrokerDeregistrarJob)
			Ω(props).ShouldNot(BeNil())
			Ω(props.Broker).ShouldNot(BeNil())
			Ω(props.Broker.Name).Should(Equal("p-rabbitmq"))

			Ω(props.Cf).ShouldNot(BeNil())
			Ω(props.Cf.ApiUrl).Should(Equal("https://api.sys.example.com"))
			Ω(props.Cf.AdminUsername).Should(Equal("system_services"))
			Ω(props.Cf.AdminPassword).Should(Equal(controlSystemServicesPassword))
			Ω(props.Cf.SkipSslValidation).Should(BeTrue())
		})
	})
})
