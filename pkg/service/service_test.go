//go:generate echo $PWD/$GOPACKAGE/$GOFILE
//go:generate mockgen -package service_test -destination $PWD/pkg/service/mock_contract_test.go github.com/kamilsk/click/pkg/service Storage
package service_test

import (
	_ "github.com/golang/mock/gomock"
	_ "github.com/golang/mock/mockgen/model"
)
