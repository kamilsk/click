//go:generate echo $PWD/$GOPACKAGE/$GOFILE
//go:generate mockgen -package server_test -destination $PWD/pkg/server/mock_contract_test.go github.com/kamilsk/click/pkg/server Service
package server_test
