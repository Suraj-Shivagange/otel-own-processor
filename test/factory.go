package splitbatchprocessor

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
)

const (
	// The value of "type" key in configuration.
	typeStr = "splitbatch"
)

// NewFactory creates a factory for the routing processor.
func NewFactory() component.ProcessorFactory {
	// TODO: find a more appropriate way to get this done, as we are swallowing the error here

	return component.NewProcessorFactory(
		typeStr,
		createDefaultConfig,
	)
}

func createDefaultConfig() config.Processor {
	return &Config{
		ProcessorSettings: config.NewProcessorSettings(config.NewComponentID(typeStr)),
	}
}
