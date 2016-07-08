package bayesian

const zeroProbability float64 = 0.00000000001

// Classifier structure which encapsulates the frequency table as a dependency
type Classifier struct {
	cFrequency ClassFrequencyTable
}

// Sets the classifier storage
func (c *Classifier) SetStorage(storage ClassFrequencyTable) {
	c.cFrequency = storage
}

// Teaches the given class the given node
func (c *Classifier) Learn(class string, node string) {
	err := c.cFrequency.IncrementNodeFrequency(class, node, uint(1))

	if nil != err {
		panic(err)
	}
}

// Teaches the given class the given node
func (c *Classifier) LearnBulk(class string, node []string) {
	err := c.cFrequency.IncrementBulkNodeFrequencies(class, node)

	if nil != err {
		panic(err)
	}
}

// Proxies the functionality to the frequency table
func (c *Classifier) AddClasses(classes []string) {
	err := c.cFrequency.AddClasses(classes)

	if nil != err {
		panic(err)
	}
}

// Calculate the score of all classes within the giving documents
func (c *Classifier) Probability(document []string) map[string]float64 {
	class := c.cFrequency.Classes()
	lenClasses := len(class)

	score := make(map[string]float64, lenClasses)

	totalNodes := c.cFrequency.AllNodesFrequencies()
	totalNodeClassFreq := c.cFrequency.TotalClassNodesFrequencies()

	var scoreSum float64 = 0
	for j := 0; j < lenClasses; j++ {
		totalClassNodes := totalNodeClassFreq[class[j]]
		clsPrior := float64(totalClassNodes) / float64(totalNodes)
		clsScore := clsPrior

		for k := 0; k < len(document); k++ {
			nodeFrequency := c.cFrequency.NodeFrequencyInClass(class[j], document[k])
			var nodeProbability float64

			if 0 == nodeFrequency {
				nodeProbability = zeroProbability
			} else {
				nodeProbability = float64(nodeFrequency) / float64(totalClassNodes)
			}
			clsScore *= nodeProbability
		}

		score[class[j]] = clsScore
		scoreSum += clsScore
	}

	for i := 0; i < lenClasses; i++ {
		score[class[i]] /= scoreSum
	}

	return score
}

// Factory method to initialize a new classifier
func New() *Classifier {
	return &Classifier{
		NewFrequencyStorage(),
	}
}
