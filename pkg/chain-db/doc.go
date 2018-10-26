/*
package chain-db defines an append only K,V store, permanently storing all versions of the object

An object in chain-db starts out by defining all possible fields through protocol buffers, and then only being
able to amend the object. Gets always return a Chain object, which contains every single put on the object, verifyable by
the hashes
*/
package append
