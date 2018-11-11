package chaindb

// ObjectHash returns the latest object hash belonging to a key
func (db *DB) ObjectHash(key string, option *Option) (uint64, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	return db.objectHash(key, option)
}

func (db *DB) objectHash(key string, option *Option) (uint64, error) {
	if option.Cached {
		return db.config.Cache.GetObjHash(key)
	}

	chain, err := db.get(key, option)
	if err != nil {
		return 0, err
	}

	return chain[chain.Len()-1].State.ObjHash, nil
}

// ChainHash returns the current ChainHash belonging to a key
func (db *DB) ChainHash(key string, option *Option) (uint64, error) {
	if option.Cached {
		return db.config.Cache.GetChainHash(key)
	}

	chain, err := db.get(key, option)
	if err != nil {
		return 0, err
	}
	return chain[chain.Len()-1].State.ChainHash, nil

}

// Get returns a chain from the DB, or an error if it is not present
func (db *DB) Get(key string, option *Option) (Chain, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	if option == nil {
		option = defaultOpt()
	}

	return db.get(key, option)
}

// get does the same as Get, but is not threadsafe and does not set a default op
func (db *DB) get(key string, option *Option) (Chain, error) {
	chunks, err := db.config.Store.Get(key)

	if err != nil {
		return nil, err
	}

	if len(chunks) == 0 {
		return nil, ErrKeyNotExist
	}

	ok, _, err := db.CheckIntegrity(chunks)
	if err != nil {

		if !ok {
			return chunks, err
		}
		// no need to return the chain if the error was not an integrity error, it is an io err or something then.
		return nil, err
	}
	return chunks, nil
}
