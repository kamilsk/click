//go:generate echo $PWD/$GOPACKAGE/$GOFILE
//go:generate mockgen -package chi_test -destination $PWD/server/router/chi/mock_contract_test.go github.com/kamilsk/click/server/router Server
package chi_test
