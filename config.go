package main

// We take in account words until n-depth before the word analyzed, minimum is of course 1
var depth int = 50

// Number of words to generate
var nbWords int = 50

// Random pick in best fitting words array possibilities
var randPickOverBest int = 1

// Influences for score attribution
var coefOccur float64 = 100.5
var coefDistance float64 = 20.0
var coefWordLength float64 = 100.0

// Minim occurrences of a word followed by another to take take it in account for generation
var minOccurToAccept int = 1
