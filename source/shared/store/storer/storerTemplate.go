package storer

type templateData struct {
	ImportPrefix    string
	RecordName      string
	StoringFilePath string
}

var template = `package storer

import (
	_record_ "{{ .ImportPrefix }}/shared/store/record"
)

/* KICKFYNE NOTE WELL:
This defines the interface which is implemented by {{ .RecordName }}Store in {{ .StoringFilePath }}.
There fore, any changes made here must be reflected in {{ .StoringFilePath }}.
*/

// {{ .RecordName }}Storer defines the behavior (API) of a store of /shared/store/_record_.{{ .RecordName }} records.
type {{ .RecordName }}Storer interface {

	// Open opens the data-store.
	// Keeps the file in memory.
	// Returns the error.
	Open() (err error)

	// Close closes the data-store.
	// Returns the error.
	Close() (err error)

	// IsFull returns if the store is full.
	IsFull() (isFull bool)

	// NextID returns the next available ID or an error.
	NextID() (nextID uint64, err error)

	// Get retrieves one _record_.{{ .RecordName }} from the data-store.
	// Param id is the record ID.
	// Returns a _record_.{{ .RecordName }} and error.
	// When no record is found, the returned _record_.{{ .RecordName }} is nil and the returned error is nil.
	Get(id uint64) (r _record_.{{ .RecordName }}, err error)

	// GetAll retrieves all of the _record_.{{ .RecordName }} records from the data-store.
	// Returns a slice of _record_.{{ .RecordName }} and error.
	// When no records are found, the returned slice length is 0 and the returned error is nil.
	GetAll() (rr []_record_.{{ .RecordName }}, err error)

	// Update updates the _record_.{{ .RecordName }} in the data-store.
	// Param newR is the _record_.{{ .RecordName }} to be updated.
	// If newR is a new record then r.ID is updated as well.
	// Returns the error.
	Update(newR _record_.{{ .RecordName }}) (updatedR _record_.{{ .RecordName }}, err error)

	// UpdateAll updates a slice of _record_.{{ .RecordName }} in the data-store.
	// Param newRR is the slice of _record_.{{ .RecordName }} to be updated.
	// If any record in newRR is new then it's ID is updated as well.
	// Returns the error.
	UpdateAll(newRR []_record_.{{ .RecordName }}) (updatedRR []_record_.{{ .RecordName }}, err error)

	// Remove removes the _record_.{{ .RecordName }} from the data-store.
	// Param id is the record ID of the _record_.{{ .RecordName }} to be removed.
	// If the record is not found returns a nil error.
	// Returns the error.
	Remove(id uint64) (err error)
}

`
