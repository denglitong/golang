package main

import (
	"fmt"
	"github.com/alecthomas/inject"
	"github.com/denglitong/golang/inject_demo/db"
)

func main() {
	injector := inject.New()
	injector.Install(
		&db.MongoModule{},
	)
	injector.Call(runShowSampleUsage)

	prepareBindings(injector)
	injector.Call(runShowBindings)
}

func runShowSampleUsage(
	db db.Database,
) {
	if pong, err := db.Ping(); err != nil {
		panic(err)
	} else {
		fmt.Println(pong)
	}
}

func prepareBindings(injector *inject.Injector) {
	var intToBind int = 3
	injector.Bind(intToBind)
	// bind to a value is equivalent to Literal
	var strToBind string = "Hello, world!"
	injector.Bind(inject.Literal(strToBind))

	// binding function will be called each time a float32 is requested:
	// binding a function with return type will bind to its return value
	injector.Bind(func() float32 {
		return 3.14
	})

	// Wrap the function in Singleton() to ensure it is called only once:
	injector.Bind(inject.Singleton(func() float64 {
		return 12.56
	}))

	// in golang, function is the first-class object, we can use it as a type;
	// binding func will bind the implementation(value) to the type of the func signature(type)
	injector.Bind(inject.Literal(func(s1, s2 string) string {
		return fmt.Sprintf("function literal binding: %s => %s", s1, s2)
	}))

	injector.Bind(inject.Mapping(map[string]int{
		"key1": 1,
		"key2": 2,
	}))
	injector.Bind(inject.Sequence([]int{1, 2, 3}))

	// named binding can be achieved with type alias
	injector.Bind(SourcedEventTopic("collections-fct-v1"))
	injector.Bind(DomainEventTopic("collections-uact-v1"))

	// bind interface to implementation
	injector.BindTo((*Greeter)(nil), &DefaultGreeterImpl{Val: "Golang Inject!"})
	injector.BindTo((*CustomerGreeter)(nil), &HiGreeterImpl{Val: "Golang Inject!"})
}

type SourcedEventTopic string
type DomainEventTopic string

type Greeter interface {
	Say() string
}

type CustomerGreeter Greeter

type DefaultGreeterImpl struct {
	Val string
}

func (s *DefaultGreeterImpl) Say() string {
	return "How's going:" + s.Val
}

type HiGreeterImpl struct {
	Val string
}

func (s *HiGreeterImpl) Say() string {
	return "Welcome:" + s.Val
}

func runShowBindings(
	bindingInt int,
	bindingStr string,
	bindingFloat32 float32,
	bindingFloat64 float64,
	bindingFunc func(string, string) string,
	bindingMapString2Int map[string]int,
	bindingSequenceInt []int,
	sourcedEventTopic SourcedEventTopic,
	domainEventTopic DomainEventTopic,
	bindingInterfaceDefaultImpl Greeter,
	bindingInterfaceCustomerImpl CustomerGreeter,
) {
	fmt.Println("bindings int:", bindingInt)
	fmt.Println("bindings Literal string:", bindingStr)
	fmt.Println("bindings float32:", bindingFloat32)
	fmt.Println("bindings Singleton float64:", bindingFloat64)
	fmt.Println("bindings function:", bindingFunc("key", "value"))
	fmt.Println("bindings map[string]int:", bindingMapString2Int)
	fmt.Println("bindings []int:", bindingSequenceInt)
	fmt.Println("bindings string named SourcedEventTopic:", sourcedEventTopic)
	fmt.Println("bindings string named DomainEventTopic:", string(domainEventTopic))
	fmt.Println("bindings interface default implementation:", bindingInterfaceDefaultImpl.Say())
	fmt.Println("bindings interface custome implementation:", bindingInterfaceCustomerImpl.Say())
}
