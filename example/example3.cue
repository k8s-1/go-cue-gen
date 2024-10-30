@if(dev)

package example

// -t env=value
environment: string @tag(env) | *"novalue"
