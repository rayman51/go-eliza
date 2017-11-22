// Ray Mannion
// G00340315
// Data Representation & Query project
// November 2017
package main

import (
	"fmt"
	"math/rand"
	"net/http" // imports
	"regexp"   // used
	"strings"
	"text/template"
)

var user string

// got the reponses from this site
//https://www.smallsurething.com/implementing-the-famous-eliza-chatbot-in-python/
var responses = [][]string{ // responses for chatbot
	{`my name is ([^.?!]*)[.?!]?`,
		"hi $1, how can i help",
	},
	/*
		{`will liverpool win the league`,
			"NEVER",
		},
	*/
	{`Hello([^.?!]*)[.?!]?`,
		"Hello... I'm glad you could drop by today.",
		"Hi there... how are you today?",
		"Hello, how are you feeling today?"},

	{`What ([^.?!]*)[.?!]?`,
		"Why do you ask?",
		"How would an answer to that help you?",
		"What do you think?"},

	{`How ([^.?!]*)[.?!]?`,
		"How do you suppose?",
		"Perhaps you can answer your own question.",
		"What is it you're really asking?"},

	{`Because ([^.?!]*)[.?!]?`,
		"Is that the real reason?",
		"What other reasons come to mind?",
		"Does that reason apply to anything else?",
		"If $1?, what else must be true?"},

	{`([^.?!]*)[.?!]? sorry ([^.?!]*)[.?!]?`,
		"There are many times when no apology is needed.",
		"What feelings do you have when you apologize?"},

	{`I think ([^.?!]*)[.?!]?`,
		"Do you doubt $1??",
		"Do you really think so?",
		"But you're not sure $1??"},

	{`([^.?!]*)[.?!]? friend ([^.?!]*)[.?!]?`,
		"Tell me more about your friends.",
		"When you think of a friend, what comes to mind?",
		"Why don't you tell me about a childhood friend?"},

	{`Yes`,
		"You seem quite sure.",
		"OK, but can you elaborate a bit?"},

	{`([^.?!]*)[.?!]? computer([^.?!]*)[.?!]?`,
		"Are you really talking about me?",
		"Does it seem strange to talk to a computer?",
		"How do computers make you feel?",
		"Do you feel threatened by computers?"},

	{`Is it ([^.?!]*)[.?!]?`,
		"Do you think it is $1??",
		"Perhaps it's $1? -- what do you think?",
		"If it were $1?, what would you do?",
		"It could well be that $1?."},

	{`It is ([^.?!]*)[.?!]?`,
		"You seem very certain.",
		"If I told you that it probably isn't $1?, what would you feel?"},

	{`Can you ([^\?]*)\??`,
		"What makes you think I can't $1??",
		"If I could $1?, then what?",
		"Why do you ask if I can $1??"},

	{`Can I ([^\?]*)\??`,
		"Perhaps you don't want to $1?.",
		"Do you want to be able to $1??",
		"If you could $1?, would you?"},

	{`You are ([^.?!]*)[.?!]?`,
		"Why do you think I am $1??",
		"Does it please you to think that I'm $1??",
		"Perhaps you would like me to be $1?.",
		"Perhaps you're really talking about yourself?"},

	{`I need ([^.?!]*)[.?!]?`,
		"Why do you need $1??",
		"Would it really help you to get $1??",
		"Are you sure you need $1??"},

	{`Why don\'?t you ([^\?]*)\??`,
		"Do you really think I don't $1??",
		"Perhaps eventually I will $1?.",
		"Do you really want me to $1??"},

	{`Why can\'?t I ([^\?]*)\??`,
		"Do you think you should be able to $1??",
		"If you could $1?, what would you do?",
		"I don't know -- why can't you $1??",
		"Have you really tried?"},

	{`I can\'?t ([^.?!]*)[.?!]?`,
		"How do you know you can't $1??",
		"Perhaps you could $1? if you tried.",
		"What would it take for you to $1??"},

	{`I am ([^.?!]*)[.?!]?`,
		"Did you come to me because you are $1??",
		"How long have you been $1??",
		"How do you feel about being $1??"},

	{`I\'?m ([^.?!]*)[.?!]?`,
		"How does being $1? make you feel?",
		"Do you enjoy being $1??",
		"Why do you tell me you're $1??",
		"Why do you think you're $1??"},

	{`Are you ([^\?]*)\??`,
		"Why does it matter whether I am $1??",
		"Would you prefer it if I were not $1??",
		"Perhaps you believe I am $1?.",
		"I may be $1? -- what do you think?"},

	{`You\'?re ([^.?!]*)[.?!]?`,
		"Why do you say I am $1??",
		"Why do you think I am $1??",
		"Are we talking about you, or me?"},

	{`I don\'?t ([^.?!]*)[.?!]?`,
		"Don't you really $1??",
		"Why don't you $1??",
		"Do you want to $1??"},

	{`I feel ([^.?!]*)[.?!]?`,
		"Good, tell me more about these feelings.",
		"Do you often feel $1??",
		"When do you usually feel $1??",
		"When you feel $1?, what do you do?"},

	{`I have ([^.?!]*)[.?!]?`,
		"Why do you tell me that you've $1??",
		"Have you really $1??",
		"Now that you have $1?, what will you do next?"},

	{`I would ([^.?!]*)[.?!]?`,
		"Could you explain why you would $1??",
		"Why would you $1??",
		"Who else knows that you would $1??"},

	{`Is there ([^.?!]*)[.?!]?`,
		"Do you think there is $1??",
		"It's likely that there is $1?.",
		"Would you like there to be $1??"},

	{`My ([^.?!]*)[.?!]?`,
		"I see, your $1?.",
		"Why do you say that your $1??",
		"When your $1?, how do you feel?"},

	{`You ([^.?!]*)[.?!]?`,
		"We should be discussing you, not me.",
		"Why do you say that about me?",
		"Why do you care whether I $1??"},

	{`Why ([^.?!]*)[.?!]?`,
		"Why don't you tell me the reason why $1??",
		"Why do you think $1??"},

	{`I want ([^.?!]*)[.?!]?`,
		"What would it mean to you if you got $1??",
		"Why do you want $1??",
		"What would you do if you got $1??",
		"If you got $1?, then what would you do?"},

	{`([^.?!]*)[.?!]? mother([^.?!]*)[.?!]?`,
		"Tell me more about your mother.",
		"What was your relationship with your mother like?",
		"How do you feel about your mother?",
		"How does this relate to your feelings today?",
		"Good family relations are important."},

	{`([^.?!]*)[.?!]? father([^.?!]*)[.?!]?`,
		"Tell me more about your father.",
		"How did your father make you feel?",
		"How do you feel about your father?",
		"Does your relationship with your father relate to your feelings today?",
		"Do you have trouble showing affection with your family?"},

	{`([^.?!]*)[.?!]? child([^.?!]*)[.?!]?`,
		"Did you have close friends as a child?",
		"What is your favorite childhood memory?",
		"Do you remember any dreams or nightmares from childhood?",
		"Did the other children sometimes tease you?",
		"How do you think your childhood experiences relate to your feelings today?"},

	{`([^.?!]*)[.?!]?\?`,
		"Why do you ask that?",
		"Please consider whether you can answer your own question.",
		"Perhaps the answer lies within yourself?",
		"Why don't you tell me?"},

	{`quit`,
		"Thank you for talking with me.",
		"Good-bye.",
		"Thank you, that will be $150.  Have a good day!"},

	{`([^.?!]*)[.?!]?`,
		"Please tell me more.",
		"Let's change focus a bit... Tell me about your family.",
		"Can you elaborate on that?",
		"Why do you say that $1??",
		"I see.",
		"Very interesting.",
		"$1?.",
		"I see.  And what does that tell you?",
		"How does that make you feel?",
		"How do you feel when you say that?"},
}

func templateHandler(w http.ResponseWriter, r *http.Request) { //Handle Http requests

	t, _ := template.ParseFiles("chat/eliza.html") //Parse the template File
	t.Execute(w, t)                                // Execute the Tmpl file

}

func responseFromEliza(usersAnswer string) string {

	var counter int

	for _, h := range responses {
		row := 0
		for {
			row = rand.Intn(len(h)) //Find the length of the current row, Used for random row
			if row != 0 {
				break
			}
		}
		re := regexp.MustCompile("(?i)" + responses[counter][0]) //Read the find index in the row - for the question

		if matched := re.MatchString(usersAnswer); matched { //Compare question with the users input
			return re.ReplaceAllString(usersAnswer, responses[counter][row]) //Replace the question with the answer for output

		}
		counter++ //Increment the index
	}

	//Catch all if loop Fails
	answers := []string{
		"I’m not sure what you’re trying to say. Could you explain it to me?",
		"How does that make you feel?",
		"Why do you say that?",
	}

	return answers[rand.Intn(len(answers))] //Return random answer

}

func receiveAjax(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" { //Post the users input

		Q := r.FormValue("Question") //Get the value from the input field

		fmt.Fprintf(w, responseFromEliza(Q))

	}
}

func Reflect(input string) string {
	// Split the input on word boundaries.
	boundaries := regexp.MustCompile(`\b`)
	tokens := boundaries.Split(input, -1)

	// List the reflections.
	reflections := [][]string{
		{`I`, `you`},
		{`me`, `you`},
		{`you`, `me`},
		{`my`, `your`},
		{`your`, `my`},
		{`Are`, `Am`},
	}

	// Loop through each token, reflecting it if there's a match.
	for i, token := range tokens {
		for _, reflection := range reflections {
			if matched, _ := regexp.MatchString(reflection[0], token); matched {
				tokens[i] = reflection[1]
				break
			}
		}
	}

	// Put the tokens back together.
	return strings.Join(tokens, ``)
}
func main() {

	http.Handle("/", http.FileServer(http.Dir("./"))) //Handle http request
	http.HandleFunc("/chat", templateHandler)         //handle requests for templates
	http.HandleFunc("/ajax", receiveAjax)
	http.ListenAndServe(":8080", nil) //Listen and report from port 8080
}
