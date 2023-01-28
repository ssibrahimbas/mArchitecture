package adapters

import (
	"github.ssibrahimbas/mArchitecture/boilerplate/src/adapters/memory"
	"github.ssibrahimbas/mArchitecture/boilerplate/src/adapters/mysql"
)

var (
	MySQL  = mysql.New()
	Memory = memory.New()
)
