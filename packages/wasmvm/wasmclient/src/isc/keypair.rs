// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use crypto::hashes::{blake2b::Blake2b256, Digest};
use crypto::signatures::ed25519;
use std::{convert::TryInto, fmt::Debug};
use wasmlib::*;

pub struct KeyPair {
    private_key: ed25519::SecretKey,
    pub public_key: ed25519::PublicKey,
}

impl KeyPair {
    pub fn address(&self) -> ScAddress {
        let mut addr: Vec<u8> = Vec::with_capacity(wasmlib::SC_LENGTH_ED25519);
        addr[0] = wasmlib::SC_ADDRESS_ED25519;
        let hash = Blake2b256::digest(self.public_key.to_bytes());
        addr.copy_from_slice(&hash[..]);
        return wasmlib::address_from_bytes(&addr);
    }
    pub fn sign(&self, data: &[u8]) -> Vec<u8> {
        return self.private_key.sign(data).to_bytes().to_vec();
    }
    pub fn from_sub_seed(seed: &[u8], n: u64) -> KeyPair {
        let index_bytes = uint64_to_bytes(n);
        let mut hash_of_index_bytes = Blake2b256::digest(index_bytes.to_owned());
        for i in 0..seed.len() {
            hash_of_index_bytes[i] ^= seed[i];
        }
        let public_key =
            ed25519::PublicKey::try_from_bytes(hash_of_index_bytes.try_into().unwrap()).unwrap();
        let private_key = ed25519::SecretKey::from_bytes(hash_of_index_bytes.try_into().unwrap());
        return KeyPair {
            public_key: public_key,
            private_key: private_key,
        };
    }
}

impl Clone for KeyPair {
    fn clone(&self) -> Self {
        return KeyPair {
            private_key: ed25519::SecretKey::from_bytes(self.private_key.to_bytes()),
            public_key: self.public_key.clone(),
        };
    }
}

impl Debug for KeyPair {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> Result<(), std::fmt::Error> {
        f.debug_tuple("KeyPair").field(&self.public_key).finish()
    }
}

impl PartialEq for KeyPair {
    fn eq(&self, other: &Self) -> bool {
        // FIXME this may not be enough
        return self.public_key == other.public_key;
    }
}
