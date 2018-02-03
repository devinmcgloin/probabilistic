# Probabilistic Datastructures and Algorithms
[![Build Status](https://travis-ci.org/devinmcgloin/probabilistic.svg?branch=master)](https://travis-ci.org/devinmcgloin/probabilistic)

This Library implements Probabilistic Datastructures for Golang.

```
go get github.com/devinmcgloin/probabilistic
```

## Implementations

Currently includes Bloom Filters, Min Sketch Count and Min Hashing. I have plans
to add HyperLogLog++, SkipLists and Treaps. They'll arrive soonish or maybe not
at all. If you want to help out with any of these or improve the hashing
algorithms used feel free to make a pull request.

## Usage

* All of these data structures take error thresholds, if you require less
    accuracy it's important to reduce those thresholds for performance reasons.
