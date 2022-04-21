package grocery

import (
	"github.com/google/uuid"
)

// StoreOptions provides options that may be passed to StoreWithOptions if the
// default behavior of Store needs to be changed.
type StoreOptions struct {
	// ID is the ID this object must be stored with. Store will fail if an
	// object with the type of the pointer being passed already exists with
	// this ID, unless Overwrite is set to true.
	ID string

	// Set to true if you would like to overwrite existing objects stored with
	// this ID.
	Overwrite bool

	// All other options inherit from UpdateOptions.
	*UpdateOptions
}

// Store saves an object in Redis. As with all other Grocery operations, the
// name of the pointer's struct type is used as a prefix for the object's key.
// The object's ID is then randomly generated, and the object is stored at
// prefix:id. If you would like to set a specific ID, use StoreWithOptions.
func Store(ptr interface{}) (string, error) {
	id := uuid.NewString()
	return id, StoreWithOptions(ptr, &StoreOptions{id, false, nil})
}

// StoreWithOptions saves an object in Redis, like Store, but with options.
func StoreWithOptions(ptr interface{}, opts *StoreOptions) error {
	if opts.UpdateOptions == nil {
		opts.UpdateOptions = &UpdateOptions{}
	}

	opts.UpdateOptions.isStore = true
	opts.UpdateOptions.storeOverwrite = opts.Overwrite

	return updateInternal(opts.ID, ptr, opts.UpdateOptions)
}
