package log

import (
	"context"

	"shopping-cart/common/log"
	"shopping-cart/common/server"
)

const (
	// nameField nameField
	nameField = "name"
)

// newLogger -
func newLogger(name ...string) *log.Logger {
	if len(name) == 0 {
		return log.StdLogger()
	}
	fields := []log.Field{
		log.ZapString(nameField, name[0]),
	}
	var fieldsInterface []interface{}
	fieldsInterface = append(fieldsInterface, nameField, name[0])
	return log.StdLogger().WithFields(fields, fieldsInterface)
}

// GetLogger -
func GetLogger(ctx context.Context, name string) *log.Logger {
	if srv := server.FromContext(ctx); srv != nil {
		if len(name) == 0 {
			return srv.Log().WithFields(nil, nil)
		}
		fields := []log.Field{
			log.ZapString(nameField, name),
		}
		var fieldsInterface []interface{}
		fieldsInterface = append(fieldsInterface, nameField, name)
		return srv.Log().WithFields(fields, fieldsInterface)
	}
	return newLogger(name)
}
