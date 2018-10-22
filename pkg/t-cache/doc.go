/*
package cache implements a simple, transactional in-memory cache.

Calling Lock() on the cache will return a Committer. Either Commit or Rollback must be called on the committer
before releasing the lock, else a panic is thrown

The cache is based on a simple map, and should not be used for large workloads.
*/
package cache
