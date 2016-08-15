package main

import (
	"github.com/enaml-ops/omg-product-bundle/products/concourse/plugin"
	"github.com/enaml-ops/pluginlib/product"
)

var Version string = "v0.0.0"

func main() {
	product.Run(&concourseplugin.ConcoursePlugin{
		PluginVersion: Version,
	})
}
