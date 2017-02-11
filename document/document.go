package document

// Config is the global configuration object for target
var Config = struct {
	// Retrieve is the global document retriever, override for testing with genRetrieve
	Retrieve retriever
	Reparent bool
}{
	Retrieve: genRetrieve(nil),
	Reparent: false,
}
