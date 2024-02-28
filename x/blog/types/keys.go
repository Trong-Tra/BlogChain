package types

const (
	// ModuleName defines the module name
	ModuleName = "blog"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_blog"

	// PostKey defines store and retrieve posts from the db
	PostKey = "Post/value/"

	// PostCountKey defines to keep track of the ID of the lastest post added to the store
	PostCountKey = "Post/count/"
)

var (
	ParamsKey = []byte("p_blog")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
