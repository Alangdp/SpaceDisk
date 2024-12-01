package core

// Fnv1aHash computes the FNV-1a hash of the given string and returns the resulting hash as a uint32.
//
// The FNV-1a hash algorithm uses a prime number and a specific offset to process the string byte by byte.
// This function applies the FNV-1a hashing to the input string and returns the 32-bit hash value.
//
// Example:
//     hash := Fnv1aHash("example")
//     fmt.Println("Hash value:", hash)
//
// For more details on the FNV-1a algorithm, refer to:
// https://en.wikipedia.org/wiki/Fowler%E2%80%93Noll%E2%80%93Vo_hash_function
func Fnv1aHash(str string) uint32 {
	const fnvOffset32 = 2166136261
	const fnvPrime32 = 16777619

	hash := uint32(fnvOffset32)
	for i := 0; i < len(str); i++ {
		hash ^= uint32(str[i])
		hash *= fnvPrime32
	}
	return hash
}
