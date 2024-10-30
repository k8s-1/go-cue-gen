package main

import (
	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/cue/load"
	"cuelang.org/go/encoding/yaml"
	"fmt"
	"log"
)

func main() {
	ctx := cuecontext.New()

	config := &load.Config{
		// Use the 'dev' tag.
		Tags: []string{"dev", "env=value"},
	}

	filePath := "./main.cue"

	// Load the instance.
	insts := load.Instances([]string{filePath}, config)
	if len(insts) == 0 || insts[0].Err != nil {
		log.Fatalf("error loading CUE instances: %v", insts[0].Err)
	}

	v := ctx.BuildInstance(insts[0])
	if v.Err() != nil {
		log.Fatalf("error building instance: %v", v.Err())
	}

	// Print the instance value.
	fmt.Printf("CUE Value: %v\n", v)

	// Encode the instance to YAML.
	b, err := yaml.Encode(v)
	if err != nil {
		log.Fatalf("error encoding to YAML: %v", err)
	}

	// Print the YAML output.
	fmt.Printf("YAML Output:\n%s\n", string(b))
}
