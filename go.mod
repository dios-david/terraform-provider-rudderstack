module terraform-provider-rudderstack

go 1.16

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/hashicorp/terraform-plugin-framework v0.5.0
	github.com/rudderlabs/cp-client-go v0.0.13
	golang.org/x/net v0.0.0-20210929193557-e81a3d93ecf6 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
)

// replace github.com/rudderlabs/cp-client-go v0.0.13 => ../cp-client-go

replace github.com/hashicorp/terraform-plugin-framework v0.5.0 => ../terraform-plugin-framework
