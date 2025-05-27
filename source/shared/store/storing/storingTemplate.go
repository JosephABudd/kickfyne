package storing

type templateData struct {
	ImportPrefix   string
	RecordName     string
	StorerFilePath string
}

var template = `package storing

import (
	"bytes"
	"fmt"
	"math"
	"sort"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/storage"
	"gopkg.in/yaml.v3"

	_paths_ "{{ .ImportPrefix }}/shared/paths"
	_record_ "{{ .ImportPrefix }}/shared/store/record"
)


/* KICKFYNE NOTE WELL:
This is the implementation of the {{ .RecordName }}Storer interface defined in {{ .StorerFilePath }}.
There fore, any changes made here must be reflected in {{ .StorerFilePath }}.
*/

var Err{{ .RecordName }}Full = fmt.Errorf("{{ .RecordName }}Store is full")

type by{{ .RecordName }}ID []_record_.{{ .RecordName }}

func (bcn by{{ .RecordName }}ID) Len() int           { return len(bcn) }
func (bcn by{{ .RecordName }}ID) Swap(i, j int)      { bcn[i], bcn[j] = bcn[j], bcn[i] }
func (bcn by{{ .RecordName }}ID) Less(i, j int) bool { return bcn[i].ID < bcn[j].ID }

type {{ .RecordName }}Data struct {
	LastID  uint64
	Records []_record_.{{ .RecordName }}
}

// {{ .RecordName }}Store is the API of the {{ .RecordName }} store.
// It is the implementation of the interface in /domain/store/storer/{{ .RecordName }}.go.
type {{ .RecordName }}Store struct {
	uri  fyne.URI
	lock sync.Mutex
	data {{ .RecordName }}Data
}

// New{{ .RecordName }}Store constructs a new {{ .RecordName }}Store.
// Param db is an open bolt data-store.
// Returns a pointer to the new {{ .RecordName }}Store.
func New{{ .RecordName }}Store() (store *{{ .RecordName }}Store) {
	store = &{{ .RecordName }}Store{
		uri: _paths_.StoreURI("{{ .RecordName }}.yaml"),
	}
	return
}

// IsFull returns if the store is full.
func (store *{{ .RecordName }}Store) IsFull() (isFull bool) {
	isFull = (store.data.LastID == math.MaxUint64)
	return
}

// NextID returns the next available id.
// Returns the error if there are no more ids.
func (store *{{ .RecordName }}Store) NextID() (nextID uint64, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("{{ .RecordName }}Store.NextID: %w", err)
		}
	}()

	store.lock.Lock()
	defer store.lock.Unlock()

	nextID, err = store.nextID()
	return
}

// Open opens the bolt data-store.
// Keeps the file in memory.
// Returns the error.
func (store *{{ .RecordName }}Store) Open() (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("{{ .RecordName }}Store.Open: %w", err)
		}
	}()

	store.lock.Lock()
	defer store.lock.Unlock()

	if err = store.readAll(); err != nil {
		return
	}
	sort.Sort(by{{ .RecordName }}ID(store.data.Records))

	return
}

// Close closes the data-store.
// Returns the error.
func (store *{{ .RecordName }}Store) Close() (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("{{ .RecordName }}Store.Close: %w", err)
		}
	}()

	// The store is always closed.
	return
}

// Get retrieves one _record_.{{ .RecordName }} from the data-store.
// Param id is the record ID.
// Returns a _record_.{{ .RecordName }} and error.
// When no record is found, the returned _record_.{{ .RecordName }} is nil and the returned error is nil.
// Use {{ .RecordName }}.IsZero() to determine if the returned record is zero meaning not found.
func (store *{{ .RecordName }}Store) Get(id uint64) (r _record_.{{ .RecordName }}, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("{{ .RecordName }}Store.Get: %w", err)
		}
	}()

	store.lock.Lock()
	defer store.lock.Unlock()

	for _, rec := range store.data.Records {
		if rec.ID == id {
			r = rec
			return
		}
	}
	// Not found. No error.
	return
}

// GetAll retrieves all of the _record_.{{ .RecordName }} records from the data-store.
// Returns a slice of _record_.{{ .RecordName }} and error.
// When no records are found, the returned slice length is 0 and the returned error is nil.
func (store *{{ .RecordName }}Store) GetAll() (rr []_record_.{{ .RecordName }}, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("{{ .RecordName }}Store.GetAll: %w", err)
		}
	}()

	store.lock.Lock()
	defer store.lock.Unlock()

	rr = make([]_record_.{{ .RecordName }}, len(store.data.Records))
	copy(rr, store.data.Records)
	return
}

// Update updates the _record_.{{ .RecordName }} in the data-store.
// Param newR is the _record_.{{ .RecordName }} to be updated.
// If newR is a new record then updatedR has the new ID.
// Returns the updated record and the error.
func (store *{{ .RecordName }}Store) Update(newR _record_.{{ .RecordName }}) (updatedR _record_.{{ .RecordName }}, err error) {

	defer func() {
		if err == nil {
			sort.Sort(by{{ .RecordName }}ID(store.data.Records))
			err = store.writeAll()
		}
		if err != nil {
			err = fmt.Errorf("{{ .RecordName }}Store.Update: %w", err)
		}
	}()

	store.lock.Lock()
	defer store.lock.Unlock()

	updatedR = newR

	// Add or replace the _record_.
	// Updating an existing record so replace it.
	if updatedR.ID != 0 {
		for i, r := range store.data.Records {
			if r.ID == updatedR.ID {
				// Found the record so just update it.
				store.data.Records[i] = updatedR
				return
			}
		}
	}
	// New _record_.
	if updatedR.ID == 0 {
		if updatedR.ID, err = store.nextID(); err != nil {
			return
		}
	}
	// Add this new record to the list of records and sort it.
	store.data.Records = append(store.data.Records, updatedR)
	return
}

// UpdateAll updates a slice of _record_.{{ .RecordName }} in the data-store.
// Param newRR is the slice of _record_.{{ .RecordName }} to be updated.
// Returns the updated version of each added _record_.
// Returns the error.
func (store *{{ .RecordName }}Store) UpdateAll(newRR []_record_.{{ .RecordName }}) (updatedRR []_record_.{{ .RecordName }}, err error) {

	defer func() {
		if err == nil {
			sort.Sort(by{{ .RecordName }}ID(store.data.Records))
			err = store.writeAll()
		}
		if err != nil {
			err = fmt.Errorf("{{ .RecordName }}Store.UpdateAll: %w", err)
		}
	}()

	store.lock.Lock()
	defer store.lock.Unlock()

	updatedRR = make([]_record_.{{ .RecordName }}, len(newRR))
	copy(updatedRR, newRR)

	for _, updatedR := range updatedRR {
		switch {
		case updatedR.ID == 0:
			// New record without an id.
			if updatedR.ID, err = store.nextID(); err != nil {
				return
			}
			store.data.Records = append(store.data.Records, updatedR)
		case updatedR.ID != 0:
			found := false
			// Updating an existing record so replace it.
			for i, r := range store.data.Records {
				if r.ID == updatedR.ID {
					found = true
					store.data.Records[i] = updatedR
					break
				}
			}
			if !found {
				// New record with an id.
				store.data.Records = append(store.data.Records, updatedR)
			}
		}
	}
	return
}

// Remove removes the _record_.{{ .RecordName }} from the data-store.
// Param id is the record ID of the _record_.{{ .RecordName }} to be removed.
// If the record is not found returns a nil error.
// Returns the error.
func (store *{{ .RecordName }}Store) Remove(id uint64) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("{{ .RecordName }}Store.Remove: %w", err)
		}
	}()

	store.lock.Lock()
	defer store.lock.Unlock()

	// Find the _record_.
	var found bool
	var records []_record_.{{ .RecordName }}
	l := len(store.data.Records)
	for i, r := range store.data.Records {
		if r.ID == id {
			found = true
			records = make([]_record_.{{ .RecordName }}, l-1)
			if i > 0 {
				// Copy records preceding this unwanted _record_.
				copy(records, store.data.Records[:i])
			}
			// Skip over this unwanted _record_.
			if i++; i < l {
				// Copy records following this unwanted _record_.
				records = append(records, store.data.Records[i:]...)
			}
			break
		}
	}
	if !found {
		// No error if not found.
		return
	}
	store.data.Records = records
	err = store.writeAll()
	return
}

func (store *{{ .RecordName }}Store) readAll() (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("{{ .RecordName }}Store.readAll: %w", err)
		}
	}()

	// If the file doesn't exists then setup the data.
	var exists bool
	if exists, err = storage.Exists(store.uri); err != nil {
		return
	}
	if !exists {
		store.data.Records = make([]_record_.{{ .RecordName }}, 0, 1024)
		return
	}

	// Open.
	var rc fyne.URIReadCloser
	if rc, err = storage.Reader(store.uri); err != nil {
		return
	}
	defer func() {
		closeErr := rc.Close()
		if err == nil {
			err = closeErr
		}
	}()

	// Read.
	buffer := bytes.Buffer{}
	if _, err = buffer.ReadFrom(rc); err != nil {
		return
	}
	err = yaml.Unmarshal(buffer.Bytes(), &store.data)
	return
}

func (store *{{ .RecordName }}Store) writeAll() (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("{{ .RecordName }}Store.writeAll: %w", err)
		}
	}()

	// Open.
	var wc fyne.URIWriteCloser
	if wc, err = storage.Writer(store.uri); err != nil {
		return
	}
	defer func() {
		closeErr := wc.Close()
		if err == nil {
			err = closeErr
		}
	}()

	// Convert.
	var bb []byte
	if bb, err = yaml.Marshal(&store.data); err != nil {
		return
	}

	// Write.
	_, err = wc.Write(bb)
	return
}

// nextID returns the next available id.
// Returns the error if there are no more ids.
func (store *{{ .RecordName }}Store) nextID() (nextID uint64, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("{{ .RecordName }}Store.nextID: %w", err)
		}
	}()


	if store.IsFull() {
		err = Err{{ .RecordName }}Full
		return
	}

	store.data.LastID++
	nextID = store.data.LastID

	return
}

`
