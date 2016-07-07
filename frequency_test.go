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

// Testing the retrieving of the added classes functionality
func TestRamClassStorage_Classes(t *testing.T) {
	// setup
	cls := []string{"testClass1", "testClass2"}
	fs := NewFrequencyStorage()
	fs.AddClasses(cls)
	// execute
	actualCls := fs.Classes()
	// assert
	assert.Equal(t, cls, actualCls)
}

// Testing the incrementing of a node within a classes functionality
func TestRamFrequencyStorage_IncrementNodeFrequency(t *testing.T) {
	// setup
	cls := []string{"testClass1", "testClass2"}
	fs := NewFrequencyStorage()
	fs.AddClasses(cls)
	// execute
	fs.IncrementNodeFrequency(cls[0], "testNode1Class1", 1)
	fs.IncrementNodeFrequency(cls[0], "testNode2Class1", 1)
	// assert

	// appended nodes
	assert.Equal(t, uint(1), fs.storage[cls[0]].freq["testNode1Class1"])
	assert.Equal(t, uint(1), fs.storage[cls[0]].freq["testNode2Class1"])
	// unusual cases
	assert.Equal(t, uint(0), fs.storage[cls[0]].freq["testNode3Class1"])
	// total
	assert.Equal(t, uint(2), fs.storage[cls[0]].total)
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
