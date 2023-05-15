package main

// We take in account words until n-depth before the word analyzed, minimum is of course 1
var depth int = 15

// Number of words to generate
var nbWords int = 25

// Random pick in best fitting words array possibilities
var randPickOverBest int = 3

// Influences for score attribution
var coefOccur float64 = 0.5
var coefDistance float64 = 0.3
var coefWordLength float64 = 7.0

// Minim occurrences of a word followed by another to take take it in account for generation
var minOccurToAccept int = 2
