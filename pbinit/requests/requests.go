package requests

import (
	"context"

	"github.com/cyverse-de/go-mod/gotelnats"
	"github.com/cyverse-de/go-mod/pbinit/common"
	"github.com/cyverse-de/p/go/requests"

	"go.opentelemetry.io/otel/trace"
)

func NewByUsername() *requests.ByUsername {
	return &requests.ByUsername{
		Header: gotelnats.NewHeader(),
	}
}

func NewByUserID() *requests.ByUserID {
	return &requests.ByUserID{
		Header: gotelnats.NewHeader(),
	}
}

func NewNoParams() *requests.NoParams {
	return &requests.NoParams{
		Header: gotelnats.NewHeader(),
	}
}

func NewByUUID() *requests.ByUUID {
	return &requests.ByUUID{
		Header: gotelnats.NewHeader(),
	}
}

func NewByUUIDAndUsername() *requests.ByUUIDAndUsername {
	return &requests.ByUUIDAndUsername{
		Header: gotelnats.NewHeader(),
	}
}

func NewByUUIDAndUserID() *requests.ByUUIDAndUserID {
	return &requests.ByUUIDAndUserID{
		Header: gotelnats.NewHeader(),
	}
}

func NewAssociateByUUIDs() *requests.AssociateByUUIDs {
	return &requests.AssociateByUUIDs{
		Header: gotelnats.NewHeader(),
	}
}

func InitByUsername(request *requests.ByUsername, subject string) (context.Context, trace.Span) {
	if request.Header == nil {
		request.Header = gotelnats.NewHeader()
	}
	return common.Init(request.Header, subject)
}

func InitByUserID(request *requests.ByUserID, subject string) (context.Context, trace.Span) {
	if request.Header == nil {
		request.Header = gotelnats.NewHeader()
	}
	return common.Init(request.Header, subject)
}

func InitNoParams(request *requests.NoParams, subject string) (context.Context, trace.Span) {
	if request.Header == nil {
		request.Header = gotelnats.NewHeader()
	}
	return common.Init(request.Header, subject)
}

func InitByUUID(request *requests.ByUUID, subject string) (context.Context, trace.Span) {
	if request.Header == nil {
		request.Header = gotelnats.NewHeader()
	}
	return common.Init(request.Header, subject)
}

func InitByUUIDAndUsername(request *requests.ByUUIDAndUsername, subject string) (context.Context, trace.Span) {
	if request.Header == nil {
		request.Header = gotelnats.NewHeader()
	}
	return common.Init(request.Header, subject)
}

func InitByUUIDAndUserID(request *requests.ByUUIDAndUserID, subject string) (context.Context, trace.Span) {
	if request.Header == nil {
		request.Header = gotelnats.NewHeader()
	}
	return common.Init(request.Header, subject)
}

func InitAssociateByUUIDs(request *requests.AssociateByUUIDs, subject string) (context.Context, trace.Span) {
	if request.Header == nil {
		request.Header = gotelnats.NewHeader()
	}
	return common.Init(request.Header, subject)
}
