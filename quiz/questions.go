package quiz

// Questions dictates what can or cannot be in a question.
type Questions []struct {
	Question        string
	Answer          string
	PossibleAnswers []string
}

var quizQuestions Questions = Questions{
	{
		Question: "What is roleplay?",
		Answer:   "Acting in the capacity of a character, portrayed as a real person.",
		PossibleAnswers: []string{
			"Acting in the capacity of a character, portrayed as a real person.",
			"Forcing actions upon a player.",
			"Beating a player up and rolling them up into a sushi roll.",
		},
	},
	{
		Question: "What is metagaming?",
		Answer:   "Using OOC information in IC situations to provide you with an OOC advantage.",
		PossibleAnswers: []string{
			"Using OOC information in IC situations to provide you with an OOC advantage.",
			"Jumping repeatedly to a location.",
			"Climbing the tree, because that's a new meta.",
		},
	},
	{
		Question: "What does OOC and IC stand for?",
		Answer:   "Out of Character, and In Character.",
		PossibleAnswers: []string{
			"Out of Character, and In Character.",
			"Out of Car and In Car.",
			"Something related to Law Enforcement Officers.",
		},
	},
	{
		Question: "Are you allowed to rush taze?",
		Answer:   "No",
		PossibleAnswers: []string{
			"No.",
			"Yes.",
			"Only in certain cases.",
		},
	},
	{
		Question: "You see a mapping bug— An ideal spot for hiding from cops. What do you do?",
		Answer:   "Take a screenshot of it and post a bug report on the forums.",
		PossibleAnswers: []string{
			"Take a screenshot of it and post a bug report on the forums",
			"Hide— no one will find me!",
			"Take a screenshot of it and show it to your buddies.",
		},
	},
	{
		Question: "An admin bans you and you believe the ban is unjustified. What should you do?",
		Answer:   "Post a ban appeal on the forums, explaining the situation fully.",
		PossibleAnswers: []string{
			"Post a ban appeal on the forums, explaining the situation fully.",
			"Curse the admin over /b.",
			"Take a screenshot and immediately approach the Head of Staff.",
		},
	},
	{
		Question: "Which of the valid actions are considered valid?",
		Answer:   "/me taps on John's shoulder, trying to get his attention.",
		PossibleAnswers: []string{
			"/me taps on John's shoulder, trying to get his attention",
			"/me climbs a tree and suddenly starts flying.",
			"/do I'd punch the building and it would topple over.",
		},
	},
	{
		Question: "Are you allowed to perform illegal activites in safe zones?",
		Answer:   "Only sometimes.",
		PossibleAnswers: []string{
			"Only sometimes.",
			"No.",
			"Yes.",
		},
	},
	{
		Question: "Deathmatching is defined as:",
		Answer:   "Killing or attempting to kill a player without a valid roleplay reason to do so.",
		PossibleAnswers: []string{
			"Killing or attempting to kill a player without a valid roleplay reason to do so.",
			"Running from a cop— zig zagging.",
			"Shooting the moon, because who likes to be blinded by it?",
		},
	},
	{
		Question: "You get stopped by a Law Enforcement Officer. You decide to elude rather than comply— but you don't get far. How should your character compose themselves?",
		Answer:   "Roleplay the situation accordingly.",
		PossibleAnswers: []string{
			"Roleplay the situation accordingly.",
			"Roleplay taking out a Deagle and shooting the Officer.",
			"Just jump around— who cares?!",
		},
	},
	{
		Question: "When are you allowed to insult a fellow player?",
		Answer:   "Whenever you can justify doing it ICly.",
		PossibleAnswers: []string{
			"Whenever you can justify doing it ICly.",
			"Never.",
			"Whenever I want, wherever I want. I control these lands.",
		},
	},
	{
		Question: "Which sentence is gramatically correct?",
		Answer:   "John, you're such a lad.",
		PossibleAnswers: []string{
			"John, you're such a lad.",
			"John your such a lad.",
			"John you're sucher lad.",
		},
	},
	{
		Question: "You crash your car into another player's car— how should you compose yourself?",
		Answer:   "Roleplay the car crash accordingly.",
		PossibleAnswers: []string{
			"Roleplay the car crash accordingly.",
			"Run away! f——k them, right?",
			"Get out and shoot them for violating your rights.",
		},
	},
	{
		Question: "Revenge Killing is defined as:",
		Answer:   "Killing your killer because they killed you, just to get revenge on them.",
		PossibleAnswers: []string{
			"Killing your killer because they killed you, just to get revenge on them.",
			"Making popcorn.",
			"Killing a player to avenge a friend of yours.",
		},
	},
	{
		Question: "Is disgusting roleplay allowed?",
		Answer:   "If both parties agree to it, yes.",
		PossibleAnswers: []string{
			"If both parties agree to it, yes.",
			"Not at all.",
			"Obviously. I can perform disgusting roleplay whenever I wish to.",
		},
	},
}

// GetQuestions returns all quiz questions.
func GetQuestions() Questions {
	return quizQuestions
}
