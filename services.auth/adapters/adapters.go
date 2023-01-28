package adapters

import (
	"github.ssibrahimbas/mArchitecture/auth/adapters/memory"
	"github.ssibrahimbas/mArchitecture/auth/adapters/mongo"
	"github.ssibrahimbas/mArchitecture/auth/adapters/mysql"
)

var (
	MySQL  = mysql.New()
	Memory = memory.New()
	Mongo  = mongo.New()
)
