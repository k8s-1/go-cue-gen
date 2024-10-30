package main

import (
	"fmt"
	"log"

	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/cue/load"

	"cuelang.org/go/encoding/yaml"
)

func main() {
	ctx := cuecontext.New()
	insts := load.Instances([]string{"."}, nil)
	v := ctx.BuildInstance(insts[0])
	fmt.Printf("%v\n", v)

	b, err := yaml.Encode(v)
	if err != nil {
		log.Fatal(err)
	}

  fmt.Printf("%s\n", string(b))
}
