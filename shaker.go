package main

import (
	"fmt"
	"math/rand"
	"time"
)

var magicAnswers = [...]string{
	// Positive outcomes
	"It is certain",
	"It is decidedly so",
	"Without a doubt",
	"Yes definitely",
	"You may rely on it",
	"As I see it, yes",
	"Most likely",
	"Outlook good",
	"Yes",
	"Signs point to yes",

	// Neutral outcomes
	"Reply hazy try again",
	"Ask again later",
	"Better not tell you now",
	"Cannot predict now",
	"Concentrate and ask again",

	// Negative outcomes
	"Don't count on it",
	"My reply is no",
	"My sources say no",
	"Outlook not so good",
	"Very doubtful",
}

// Shake randomize Magic 8-Ball mechanism for picking
// one answer from possible twenty answers
func Shake() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf(magicAnswers[rand.Intn(len(magicAnswers))])
}
