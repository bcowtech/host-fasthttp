package middleware

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/bcowtech/host"
	"github.com/bcowtech/host-fasthttp/internal"
	"github.com/bcowtech/structproto"
	"github.com/bcowtech/structproto/tagresolver"
	"github.com/bcowtech/structproto/util/reflectutil"
)

var _ structproto.StructBinder = new(ResourceManagerBinder)

type ResourceManagerBinder struct {
	router     internal.Router
	appContext *host.AppContext
}

func (b *ResourceManagerBinder) Init(context *structproto.StructProtoContext) error {
	return nil
}

func (b *ResourceManagerBinder) Bind(field structproto.FieldInfo, rv reflect.Value) error {
	if !rv.IsValid() {
		return fmt.Errorf("specifiec argument 'rv' is invalid")
	}

	// assign zero if rv is nil
	rvResource := reflectutil.AssignZero(rv)
	binder := &ResourceBinder{
		resourceType: rvResource.Type().Name(),
		components: map[string]reflect.Value{
			host.APP_CONFIG_FIELD:           b.appContext.Config(),
			host.APP_SERVICE_PROVIDER_FIELD: b.appContext.ServiceProvider(),
		},
	}
	err := b.preformBindResource(rvResource, binder)
	if err != nil {
		return err
	}

	// register RequestHandlers
	return b.registerRoute(field.Name(), rvResource)
}

func (b *ResourceManagerBinder) Deinit(context *structproto.StructProtoContext) error {
	return nil
}

func (b *ResourceManagerBinder) preformBindResource(target reflect.Value, binder *ResourceBinder) error {
	prototype, err := structproto.Prototypify(target,
		&structproto.StructProtoResolveOption{
			TagResolver: tagresolver.NoneTagResolver,
		})
	if err != nil {
		return err
	}

	err = prototype.Bind(binder)
	if err != nil {
		return err
	}

	// TODO: see if work, when move the following statements into ResourceBinder.Deinit()
	// call resource.Init()
	ctx := structproto.StructProtoContext(*prototype)
	rv := ctx.Target()
	if rv.CanAddr() {
		rv = rv.Addr()
		// call resource.Init()
		fn := rv.MethodByName(host.APP_COMPONENT_INIT_METHOD)
		if fn.IsValid() {
			if fn.Kind() != reflect.Func {
				log.Fatalf("[bcowtech/host-fasthttp] cannot find func %s() within type %s\n", host.APP_COMPONENT_INIT_METHOD, rv.Type().String())
			}
			if fn.Type().NumIn() != 0 || fn.Type().NumOut() != 0 {
				log.Fatalf("[bcowtech/host-fasthttp] %s.%s() type should be func()\n", rv.Type().String(), host.APP_COMPONENT_INIT_METHOD)
			}
			fn.Call([]reflect.Value(nil))
		}
	}
	return nil
}

func (b *ResourceManagerBinder) registerRoute(url string, rvResource reflect.Value) error {
	// register RequestHandlers
	count := rvResource.Type().NumMethod()
	for i := 0; i < count; i++ {
		method := rvResource.Type().Method(i)

		rvMethod := rvResource.Method(method.Index)
		if isRequestHandler(rvMethod) {
			handler := asRequestHandler(rvMethod)
			if handler != nil {
				// TODO: validate path make comply RFC3986
				b.router.Add(strings.ToUpper(method.Name), url, handler)
			}
		}
	}
	return nil
}
