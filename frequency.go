package bayesian

import (
	"errors"
)

// Error for existence class entry
var errClsAlreadyExists = errors.New("class already exists")

// Error for existence class entry
var errClsNotExisted = errors.New("class does not exists")

// The storage interface for storing classes nodes frequencies
type ClassFrequencyTable interface {
	Classes() []string
	AddClasses([]string) error
	IncrementNodeFrequency(cls string, node string, i uint) error
	IncrementBulkNodeFrequencies(cls string, node []string) error
	NodeFrequencyInClass(class string, node string) uint
	TotalClassNodesFrequencies() map[string]uint
	AllNodesFrequencies() uint
}

// Encapsulates the class information
type class struct {
	freq  map[string]uint
	total uint
}

// IN-RAM class nodes frequency storage
type RamFrequencyStorage struct {
	storage map[string]*class
}

// Returns the nodes count within a given class
func (r *RamFrequencyStorage) Classes() []string {
	classes := make([]string, 0, len(r.storage))
	for k := range r.storage {
		classes = append(classes, k)
	}

	return classes
}

// Returns the nodes count within a given class
func (r *RamFrequencyStorage) AddClasses(c []string) error {
	clsLen := len(c)
	r.storage = make(map[string]*class, clsLen)

	// Iterate over storage map & add each class
	for i := 0; i < clsLen; i++ {
		// If the storage has already the class
		if nil != r.storage[c[i]] {
			return errClsAlreadyExists
		}

		freq := make(map[string]uint)
		r.storage[c[i]] = &class{
			freq,
			0,
		}
	}

	return nil
}

// Increments the node frequency
func (r *RamFrequencyStorage) IncrementNodeFrequency(cls string, node string, i uint) error {
	if nil == r.storage[cls] {
		return errClsNotExisted
	}
	r.storage[cls].freq[node] += i
	r.storage[cls].total += i

	return nil
}

// Increments the node frequency
func (r *RamFrequencyStorage) IncrementBulkNodeFrequencies(cls string, node []string) error {
	if nil == r.storage[cls] {
		return errClsNotExisted
	}
	lenNodes := len(node)
	for i := 0; i < lenNodes; i++ {
		r.storage[cls].freq[node[i]] += 1
	}
	r.storage[cls].total += uint(lenNodes)

	return nil
}

// Returns the nodes count within a given class
func (r *RamFrequencyStorage) NodeFrequencyInClass(c string, n string) uint {
	return r.storage[c].freq[n]
}

// Increments the node frequency
func (r *RamFrequencyStorage) TotalClassNodesFrequencies() map[string]uint {
	count := make(map[string]uint, len(r.storage))

	for clsName, cls := range r.storage {
		count[clsName] = cls.total
	}

	return count
}

// Increments the node frequency
func (r *RamFrequencyStorage) AllNodesFrequencies() uint {
	var sum uint = 0
	for _, cls := range r.storage {
		sum += cls.total
	}

	return sum
}

// Creates new frequency storage
func NewFrequencyStorage() *RamFrequencyStorage {
	return &RamFrequencyStorage{}
}
