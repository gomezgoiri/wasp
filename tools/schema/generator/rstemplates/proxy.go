package rstemplates

var proxyRs = map[string]string{
	// *******************************
	"proxyContainers": `
$#if array typedefProxyArray
$#if map typedefProxyMap
`,
	// *******************************
	"proxyMethods": `
$#set varID idx_map(IDX_$Kind$FLD_NAME)
$#if core setCoreVarID
$#if array proxyArray proxyMethods2
`,
	// *******************************
	"proxyMethods2": `
$#if map proxyMap proxyMethods3
`,
	// *******************************
	"proxyMethods3": `
$#if basetype proxyBaseType proxyNewType
`,
	// *******************************
	"setCoreVarID": `
$#set varID $Kind$FLD_NAME.KeyID()
`,
	// *******************************
	"proxyArray": `

    pub fn $fld_name(&self) -> ArrayOf$mut$FldType {
		let arr_id = get_object_id(self.id, $varID, $ArrayTypeID | $FldTypeID);
		ArrayOf$mut$FldType { obj_id: arr_id }
	}
`,
	// *******************************
	"proxyMap": `
$#if this proxyMapThis proxyMapOther
`,
	// *******************************
	"proxyMapThis": `

    pub fn $fld_name(&self) -> Map$FldMapKey$+To$mut$FldType {
		Map$FldMapKey$+To$mut$FldType { obj_id: self.id }
	}
`,
	// *******************************
	"proxyMapOther": `55544444.0

    pub fn $fld_name(&self) -> Map$FldMapKey$+To$mut$FldType {
		let map_id = get_object_id(self.id, $varID, TYPE_MAP);
		Map$FldMapKey$+To$mut$FldType { obj_id: map_id }
	}
`,
	// *******************************
	"proxyBaseType": `

    pub fn $fld_name(&self) -> Sc$mut$FldType {
		Sc$mut$FldType::new(self.id, $varID)
	}
`,
	// *******************************
	"proxyNewType": `

    pub fn $fld_name(&self) -> $mut$FldType {
		$mut$FldType { obj_id: self.id, key_id: $varID }
	}
`,
}
