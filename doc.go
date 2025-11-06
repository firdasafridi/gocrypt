// Package gocrypt provide a library for do encryption in struct with go field.
//
// # Package gocrypt provide in struct tag encryption or inline encryption and decryption
//
// The package supported:
//
// # DES3 — Triple Data Encryption Standard
//
// # AES — Advanced Encryption Standard
//
// AES-256-GCM — Advanced Encryption Standard with 256-bit key and Galois/Counter Mode (cross-language compatible)
//
// # RC4 — stream chipper
//
// The AES cipher is the current U.S. government standard for all software, and is recognized worldwide.
//
// The AES-256-GCM implementation (tag: aes256gcm) is designed for cross-language compatibility,
// ensuring encrypted data can be decrypted in other languages (e.g., JavaScript) using the same format.
//
// The DES ciphers are primarily supported for PBE standard that provides the option of generating an encryption key based on a passphrase.
//
// The RC4 is supplied for situations that call for fast encryption, but not strong encryption. RC4 is ideal for situations that require a minimum of encryption.
package gocrypt
