# Naive Bayesian library for Go [![Build Status](https://travis-ci.org/goglue/bayesian.svg?branch=master)](https://travis-ci.org/goglue/bayesian) [![Coverage Status](https://coveralls.io/repos/github/goglue/bayesian/badge.svg?branch=master)](https://coveralls.io/github/goglue/bayesian?branch=master)

- Install

```
    go get github.com/goglue/bayesian
```

- Usage

```go
import "github.com/goglue/bayesian"

func main() {
        // new structure of the classifier
        classifier := bayesian.New()
        classes := []string{"laptops", "computers"}
        
        // add classes to the classifier
        classifier.AddClasses(classes)
        
        // - teach laptops class
        classifier.Learn("laptops", "wifi")
        classifier.Learn("laptops", "wifi")
        classifier.Learn("laptops", "ram")
        classifier.Learn("laptops", "hdd")
        
        // - teach computers class (LearnBulk)
        comLearnings := []string{"wifi", "monitor", "monitor", "ram", "hdd"}
        classifier.LearnBulk("computers", comLearnings)
        
        // setup is done, now get the scores of a given docs
        nodes := []string{"wifi", "ram", "hdd"}
        scores := classifier.Probability(nodes)
        /*
        scores :
        map[laptops:0.7575757575757576 computers:0.24242424242424246]
        */
}
```

- Frequency table
You can implement this interface and set the classifier storage to it

```go
type ClassFrequencyTable interface {
        Classes() []string
        AddClasses([]string) error
        IncrementNodeFrequency(cls string, node string, i uint) error
        IncrementBulkNodeFrequencies(cls string, node []string) error
        NodeFrequencyInClass(class string, node string) uint
        TotalClassNodesFrequencies() map[string]uint
        AllNodesFrequencies() uint
}

classifier.SetStorage(someRandomStorage)
```
