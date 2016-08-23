package cloudfoundry_test

import (
	"fmt"

	"github.com/enaml-ops/omg-product-bundle/products/cloudfoundry/enaml-gen/cloud_controller_worker"
	. "github.com/enaml-ops/omg-product-bundle/products/cloudfoundry/plugin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/yaml.v2"
)

var _ = Describe("Cloud Controller Worker Partition", func() {
	Context("When initialized with a complete set of arguments", func() {
		var cloudControllerWorker InstanceGrouper
		BeforeEach(func() {
			plugin := new(Plugin)
			c := plugin.GetContext([]string{
				"cloudfoundry",
				"--consul-ip", "1.0.0.1",
				"--consul-ip", "1.0.0.2",
				"--az", "az",
				"--stemcell-name", "stemcell",
				"--consul-encryption-key", "consulencryptionkey",
				"--consul-ca-cert", "consul-ca-cert",
				"--consul-agent-cert", "consul-agent-cert",
				"--consul-agent-key", "consul-agent-key",
				"--consul-server-cert", "consulservercert",
				"--consul-server-key", "consulserverkey",
				"--cc-worker-vm-type", "ccworkervmtype",
				"--network", "foundry",
				"--cc-staging-upload-user", "staginguser",
				"--cc-staging-upload-password", "stagingpassword",
				"--cc-bulk-api-user", "bulkapiuser",
				"--cc-bulk-api-password", "bulkapipassword",
				"--cc-db-encryption-key", "dbencryptionkey",
				"--cc-internal-api-user", "internalapiuser",
				"--cc-internal-api-password", "internalapipassword",
				"--system-domain", "sys.yourdomain.com",
				"--app-domain", "apps.yourdomain.com",
				"--allow-app-ssh-access",
				"--nfs-server-address", "10.0.0.19",
				"--nfs-share-path", "/var/vcap/nfs",
				"--metron-secret", "metronsecret",
				"--metron-zone", "metronzoneguid",
				"--mysql-proxy-ip", "10.0.0.3",
				"--db-ccdb-username", "ccdbuser",
				"--db-ccdb-password", "ccdbpass",
				"--nats-user", "natsuser",
				"--nats-pass", "natspass",
				"--nats-port", "4333",
				"--nats-machine-ip", "10.0.0.4",
			})

			cloudControllerWorker = NewCloudControllerWorkerPartition(c)
		})

		It("then should not be nil", func() {
			Ω(cloudControllerWorker).ShouldNot(BeNil())
		})

		It("should have valid values", func() {
			Ω(cloudControllerWorker.HasValidValues()).Should(BeTrue())
		})

		It("should have the name of the Network correctly set", func() {
			igf := cloudControllerWorker.ToInstanceGroup()

			networks := igf.Networks
			Ω(len(networks)).Should(Equal(1))
			Ω(networks[0].Name).Should(Equal("foundry"))
		})

		It("should have NFS Mounter set as a job", func() {
			igf := cloudControllerWorker.ToInstanceGroup()
			nfsMounter := igf.Jobs[2]
			Ω(nfsMounter.Name).Should(Equal("nfs_mounter"))
		})

		XIt("should have NFS Mounter details set properly", func() {
			igf := cloudControllerWorker.ToInstanceGroup()

			b, _ := yaml.Marshal(igf)
			fmt.Print(string(b))
		})

		It("then it should have configured the cloud controller worker job", func() {
			igf := cloudControllerWorker.ToInstanceGroup()
			ccw := igf.GetJobByName("cloud_controller_worker")
			Ω(ccw.Release).Should(Equal(CFReleaseName))

			props := ccw.Properties.(*cloud_controller_worker.CloudControllerWorkerJob)
			Ω(props.Domain).Should(Equal("sys.yourdomain.com"))
			Ω(props.SystemDomain).Should(Equal("sys.yourdomain.com"))
			Ω(props.AppDomains).Should(ConsistOf("apps.yourdomain.com"))
			Ω(props.SystemDomainOrganization).Should(Equal("system"))

			ccdb := props.Ccdb
			Ω(ccdb.DbScheme).Should(Equal("mysql"))
			Ω(ccdb.Port).Should(Equal(3306))
			Ω(ccdb.Address).Should(Equal("10.0.0.3"))

			roles := ccdb.Roles.([]map[string]interface{})
			Ω(roles).Should(HaveLen(1))
			Ω(roles[0]).Should(HaveKeyWithValue("name", "ccdbuser"))
			Ω(roles[0]).Should(HaveKeyWithValue("password", "ccdbpass"))
			Ω(roles[0]).Should(HaveKeyWithValue("tag", "admin"))

			dbs := ccdb.Databases.([]map[string]interface{})
			Ω(dbs).Should(HaveLen(1))
			Ω(dbs[0]).Should(HaveKeyWithValue("citext", true))
			Ω(dbs[0]).Should(HaveKeyWithValue("name", "ccdb"))
			Ω(dbs[0]).Should(HaveKeyWithValue("tag", "cc"))

			cc := props.Cc
			Ω(cc.DefaultRunningSecurityGroups).Should(ConsistOf("all_open"))
			Ω(cc.DefaultStagingSecurityGroups).Should(ConsistOf("all_open"))

			quotas := cc.QuotaDefinitions.(map[string]interface{})
			Ω(quotas).Should(HaveLen(2))
			Ω(quotas).Should(HaveKey("default"))
			Ω(quotas).Should(HaveKey("runaway"))

			def := quotas["default"]
			Ω(def).Should(HaveKeyWithValue("memory_limit", 10240))
			Ω(def).Should(HaveKeyWithValue("total_services", 100))
			Ω(def).Should(HaveKeyWithValue("non_basic_services_allowed", true))
			Ω(def).Should(HaveKeyWithValue("total_routes", 1000))
			Ω(def).Should(HaveKeyWithValue("trial_db_allowed", true))

			runaway := quotas["runaway"]
			Ω(runaway).Should(HaveKeyWithValue("memory_limit", 102400))
			Ω(runaway).Should(HaveKeyWithValue("total_services", -1))
			Ω(runaway).Should(HaveKeyWithValue("non_basic_services_allowed", true))
			Ω(runaway).Should(HaveKeyWithValue("total_routes", 1000))

			nats := props.Nats
			Ω(nats.User).Should(Equal("natsuser"))
			Ω(nats.Port).Should(Equal(4333))
			Ω(nats.Password).Should(Equal("natspass"))
			Ω(nats.Machines).Should(ConsistOf("10.0.0.4"))
		})
	})
})
