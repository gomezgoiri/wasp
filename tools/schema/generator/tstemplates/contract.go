package tstemplates

var contractTs = map[string]string{
	// *******************************
	"contract.ts": `
$#emit tsHeader
$#each func FuncNameCall

type Funcs struct{}

var ScFuncs Funcs
$#each func FuncNameForCall
$#if core coreOnload
`,
	// *******************************
	"FuncNameCall": `

type $FuncName$+Call struct {
	Func    *wasmlib.Sc$initFunc$Kind
$#if param MutableFuncNameParams
$#if result ImmutableFuncNameResults
}
`,
	// *******************************
	"MutableFuncNameParams": `
	Params  Mutable$FuncName$+Params
`,
	// *******************************
	"ImmutableFuncNameResults": `
	Results Immutable$FuncName$+Results
`,
	// *******************************
	"FuncNameForCall": `

func (sc Funcs) $FuncName(ctx wasmlib.Sc$Kind$+CallContext) *$FuncName$+Call {
$#if ptrs setPtrs noPtrs
}
`,
	// *******************************
	"coreOnload": `

func OnLoad() {
	exports := wasmlib.NewScExports()
$#each func coreExportFunc
}
`,
	// *******************************
	"coreExportFunc": `
	exports.Add$Kind($Kind$FuncName, wasmlib.$Kind$+Error)
`,
	// *******************************
	"setPtrs": `
	f := &$FuncName$+Call{Func: wasmlib.NewSc$initFunc$Kind(ctx, HScName, H$Kind$FuncName$initMap)}
	f.Func.SetPtrs($paramsID, $resultsID)
	return f
`,
	// *******************************
	"noPtrs": `
	return &$FuncName$+Call{Func: wasmlib.NewSc$initFunc$Kind(ctx, HScName, H$Kind$FuncName$initMap)}
`,
}
