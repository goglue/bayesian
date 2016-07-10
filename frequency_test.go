package bayesian

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Testing the addition of classes functionality
func TestRamClassStorage_AddClasses(t *testing.T) {
	// setup
	cls := []string{"testClass1", "testClass2"}
	fs := NewFrequencyStorage()
	// execute
	fs.AddClasses(cls)
	// assert
	assert.Equal(t, 2, len(fs.storage))
	assert.NotEqual(t, nil, fs.storage[cls[0]])
	assert.NotEqual(t, nil, fs.storage[cls[1]])
}

// Testing the addition of classes functionality
func TestRamClassStorage_AddClasses_Error(t *testing.T) {
	// setup
	cls := []string{"testClass1", "testClass1", "testClass2"}
	fs := NewFrequencyStorage()
	// execute
	err := fs.AddClasses(cls)
	// assert
	assert.NotEmpty(t, err)
	assert.Equal(t, errClsAlreadyExists, err)
}

// Testing the retrieving of the added classes functionality
func TestRamClassStorage_Classes(t *testing.T) {
	// setup
	cls := []string{"testClass1", "testClass2"}
	fs := NewFrequencyStorage()
	fs.AddClasses(cls)
	// execute
	actualCls := fs.Classes()
	// assert
	assert.Contains(t, actualCls, cls[1])
	assert.Contains(t, actualCls, cls[0])
	assert.Equal(t, 2, len(actualCls))
}

// Testing the incrementing of a node within a classes functionality
func TestRamFrequencyStorage_IncrementNodeFrequency(t *testing.T) {
	// setup
	cls := []string{"testClass1", "testClass2"}
	fs := NewFrequencyStorage()
	fs.AddClasses(cls)
	// execute
	err1 := fs.IncrementNodeFrequency(cls[0], "testNode1Class1", 1)
	err2 := fs.IncrementNodeFrequency(cls[0], "testNode2Class1", 1)
	// assert
	assert.Empty(t, err1)
	assert.Empty(t, err2)
	assert.Equal(t, uint(1), fs.storage[cls[0]].freq["testNode1Class1"])
	assert.Equal(t, uint(1), fs.storage[cls[0]].freq["testNode2Class1"])
	assert.Equal(t, uint(0), fs.storage[cls[0]].freq["testNode3Class1"])
	assert.Equal(t, uint(2), fs.storage[cls[0]].total)
}

// Testing the incrementing of a node within a classes functionality
func TestRamFrequencyStorage_IncrementNodeFrequency_Error(t *testing.T) {
	// setup
	cls := []string{"testClass1", "testClass2"}
	fs := NewFrequencyStorage()
	fs.AddClasses(cls)
	// execute
	err := fs.IncrementNodeFrequency("errorClass", "testNode1Class1", 1)
	// assert
	assert.NotEmpty(t, err)
	assert.Equal(t, errClsNotExisted, err)
}

// Testing the incrementing of a node within a classes functionality
func TestRamFrequencyStorage_IncrementBulkNodeFrequency(t *testing.T) {
	// setup
	cls := []string{"testClass1", "testClass2"}
	fs := NewFrequencyStorage()
	fs.AddClasses(cls)
	nodes := []string{"testNode1Class1", "testNode2Class1"}
	// execute
	fs.IncrementBulkNodeFrequencies(cls[0], nodes)
	// assert
	assert.Equal(t, uint(1), fs.storage[cls[0]].freq["testNode1Class1"])
	assert.Equal(t, uint(1), fs.storage[cls[0]].freq["testNode2Class1"])
	assert.Equal(t, uint(0), fs.storage[cls[0]].freq["testNode3Class1"])
	assert.Equal(t, uint(2), fs.storage[cls[0]].total)
}

// Testing the incrementing of a node within a classes functionality
func TestRamFrequencyStorage_IncrementBulkNodeFrequency_Error(t *testing.T) {
	// setup
	cls := []string{"testClass1", "testClass2"}
	fs := NewFrequencyStorage()
	fs.AddClasses(cls)
	nodes := []string{"testNode1Class1", "testNode2Class1"}
	// execute
	err := fs.IncrementBulkNodeFrequencies("something", nodes)
	// assert
	assert.NotEmpty(t, err)
	assert.Equal(t, errClsNotExisted, err)
}

// Testing retrieving node count functionality
func TestRamFrequencyStorage_NodeFrequency(t *testing.T) {
	// setup
	cls := []string{"testClass1", "testClass2"}
	fs := NewFrequencyStorage()
	fs.AddClasses(cls)
	fs.IncrementNodeFrequency(cls[0], "testNode1Class1", 1)
	fs.IncrementNodeFrequency(cls[0], "testNode2Class1", 1)
	// execute
	c1 := fs.NodeFrequencyInClass(cls[0], "testNode1Class1")
	c2 := fs.NodeFrequencyInClass(cls[0], "testNode2Class1")
	// assert
	assert.Equal(t, uint(1), c1)
	assert.Equal(t, uint(1), c2)
}

// Testing the counting of total class nodes
func TestRamFrequencyStorage_TotalClassNodesFrequencies(t *testing.T) {
	// setup
	cls := []string{"testClass1", "testClass2"}
	fs := NewFrequencyStorage()
	fs.AddClasses(cls)
	fs.IncrementNodeFrequency(cls[0], "testNode1Class1", 1)
	fs.IncrementNodeFrequency(cls[0], "testNode2Class1", 1)
	fs.IncrementNodeFrequency(cls[0], "testNode3Class1", 1)
	// execute
	mapCounts := fs.TotalClassNodesFrequencies()
	// assert
	assert.Equal(t, uint(3), mapCounts[cls[0]])
	assert.Equal(t, uint(0), mapCounts[cls[1]])
}

// Testing the counting of total class nodes
func TestRamFrequencyStorage_AllNodesFrequencies(t *testing.T) {
	// setup
	cls := []string{"testClass1", "testClass2"}
	fs := NewFrequencyStorage()
	fs.AddClasses(cls)
	fs.IncrementNodeFrequency(cls[0], "testNode1Class1", 1)
	fs.IncrementNodeFrequency(cls[0], "testNode2Class1", 1)
	fs.IncrementNodeFrequency(cls[0], "testNode2Class1", 1)
	fs.IncrementNodeFrequency(cls[1], "testNode3Class2", 1)
	// execute
	count := fs.AllNodesFrequencies()
	// assert
	assert.Equal(t, uint(4), count)
}
