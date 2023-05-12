package main

// We take in account words until n-depth before the word analyzed, minimum is of course 1
const depth int = 15

// Number of words to generate
const nbWords int = 25

// Random pick in best fitting words array possibilities
const randPickOverBest int = 3

// Influences for score attribution
const occurInfluence float64 = 0.5
const squareDistanceFromWordInfluence float64 = 0.3
const wordLengthInfluence float64 = 5.0

// Minim occurrences of a word followed by another to take take it in account for generation
const minOccurToAccept int = 1
