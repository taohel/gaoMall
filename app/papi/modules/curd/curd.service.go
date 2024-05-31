package curd

import (
	"gaoMall/app"
)

func Find(model interface{}, where ...interface{}) interface{} {
	app.DBR().Find(model)
	return model
}
