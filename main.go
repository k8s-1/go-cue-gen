package main

import (
	"fmt"

	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/cue/load"
)

func main() {
	ctx := cuecontext.New()
	insts := load.Instances([]string{"."}, nil)
	v := ctx.BuildInstance(insts[0])
	fmt.Printf("%v\n", v)
}
