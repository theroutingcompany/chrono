package chrono

//go:generate protoc --go_out=. chrono.proto
//go:generate go run cmd/gen.go > stubs.gen.go
