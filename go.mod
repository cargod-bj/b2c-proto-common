module github.com/cargod-bj/b2c-proto-common

go 1.14

require (
	github.com/cargod-bj/b2c-common v0.0.0-20200602083522-21f4cf946177
	github.com/golang/protobuf v1.4.2
	github.com/mitchellh/mapstructure v1.1.2
	google.golang.org/protobuf v1.23.0
)

replace (
	github.com/cargod-bj/b2c-common => ../b2c-common
)