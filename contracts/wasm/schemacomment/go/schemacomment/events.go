// COPYRIGHT OF A TEST SCHEMA DEFINITION 1
// COPYRIGHT OF A TEST SCHEMA DEFINITION 2

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package schemacomment

import (
	"github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib"
	"github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"
)

type SchemaCommentEvents struct{}

// header comment for TestEvent 1
// header comment for TestEvent 2
func (e SchemaCommentEvents) TestEvent(
	// header comment for eventParam1 1
	// header comment for eventParam1 2
	eventParam1 string,
	// header comment for eventParam2 1
	// header comment for eventParam2 2
	eventParam2 string,
) {
	evt := wasmlib.NewEventEncoder("schemacomment.testEvent")
	evt.Encode(wasmtypes.StringToString(eventParam1))
	evt.Encode(wasmtypes.StringToString(eventParam2))
	evt.Emit()
}

// header comment for TestEventNoParams 1
// header comment for TestEventNoParams 2
func (e SchemaCommentEvents) TestEventNoParams() {
	evt := wasmlib.NewEventEncoder("schemacomment.testEventNoParams")
	evt.Emit()
}