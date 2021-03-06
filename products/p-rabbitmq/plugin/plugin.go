package prabbitmq

import (
	"fmt"
	"strings"

	cli "gopkg.in/urfave/cli.v2"

	"github.com/enaml-ops/enaml"
	"github.com/enaml-ops/pluginlib/pcli"
	"github.com/enaml-ops/pluginlib/pluginutil"
	"github.com/enaml-ops/pluginlib/product"
	"github.com/xchapter7x/lo"
)

// Plugin is an omg product plugin for deploying p-rabbitmq.
type Plugin struct {
	Version string
}

// generatePassword is the default for password flags that should be generated by
// the plugin if not specified by the user
const generatePassword = "[autogenerated]"

// GetFlags returns the CLI flags accepted by the plugin.
func (p *Plugin) GetFlags() []pcli.Flag {
	return []pcli.Flag{
		pcli.CreateStringFlag("deployment-name", "the name bosh will use for the deployment", "p-rabbitmq"),
		pcli.CreateStringFlag("service-admin-password", "the password used by cloud controller for authentication", generatePassword),
		pcli.CreateStringFlag("system-domain", "the system domain"),
		pcli.CreateStringSliceFlag("az", "list of AZ names to use"),
		pcli.CreateStringFlag("rabbit-public-ip", "the public IP"),
		pcli.CreateStringFlag("rabbit-admin-password", "the admin password to use", generatePassword),
		pcli.CreateStringFlag("network", "the name of the network to use"),
		pcli.CreateStringFlag("stemcell-ver", "the version number of the stemcell you wish to use", StemcellVersion),
		pcli.CreateStringSliceFlag("rabbit-server-ip", "rabbit-mq server IPs to use"),
		pcli.CreateStringFlag("rabbit-broker-ip", "IP of the rabbitmq broker"),
		pcli.CreateStringFlag("broker-password", "password for the rabbitmq broker", generatePassword),
		pcli.CreateStringFlag("syslog-address", "the address of your syslog drain"),
		pcli.CreateIntFlag("syslog-port", "the port for your syslog connection", "514"),
		pcli.CreateStringSliceFlag("nats-machine-ip", "IP addresses of NATS machines"),
		pcli.CreateIntFlag("nats-port", "NATS port", "4222"),
		pcli.CreateStringFlag("nats-pass", "password for NATS", generatePassword),
		pcli.CreateStringFlag("haproxy-stats-password", "admin password to acces HAproxy stats dashboard", generatePassword),
		pcli.CreateStringFlag("system-services-password", "password for CF system_services account"),
		pcli.CreateBoolFlag("skip-ssl-verify", "skip SSL verification"),
		pcli.CreateStringFlag("doppler-zone", "the name zone for doppler"),
		pcli.CreateStringFlag("doppler-shared-secret", "doppler shared secret"),
		pcli.CreateStringSliceFlag("etcd-machine-ip", "IPs of etcd machines"),
		pcli.CreateBoolFlag("infer-from-cloud", "attempt to pull defaults from your targetted bosh"),
		pcli.CreateStringFlag("rabbit-broker-vm-type", "VM type for broker"),
		pcli.CreateStringFlag("rabbit-haproxy-vm-type", "VM type for ha proxy"),
		pcli.CreateStringFlag("rabbit-server-vm-type", "VM type for RabbitMQ server"),

		pcli.CreateStringFlag("vault-domain", "the location of your vault server (ie. http://10.0.0.1:8200)"),
		pcli.CreateStringFlag("vault-token", "the token to make connections to your vault"),
		pcli.CreateStringSliceFlag("vault-hash", "a list of vault hashes to pull values from"),
	}
}

// GetMeta returns metadata about the p-rabbitmq product.
func (p *Plugin) GetMeta() product.Meta {
	return product.Meta{
		Name: "p-rabbitmq",
		Stemcell: enaml.Stemcell{
			Name:    StemcellName,
			Alias:   StemcellAlias,
			Version: StemcellVersion,
		},
		Releases: []enaml.Release{
			enaml.Release{
				Name:    CFRabbitMQReleaseName,
				Version: CFRabbitMQReleaseVersion,
			},
			enaml.Release{
				Name:    ServiceMetricsReleaseName,
				Version: ServiceMetricsReleaseVersion,
			},
			enaml.Release{
				Name:    LoggregatorReleaseName,
				Version: LoggregatorReleaseVersion,
			},
			enaml.Release{
				Name:    RabbitMQMetricsReleaseName,
				Version: RabbitMQMetricsReleaseVersion,
			},
		},
		Properties: map[string]interface{}{
			"version":                  p.Version,
			"stemcell":                 StemcellVersion,
			"pivotal-rabbit-mq":        fmt.Sprintf("%s / %s", "pivotal-rabbit-mq", ProductVersion), // TODO match pivnet on name
			"cf-rabbitmq-release":      fmt.Sprintf("%s / %s", CFRabbitMQReleaseName, CFRabbitMQReleaseVersion),
			"service-metrics-release":  fmt.Sprintf("%s / %s", ServiceMetricsReleaseName, ServiceMetricsReleaseVersion),
			"loggregator-release":      fmt.Sprintf("%s / %s", LoggregatorReleaseName, LoggregatorReleaseVersion),
			"rabbitmq-metrics-release": fmt.Sprintf("%s / %s", RabbitMQMetricsReleaseName, RabbitMQMetricsReleaseVersion),
		},
	}
}

// GetProduct generates a BOSH deployment manifest for p-rabbitmq.
func (p *Plugin) GetProduct(args []string, cloudConfig []byte) ([]byte, error) {
	var err error
	flags := p.GetFlags()
	c := pluginutil.NewContext(args, pluginutil.ToCliFlagArray(flags))

	if c.Bool("infer-from-cloud") {
		inferFromCloud(cloudConfig, flags, c)
		c = pluginutil.NewContext(args, pluginutil.ToCliFlagArray(flags))
	}

	// populate flags from vault if configured to do so
	domain := c.String("vault-domain")
	tok := c.String("vault-token")
	hashes := c.StringSlice("vault-hash")
	if domain != "" && tok != "" && len(hashes) > 0 {
		lo.G.Debug("connecting to vault at", domain)
		v := pluginutil.NewVaultUnmarshal(domain, tok)
		for _, hash := range hashes {
			v.UnmarshalFlags(hash, flags)
		}
		c = pluginutil.NewContext(args, pluginutil.ToCliFlagArray(flags))
	}

	var cfg *Config
	cfg, err = configFromContext(c)

	if err != nil {
		lo.G.Error(err.Error())
		return nil, err
	}
	dm := new(enaml.DeploymentManifest)
	dm.SetName(cfg.DeploymentName)

	dm.AddRelease(enaml.Release{Name: CFRabbitMQReleaseName, Version: CFRabbitMQReleaseVersion})
	dm.AddRelease(enaml.Release{Name: ServiceMetricsReleaseName, Version: ServiceMetricsReleaseVersion})
	dm.AddRelease(enaml.Release{Name: LoggregatorReleaseName, Version: LoggregatorReleaseVersion})
	dm.AddRelease(enaml.Release{Name: RabbitMQMetricsReleaseName, Version: RabbitMQMetricsReleaseVersion})

	dm.AddStemcell(enaml.Stemcell{OS: StemcellName, Version: cfg.StemcellVersion, Alias: StemcellAlias})

	dm.AddInstanceGroup(p.NewRabbitMQServerPartition(cfg))
	dm.AddInstanceGroup(p.NewRabbitMQBrokerPartition(cfg))
	dm.AddInstanceGroup(p.NewRabbitMQHAProxyPartition(cfg))
	dm.AddInstanceGroup(p.NewRabbitMQBrokerRegistrar(cfg))
	dm.AddInstanceGroup(p.NewRabbitMQBrokerDeregistrar(cfg))

	dm.Update = enaml.Update{
		Canaries:        1,
		CanaryWatchTime: "30000-300000",
		UpdateWatchTime: "30000-300000",
		MaxInFlight:     1,
		Serial:          true,
	}

	return dm.Bytes(), err
}

func inferFromCloud(cloudConfig []byte, flags []pcli.Flag, c *cli.Context) {
	inferer := pluginutil.NewCloudConfigInferFromBytes(cloudConfig)

	vm := inferer.InferDefaultVMType()
	network := inferer.InferDefaultNetwork()
	az := inferer.InferDefaultAZ()

	for i := range flags {
		if !c.IsSet(flags[i].Name) {
			if flags[i].Name == "network" {
				lo.G.Debugf("got network '%s' from cloud config", network)
				flags[i].Value = network
			} else if flags[i].Name == "az" {
				lo.G.Debugf("got azs '%v' from cloud config", az)
				flags[i].Value = az
			} else if strings.HasSuffix(flags[i].Name, "vm-type") {
				lo.G.Debugf("got flag %s from cloud config (value=%s)", flags[i].Name, vm)
				flags[i].Value = vm
			}
		}
	}
}
