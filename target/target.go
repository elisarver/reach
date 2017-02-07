package target

// go:generate gen

// Config is the global configuration object for target
var Config = struct {
	// Retrieve is the global document retriever, override for testing with GenRetrieve
	Retrieve retriever
}{
	Retrieve: GenRetrieve(nil),
}
