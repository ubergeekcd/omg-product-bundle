package main

import (
	pgemfire "github.com/enaml-ops/omg-product-bundle/products/p-gemfire/plugin"
	"github.com/enaml-ops/pluginlib/product"
)

// Version is the version of the p-gemfire plugin.
var Version string = "v0.0.0" // overridden at link time

func main() {
	product.Run(&pgemfire.Plugin{
		Version: Version,
	})
}
