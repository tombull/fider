package bus_test

import (
	"context"
	"testing"

	"github.com/tombull/teamdream/app/models/cmd"

	"github.com/tombull/teamdream/app/pkg/bus"
	"github.com/tombull/teamdream/app/pkg/errors"

	. "github.com/tombull/teamdream/app/pkg/assert"
)

func TestBus_SimpleMessage(t *testing.T) {
	RegisterT(t)
	bus.Register(GreeterService{})
	bus.Init()
	ctx := context.WithValue(context.Background(), GreetingKey, "Good Morning")
	cmd := &SayHelloCommand{Name: "Teamdream"}
	err := bus.Dispatch(ctx, cmd)
	Expect(err).IsNil()
	Expect(cmd.Result).Equals("Good Morning Teamdream")
}

func TestBus_MessageIsNotPointer_ShouldPanic(t *testing.T) {
	RegisterT(t)
	bus.Register(GreeterService{})
	bus.Init()

	defer func() {
		if r := recover(); r != nil {
			panicText := (r.(error)).Error()
			Expect(panicText).Equals("'github.com/tombull/teamdream/app/pkg/bus_test.SayHelloCommand' is not a pointer")
		}
	}()

	cmd := SayHelloCommand{Name: "Teamdream"}
	err := bus.Dispatch(context.Background(), cmd)
	Expect(err).IsNil()
}

func TestBus_OverwriteService(t *testing.T) {
	RegisterT(t)

	bus.Register(GreeterService{})
	bus.Register(BetterGreeterService{})
	bus.Init()
	cmd := &SayHelloCommand{Name: "Teamdream"}
	err := bus.Dispatch(context.Background(), cmd)
	Expect(err).IsNil()
	Expect(cmd.Result).Equals("Hello Teamdream")
}

func TestBus_MultipleMessages(t *testing.T) {
	RegisterT(t)

	bus.Register(BetterGreeterService{})
	bus.Init()
	cmd1 := &SayHelloCommand{Name: "John"}
	cmd2 := &SayHelloCommand{Name: "Mary"}
	cmd3 := &SayHelloCommand{Name: "Bob"}
	err := bus.Dispatch(context.Background(), cmd1, cmd2, cmd3)
	Expect(err).IsNil()
	Expect(cmd1.Result).Equals("Hello John")
	Expect(cmd2.Result).Equals("Hello Mary")
	Expect(cmd3.Result).Equals("Hello Bob")
}

func TestBus_MultipleListeners(t *testing.T) {
	RegisterT(t)
	value1 := ""
	bus.AddListener(func(ctx context.Context, c *SayHelloCommand) {
		value1 = c.Name
	})

	value2 := ""
	bus.AddListener(func(ctx context.Context, c *SayHelloCommand) {
		value2 = c.Name
	})

	bus.Publish(context.Background(), &SayHelloCommand{Name: "Teamdream"})
	Expect(value1).Equals("Teamdream")
	Expect(value2).Equals("Teamdream")
}

func TestBus_MultiplePublish(t *testing.T) {
	RegisterT(t)
	value1 := ""
	bus.AddListener(func(ctx context.Context, c *SayHelloCommand) {
		value1 += c.Name
	})

	value2 := ""
	bus.AddListener(func(ctx context.Context, c *SayHelloCommand) {
		value2 += c.Name
	})

	bus.Publish(context.Background(), &SayHelloCommand{Name: "123"}, &SayHelloCommand{Name: "456"})
	Expect(value1).Equals("123456")
	Expect(value2).Equals("123456")
}

func TestBus_PublishError(t *testing.T) {
	RegisterT(t)
	boom := errors.New("BOOM")

	var err error
	bus.AddListener(func(ctx context.Context, c *cmd.LogError) {
		err = c.Err
	})

	bus.AddListener(func(ctx context.Context, c *SayHelloCommand) error {
		return boom
	})

	bus.Publish(context.Background(), &SayHelloCommand{Name: "123"})
	Expect(errors.Cause(err)).Equals(boom)
}
