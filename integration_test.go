//go:generate echo $PWD/$GOPACKAGE/$GOFILE
//go:generate mockgen -package main -destination $PWD/mock_storage_test.go github.com/kamilsk/click/service Storage
package main
