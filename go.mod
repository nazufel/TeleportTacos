module github.com/teleporttacos

go 1.19

require (
	github.com/gocql/gocql v1.3.1
	github.com/google/uuid v1.3.0
	google.golang.org/grpc v1.52.0
	google.golang.org/protobuf v1.28.1
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/hailocab/go-hostpool v0.0.0-20160125115350-e80d13ce29ed // indirect
	golang.org/x/net v0.5.0 // indirect
	golang.org/x/sys v0.4.0 // indirect
	golang.org/x/text v0.6.0 // indirect
	google.golang.org/genproto v0.0.0-20230119192704-9d59e20e5cd1 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
)

// replace github.com/teleporttacos/proto/pb => ../proto/pb
replace github.com/teleporttacos/ports => ./internal/ports
