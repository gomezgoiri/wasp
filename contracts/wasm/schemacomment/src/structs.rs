// COPYRIGHT OF A TEST SCHEMA DEFINITION 1
// COPYRIGHT OF A TEST SCHEMA DEFINITION 2


// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

#![allow(dead_code)]
#![allow(unused_imports)]

use wasmlib::*;

// comment for TestStruct1
#[derive(Clone)]
pub struct TestStruct1 {
    pub x1 : i32, // comment for x1
    pub y1 : i32, // comment for y1
}

impl TestStruct1 {
    pub fn from_bytes(bytes: &[u8]) -> TestStruct1 {
        let mut dec = WasmDecoder::new(bytes);
        TestStruct1 {
            x1 : int32_decode(&mut dec),
            y1 : int32_decode(&mut dec),
        }
    }

    pub fn to_bytes(&self) -> Vec<u8> {
        let mut enc = WasmEncoder::new();
		int32_encode(&mut enc, self.x1);
		int32_encode(&mut enc, self.y1);
        enc.buf()
    }
}

#[derive(Clone)]
pub struct ImmutableTestStruct1 {
    pub(crate) proxy: Proxy,
}

impl ImmutableTestStruct1 {
    pub fn exists(&self) -> bool {
        self.proxy.exists()
    }

    pub fn value(&self) -> TestStruct1 {
        TestStruct1::from_bytes(&self.proxy.get())
    }
}

#[derive(Clone)]
pub struct MutableTestStruct1 {
    pub(crate) proxy: Proxy,
}

impl MutableTestStruct1 {
    pub fn delete(&self) {
        self.proxy.delete();
    }

    pub fn exists(&self) -> bool {
        self.proxy.exists()
    }

    pub fn set_value(&self, value: &TestStruct1) {
        self.proxy.set(&value.to_bytes());
    }

    pub fn value(&self) -> TestStruct1 {
        TestStruct1::from_bytes(&self.proxy.get())
    }
}

// comment for TestStruct2
#[derive(Clone)]
pub struct TestStruct2 {
    pub x2 : i32, // comment for x2
    pub y2 : i32, // comment for y2
}

impl TestStruct2 {
    pub fn from_bytes(bytes: &[u8]) -> TestStruct2 {
        let mut dec = WasmDecoder::new(bytes);
        TestStruct2 {
            x2 : int32_decode(&mut dec),
            y2 : int32_decode(&mut dec),
        }
    }

    pub fn to_bytes(&self) -> Vec<u8> {
        let mut enc = WasmEncoder::new();
		int32_encode(&mut enc, self.x2);
		int32_encode(&mut enc, self.y2);
        enc.buf()
    }
}

#[derive(Clone)]
pub struct ImmutableTestStruct2 {
    pub(crate) proxy: Proxy,
}

impl ImmutableTestStruct2 {
    pub fn exists(&self) -> bool {
        self.proxy.exists()
    }

    pub fn value(&self) -> TestStruct2 {
        TestStruct2::from_bytes(&self.proxy.get())
    }
}

#[derive(Clone)]
pub struct MutableTestStruct2 {
    pub(crate) proxy: Proxy,
}

impl MutableTestStruct2 {
    pub fn delete(&self) {
        self.proxy.delete();
    }

    pub fn exists(&self) -> bool {
        self.proxy.exists()
    }

    pub fn set_value(&self, value: &TestStruct2) {
        self.proxy.set(&value.to_bytes());
    }

    pub fn value(&self) -> TestStruct2 {
        TestStruct2::from_bytes(&self.proxy.get())
    }
}