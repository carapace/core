/*
package chain-db defines an append only K,V store, permanently storing all versions of the object

An object in chain-db is a proto.Message, marshalled into an any.Any and embedded with metadata on state and key.

When querying, a chain of all puts on that object is returned, and the chain is verified for integrity. It is
up to the caller to unmarshal their values into appropriate structs.
*/
package append
