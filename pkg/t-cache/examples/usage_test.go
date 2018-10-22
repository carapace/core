package examples_test

import (
	"fmt"
	"github.com/carapace/core/pkg/t-cache"
)

func ExampleCommitter_Commit() {
	c := cache.New()
	save := c.Lock()
	// defer c.Unlock()

	// if commit is called before Rollback, the transaction is stored. Rollback does not error if already committed or
	// rolled back
	defer save.Rollback()

	c.Set("key", "value")
	save.Commit()

	// normally we'd defer c.Unlock
	c.Unlock()

	fmt.Println(c.Get("key"))
	// Output: value
}

func ExampleCommitter_Rollback() {
	c := cache.New()
	save := c.Lock()
	// defer c.Unlock()

	// if commit is called before Rollback, the transaction is stored. Rollback does not error if already committed or
	// rolled back

	// here we explictly rollback. Normally you'd defer this as well
	// defer save.Rollback()

	c.Set("key", "value")

	// something returned an error!
	save.Rollback()

	// normally we'd defer c.Unlock
	c.Unlock()

	fmt.Println(c.Get("key"))
	// Output: <nil>
}
