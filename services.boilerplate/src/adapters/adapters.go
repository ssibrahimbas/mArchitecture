package adapters

import (
	"clean-boilerplate/boilerplate/src/adapters/memory"
	"clean-boilerplate/boilerplate/src/adapters/mysql"
)

var (
	MySQL  = mysql.New()
	Memory = memory.New()
)
