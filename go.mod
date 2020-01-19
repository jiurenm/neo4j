module neo4j

go 1.12

require (
	github.com/golang/protobuf v1.3.2
	github.com/google/wire v0.4.0
	github.com/johnnadratowski/golang-neo4j-bolt-driver v0.0.0-20181101021923-6b24c0085aae
	github.com/kr/pretty v0.2.0 // indirect
	golang.org/x/net v0.0.0-20190404232315-eb5bcb51f2a3
	google.golang.org/grpc v1.19.0
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/yaml.v2 v2.2.7
)

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190605123033-f99c8df09eb5
	golang.org/x/net => github.com/golang/net v0.0.0-20190603091049-60506f45cf65
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190602015325-4c4f7f33c9ed
	golang.org/x/text => github.com/golang/text v0.3.2
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190606050223-4d9ae51c2468
	google.golang.org/grpc => github.com/grpc/grpc-go v1.21.1
)
