package core

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