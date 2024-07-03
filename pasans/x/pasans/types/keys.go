package types

const (
	// ModuleName defines the module name
	ModuleName = "pasans"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_pasans"
)

var (
	ParamsKey = []byte("p_pasans")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
