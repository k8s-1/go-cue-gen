import (
	"github.com/k8s-1/go-cue-gen/example"
)

//environment: string @tag(env)
someval: {
	x: string & "hello",
	y: string & "bye",
}
