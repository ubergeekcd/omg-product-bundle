package prabbitmq_test

import (
	"io/ioutil"

	"github.com/enaml-ops/enaml"
	prabbitmq "github.com/enaml-ops/omg-product-bundle/products/p-rabbitmq/plugin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	logging "github.com/op/go-logging"
)

var _ = Describe("prabbitmq plugin", func() {

	BeforeSuite(func() {
		// suppress logging for tests
		logging.SetBackend(logging.NewLogBackend(ioutil.Discard, "", 0))
	})

	Context("when generating a manifest with incomplete input", func() {
		var (
			p  *prabbitmq.Plugin
			dm *enaml.DeploymentManifest
		)

		BeforeEach(func() {
			p = new(prabbitmq.Plugin)
			manifestBytes := p.GetProduct([]string{"foo"}, []byte{})
			dm = enaml.NewDeploymentManifest(manifestBytes)
		})

		XIt("should have set the deployment name", func() {
			// TODO
		})

		It("should have the correct releases", func() {
			hasRelease := func(name, version string) bool {
				for i := range dm.Releases {
					if dm.Releases[i].Name == name && dm.Releases[i].Version == version {
						return true
					}
				}
				return false
			}

			Ω(hasRelease(prabbitmq.CFRabbitMQReleaseName, prabbitmq.CFRabbitMQReleaseVersion)).Should(BeTrue())
			Ω(hasRelease(prabbitmq.ServiceMetricsReleaseName, prabbitmq.ServiceMetricsReleaseVersion)).Should(BeTrue())
			Ω(hasRelease(prabbitmq.LoggregatorReleaseName, prabbitmq.LoggregatorReleaseVersion)).Should(BeTrue())
			Ω(hasRelease(prabbitmq.RabbitMQMetricsReleaseName, prabbitmq.RabbitMQMetricsReleaseVersion)).Should(BeTrue())
		})

		It("should have set the update", func() {
			Ω(dm.Update.Canaries).Should(Equal(1))
			Ω(dm.Update.CanaryWatchTime).Should(Equal("30000-300000"))
			Ω(dm.Update.UpdateWatchTime).Should(Equal("30000-300000"))
			Ω(dm.Update.MaxInFlight).Should(Equal(1))
			Ω(dm.Update.Serial).Should(BeTrue())
		})

		XIt("should configure the stemcell", func() {
			Ω(dm.Stemcells).ShouldNot(BeEmpty())
		})
	})
})