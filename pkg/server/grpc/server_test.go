//go:generate echo $PWD/$GOPACKAGE/$GOFILE
//go:generate mockgen -package grpc_test -destination $PWD/pkg/server/grpc/mock_storage_test.go github.com/kamilsk/click/pkg/server/grpc ProtectedStorage
package grpc_test
