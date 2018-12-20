//go:generate echo $PWD - $GOPACKAGE - $GOFILE
//go:generate mockgen -package server_test -destination mock_service_test.go github.com/kamilsk/click/pkg/server Service
package server_test
