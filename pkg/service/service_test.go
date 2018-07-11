//go:generate echo $PWD/$GOPACKAGE/$GOFILE
//go:generate mockgen -package service_test -destination $PWD/pkg/service/mock_contract_test.go github.com/kamilsk/click/pkg/service Storage
package service_test
