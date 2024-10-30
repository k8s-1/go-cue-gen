package main

import (
	"fmt"
	"log"

	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/cue/load"
	"cuelang.org/go/encoding/yaml"
	"cuelang.org/go/cue/ast"
)

func main() {
	ctx := cuecontext.New()

	config := &load.Config{
		Tags: []string{"dev"},
		TagVars: map[string]load.TagVar{
			"key": {
				Func: func() (ast.Expr, error) {
          return ast.NewString("value"), nil
				},
				Description: "A tag variable for key with value 'value'",
			},
		},
	}

	filePath := "./example/example2.cue"

	insts := load.Instances([]string{filePath}, config)
	if len(insts) == 0 || insts[0].Err != nil {
		log.Fatalf("error loading CUE instances: %v", insts[0].Err)
	}

	v := ctx.BuildInstance(insts[0])
	fmt.Printf("%v\n", v)

	b, err := yaml.Encode(v)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", string(b))
}
