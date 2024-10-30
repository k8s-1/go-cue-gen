package main

import (
	"fmt"
	"log"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/cue/format"
	"cuelang.org/go/cue/load"
)

func main() {
	// Create a new CUE context
	ctx := cuecontext.New()

	// Load CUE files from the current directory
	cueInstances := load.Instances([]string{"."}, &load.Config{})
	if len(cueInstances) == 0 {
		log.Fatal("no CUE files found")
	}
	if cueInstances[0].Err != nil {
		log.Fatalf("error loading CUE files: %v", cueInstances[0].Err)
	}

	// Build the first instance (adjust if you have multiple CUE files)
	value := ctx.BuildInstance(cueInstances[0])
	if value.Err() != nil {
		log.Fatalf("error building CUE instance: %v", value.Err())
	}

	// Convert to YAML
	yamlData, err := formatYAML(value)
	if err != nil {
		log.Fatalf("error formatting YAML: %v", err)
	}

	// Print the YAML output
	fmt.Println(string(yamlData))
}

// Helper function to convert a CUE value to YAML
func formatYAML(value cue.Value) ([]byte, error) {
	// Marshal to YAML format
	yamlData, err := format.Encode(value, format.YAML)
	if err != nil {
		return nil, err
	}
	return yamlData, nil
}
