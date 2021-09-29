## Smart Contract Initialization

Smart contracts start out with a completely blank state. Sometimes you want to be able to
define initial state, for example if your contract is configurable. You want to be able to
pass this configuration to the contract upon deployment, so that its state reflects that
configuration once the first request comes in. To support this initialization we allow the
smart contract creator to provide an optional function named `init`.

This `init` function, when provided, will automatically be called immediately after the
first time the contract has been deployed to the VM. Note that this is a one-time
initialization call, meant to be performed by the contract deployment mechanism. ISCP will
prevent anyone else from calling this function ever again. So if you need to be able to
reconfigure the contract later on, then you will need to provide a separate configuration
function, and guard it from being accessed by anyone else than properly authorized
entities.

To show how creating a smart contract with WasmLib works we will slowly start fleshing out
the smart contract functions of the `dividend` example in this tutorial. Here is the first
part of the Go code that implements it, which contains the 'init' function:

```golang
// 'init' is used as a way to initialize a smart contract. It is an optional
// function that will automatically be called upon contract deployment. In this
// case we use it to initialize the 'owner' state variable so that we can later
// use this information to prevent non-owners from calling certain functions.
// The 'init' function takes a single optional parameter:
// - 'owner', which is the agent id of the entity owning the contract.
// When this parameter is omitted the owner will default to the contract creator.
func funcInit(ctx wasmlib.ScFuncContext, f *InitContext) {
    // The schema tool has already created a proper InitContext for this function that
    // allows us to access call parameters and state storage in a type-safe manner.
    
    // First we set up a default value for the owner in case the optional
    // 'owner' parameter was omitted.
    var owner wasmlib.ScAgentID = ctx.ContractCreator()
    
    // Now we check if the optional 'owner' parameter is present in the params map.
    if f.Params.Owner().Exists() {
        // Yes, it was present, so now we overwrite the default owner with
        // the one specified by the 'owner' parameter.
        owner = f.Params.Owner().Value()
    }
    
    // Now that we have sorted out which agent will be the owner of this contract
    // we will save this value in the 'owner' variable in state storage on the host.
    // Read the documentation on schema.json to understand why this state variable is
    // supported at compile-time by code generated from schema.json by the schema tool.
    f.State.Owner().SetValue(owner)
}
```

We define an owner variable and allow it to be something other than the default value of
contract creator. It is always a good idea to be flexible enough to be able to transfer
ownership to another entity if necessary. Remember that once a smart contract has been
deployed it is no longer possible to change this. Therefore, it is good practice thinking
through those situations that could require change in advance and allow the contract
itself to handle such changes through its state by providing a proper function interface:

```golang
// 'setOwner' is used to change the owner of the smart contract.
// It updates the 'owner' state variable with the provided agent id.
// The 'setOwner' function takes a single mandatory parameter:
// - 'owner', which is the agent id of the entity that will own the contract.
// Only the current owner can change the owner.
func funcSetOwner(ctx wasmlib.ScFuncContext, f *SetOwnerContext) {
    // Note that the schema tool has already dealt with making sure that this function
    // can only be called by the owner and that the required parameter is present.
    // So once we get to this point in the code we can take that as a given.
    
    // Save the new owner parameter value in the 'owner' variable in state storage.
    f.State.Owner().SetValue(f.Params.Owner().Value())
}
```

Note that we only define a single owner here. Proper fall-back could require multiple
owners in case the owner entity could disappear, which would allow others to take over
instead of the contract becoming immutable w.r.t. owner functionality. Again, we cannot
stress enough how important it is to think through every aspect of a smart contract before
deployment.

In the next section we will look at how a smart contract can transfer tokens.

Next: [Token Transfers](transfers.md)
