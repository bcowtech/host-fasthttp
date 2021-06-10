package middleware

import (
	"fmt"
	"reflect"

	"github.com/bcowtech/structproto"
	"github.com/bcowtech/structproto/util/reflectutil"
)

var _ structproto.StructBinder = new(ResourceBinder)

type ResourceBinder struct {
	resourceType string
	components   map[string]reflect.Value
}

func (b *ResourceBinder) Init(context *structproto.StructProtoContext) error {
	return nil
}

func (b *ResourceBinder) Bind(field structproto.FieldInfo, target reflect.Value) error {
	if v, ok := b.components[field.Name()]; ok {
		if !target.IsValid() {
			return fmt.Errorf("specifiec argument 'target' is invalid. cannot bind '%s' to '%s'",
				field.Name(),
				b.resourceType)
		}

		target = reflectutil.AssignZero(target)
		if v.Type().ConvertibleTo(target.Type()) {
			target.Set(v.Convert(target.Type()))
		}
	}
	return nil
}

func (b *ResourceBinder) Deinit(context *structproto.StructProtoContext) error {
	return nil
}
