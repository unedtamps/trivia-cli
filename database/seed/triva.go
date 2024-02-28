package seed

import (
	"TugasRPL/database/repository"
	"TugasRPL/util"
	"context"

	"github.com/gookit/slog"
)

type triva struct {
	title      string
	question   string
	answer     string
	dificulity int64
	status     int64
	choice     []string
}

var trivia = []triva{
	{
		title:      "Capital of Indonesia",
		question:   "What is the capital of Indonesia?",
		dificulity: repository.EASY,
		status:     repository.OK,
		answer:     "Jakarta",
		choice:     []string{"Jakarta", "Bandung", "Surabaya"},
	},
	{
		title:      "Tallest Mountain",
		question:   "What is the tallest mountain in the world?",
		dificulity: repository.EASY,
		answer:     "Mount Everest",
		status:     repository.OK,
		choice:     []string{"Mount Everest", "Mount Fuji", "Mount Kilimanjaro"},
	},
	{
		title:      "Chemical Symbol for Gold",
		question:   "What is the chemical symbol for gold?",
		dificulity: repository.MEDIUM, // MEDIUM
		status:     repository.OK,     // OK
		answer:     "Au",
		choice:     []string{"Au", "Ag", "Cu", "Pt"},
	},
	{
		title:      "Author of 'To Kill a Mockingbird'",
		question:   "Who is the author of the novel 'To Kill a Mockingbird'?",
		dificulity: 1, // EASY
		status:     1, // OK
		answer:     "Harper Lee",
		choice:     []string{"Harper Lee", "Ernest Hemingway", "J.K. Rowling"},
	},
	{
		title:      "Year of First Moon Landing",
		question:   "In which year did the first moon landing occur?",
		dificulity: 2, // MEDIUM
		status:     1, // OK
		answer:     "1969",
		choice:     []string{"1969", "1971", "1974", "1980"},
	},
	{
		title:      "Longest River in the World",
		question:   "What is the longest river in the world?",
		dificulity: repository.HARD, // HARD
		status:     1,               // OK
		answer:     "Amazon River",
		choice: []string{
			"Amazon River",
			"Nile River",
			"Yangtze River",
			"Mississippi River",
			"Yenisei River",
		},
	},
	{
		title:      "Element with Atomic Number 79",
		question:   "Which element has the atomic number 79?",
		dificulity: 2, // MEDIUM
		status:     1, // OK
		answer:     "Gold",
		choice:     []string{"Gold", "Silver", "Copper", "Platinum"},
	},
	{
		title:      "Author of '1984'",
		question:   "Who is the author of the dystopian novel '1984'?",
		dificulity: 2, // MEDIUM
		status:     1, // OK
		answer:     "George Orwell",
		choice:     []string{"George Orwell", "Aldous Huxley", "Ray Bradbury", "Philip K. Dick"},
	},
	{
		title:      "Largest Moon of Jupiter",
		question:   "Which is the largest moon of the planet Jupiter?",
		dificulity: 2, // MEDIUM
		status:     1, // OK
		answer:     "Ganymede",
		choice:     []string{"Ganymede", "Io", "Europa", "Callisto"},
	},
	{
		title:      "Founder of Microsoft",
		question:   "Who is the co-founder of Microsoft Corporation?",
		dificulity: 1, // EASY
		status:     1, // OK
		answer:     "Bill Gates",
		choice:     []string{"Bill Gates", "Steve Jobs", "Mark Zuckerberg"},
	},
	{
		title:      "Current President of the United States",
		question:   "Who is the current President of the United States?",
		dificulity: 2, // MEDIUM
		status:     1, // OK
		answer:     "Joe Biden",
		choice:     []string{"Joe Biden", "Donald Trump", "Barack Obama", "George W. Bush"},
	},
	{
		title:      "Number of Planets in our Solar System",
		question:   "How many planets are there in our solar system?",
		dificulity: repository.HARD, // HARD
		status:     1,               // OK
		answer:     "8",
		choice:     []string{"8", "9", "7", "10"},
	},
	{
		title:      "Chemical Symbol for Oxygen",
		question:   "What is the chemical symbol for oxygen?",
		dificulity: 1, // EASY
		status:     1, // OK
		answer:     "O2",
		choice:     []string{"O2", "CO2", "N2", "H2O"},
	},
	{
		title:      "First Woman to Win a Nobel Prize",
		question:   "Who was the first woman to win a Nobel Prize?",
		dificulity: 2, // MEDIUM
		status:     1, // OK
		answer:     "Marie Curie",
		choice:     []string{"Marie Curie", "Mother Teresa", "Rosalind Franklin", "Jane Addams"},
	},
	{
		title:      "Fastest Land Animal",
		question:   "What is the fastest land animal?",
		dificulity: repository.HARD, // HARD
		status:     1,               // OK
		answer:     "Cheetah",
		choice:     []string{"Cheetah", "Lion", "Gazelle", "Leopard"},
	},
	{
		title:      "Founder of SpaceX",
		question:   "Who is the founder of SpaceX?",
		dificulity: 2, // MEDIUM
		status:     1, // OK
		answer:     "Elon Musk",
		choice:     []string{"Elon Musk", "Jeff Bezos", "Richard Branson", "Mark Zuckerberg"},
	},
	{
		title:      "Elementary Particle with a Positive Charge",
		question:   "Which elementary particle carries a positive electric charge?",
		dificulity: 2,
		status:     1, // OK
		answer:     "Proton",
		choice:     []string{"Proton", "Electron", "Neutron", "Quark"},
	},
	{
		title:      "Largest Desert on Earth",
		question:   "What is the largest desert on Earth?",
		dificulity: repository.HARD, // HARD
		status:     1,               // OK
		answer:     "Sahara",
		choice:     []string{"Antarctica", "Sahara", "Arctic", "Gobi"},
	},
	{
		title:      "Composer of 'Moonlight Sonata'",
		question:   "Who composed the 'Moonlight Sonata'?",
		dificulity: 2, // MEDIUM
		status:     1, // OK
		answer:     "Ludwig van Beethoven",
		choice: []string{
			"Ludwig van Beethoven",
			"Wolfgang Amadeus Mozart",
			"Johann Sebastian Bach",
			"Franz Schubert",
		},
	},
	{
		title:      "Inventor of the World Wide Web",
		question:   "Who is credited with inventing the World Wide Web?",
		dificulity: 2, // EASY
		status:     1, // OK
		answer:     "Tim Berners-Lee",
		choice:     []string{"Tim Berners-Lee", "Bill Gates", "Steve Jobs", "Mark Zuckerberg"},
	},
	{
		title:      "Human DNA Base Pairs",
		question:   "How many base pairs are there in the human DNA?",
		dificulity: 3, // HARD
		status:     1, // OK
		answer:     "3 billion",
		choice:     []string{"1 billion", "2 billion", "3 billion", "4 billion", "5 billion"},
	},
	{
		title:      "Smallest Bone in the Human Body",
		question:   "What is the smallest bone in the human body?",
		dificulity: 3, // HARD
		status:     1, // OK
		answer:     "Stapes",
		choice:     []string{"Stapes", "Femur", "Tibia", "Radius", "Spine"},
	},
	{
		title:      "Hottest Place on Earth",
		question:   "Which place holds the record for the hottest temperature on Earth?",
		dificulity: 3, // HARD
		status:     1, // OK
		answer:     "Death Valley, USA",
		choice:     []string{"Death Valley, USA", "Kuwaits", "Iran", "Libya", "India"},
	},
	{
		title:      "Periodic Table Element 79",
		question:   "Which element is represented by atomic number 79 on the periodic table?",
		dificulity: 3, // HARD
		status:     1, // OK
		answer:     "Gold",
		choice:     []string{"Gold", "Platinum", "Mercury", "Lead", "Cuprum"},
	},
	{
		title:      "Mountains in the Himalayan Range",
		question:   "How many of the world's 14 highest peaks are located in the Himalayan range?",
		dificulity: 3, // HARD
		status:     1, // OK
		answer:     "8",
		choice:     []string{"6", "8", "10", "12", "14"},
	},
	{
		title:      "Nobel Prize Categories",
		question:   "In how many categories are Nobel Prizes awarded?",
		dificulity: 3, // HARD
		status:     1, // OK
		answer:     "6",
		choice:     []string{"4", "5", "6", "7", "3"},
	},
	{
		title:      "Earliest Recorded Pandemic",
		question:   "What is considered the earliest recorded pandemic in human history?",
		dificulity: 3, // HARD
		status:     1, // OK
		answer:     "Antonine Plague",
		choice:     []string{"Antonine Plague", "Black Death", "Spanish Flu", "HIV/AIDS", "Corona"},
	},
	{
		title:      "Deep Blue Supercomputer",
		question:   "Which supercomputer defeated Garry Kasparov in a chess match in 1997?",
		dificulity: 3, // HARD
		status:     1, // OK
		answer:     "IBM Deep Blue",
		choice:     []string{"IBM Deep Blue", "Cray XT5", "Blue Gene", "Titan", "Aurora"},
	},
	{
		title:      "Number of Earth's Oceans",
		question:   "How many oceans are there on Earth?",
		dificulity: 3, // HARD
		status:     1, // OK
		answer:     "5",
		choice:     []string{"4", "5", "6", "7", "9"},
	},
	{
		title:      "The Great Wall of China",
		question:   "How long is the total length of the Great Wall of China?",
		dificulity: 3, // HARD
		status:     1, // OK
		answer:     "13,171 miles",
		choice: []string{
			"8,850 miles",
			"10,210 miles",
			"11,399 miles",
			"13,171 miles",
			"7,564 miles",
		},
	},
	{
		title:      "Inventor of the Telephone",
		question:   "Who is credited with inventing the telephone?",
		dificulity: 2, // MEDIUM
		status:     1, // OK
		answer:     "Alexander Graham Bell",
		choice: []string{
			"Thomas Edison",
			"Nikola Tesla",
			"Alexander Graham Bell",
			"Guglielmo Marconi",
		},
	},
	{
		title:      "Currency of Japan",
		question:   "What is the currency of Japan?",
		dificulity: 2, // MEDIUM
		status:     1, // OK
		answer:     "Japanese Yen",
		choice:     []string{"Japanese Won", "Japanese Yen", "Japanese Dollar", "Japanese Euro"},
	},
	{
		title:      "Painter of 'Starry Night'",
		question:   "Who is the painter of the famous artwork 'Starry Night'?",
		dificulity: 2, // MEDIUM
		status:     1, // OK
		answer:     "Vincent van Gogh",
		choice: []string{
			"Pablo Picasso",
			"Claude Monet",
			"Leonardo da Vinci",
			"Vincent van Gogh",
		},
	},
	{
		title:      "Discovery of Penicillin",
		question:   "Who is credited with the discovery of penicillin?",
		dificulity: 2, // MEDIUM
		status:     1, // OK
		answer:     "Alexander Fleming",
		choice:     []string{"Marie Curie", "Louis Pasteur", "Alexander Fleming", "Joseph Lister"},
	},
	{
		title:      "Composer of 'The Four Seasons'",
		question:   "Who composed the classical piece 'The Four Seasons'?",
		dificulity: 2, // MEDIUM
		status:     1, // OK
		answer:     "Antonio Vivaldi",
		choice: []string{
			"Wolfgang Amadeus Mozart",
			"Johann Sebastian Bach",
			"Ludwig van Beethoven",
			"Antonio Vivaldi",
		},
	},
	{
		title:      "Planet Closest to the Sun",
		question:   "Which planet is closest to the Sun?",
		dificulity: 2, // MEDIUM
		status:     1, // OK
		answer:     "Mercury",
		choice:     []string{"Venus", "Earth", "Mars", "Mercury"},
	},
	{
		title:      "Formula for Water",
		question:   "What is the chemical formula for water?",
		dificulity: 2, // MEDIUM
		status:     1, // OK
		answer:     "H2O",
		choice:     []string{"CO2", "H2O", "O2", "N2"},
	},
	{
		title:      "Mount Rushmore Presidents",
		question:   "Which U.S. Presidents are depicted on Mount Rushmore?",
		dificulity: 2, // MEDIUM
		status:     1, // OK
		answer:     "George Washington, Thomas Jefferson, Theodore Roosevelt, Abraham Lincoln",
		choice: []string{
			"George Washington, John Adams, Thomas Jefferson, James Madison",
			"Abraham Lincoln, Theodore Roosevelt, Franklin D. Roosevelt, Woodrow Wilson",
			"George Washington, Thomas Jefferson, Theodore Roosevelt, Abraham Lincoln",
			"John F. Kennedy, Ronald Reagan, Bill Clinton, Barack Obama",
		},
	},
	{
		title:      "Great Barrier Reef Location",
		question:   "In which country is the Great Barrier Reef located?",
		dificulity: 2, // MEDIUM
		status:     1, // OK
		answer:     "Australia",
		choice:     []string{"Australia", "Brazil", "Mexico", "Thailand"},
	},
	{
		title:      "Composer of 'The Nutcracker'",
		question:   "Who composed the ballet 'The Nutcracker'?",
		dificulity: 2, // MEDIUM
		status:     1, // OK
		answer:     "Pyotr Ilyich Tchaikovsky",
		choice: []string{
			"Ludwig van Beethoven",
			"Johann Sebastian Bach",
			"Pyotr Ilyich Tchaikovsky",
			"Wolfgang Amadeus Mozart",
		},
	},
	{
		title:      "Number of Continents on Earth",
		question:   "How many continents are there on Earth?",
		dificulity: 1, // EASY
		status:     1, // OK
		answer:     "7",
		choice:     []string{"5", "6", "8"},
	},
	{
		title:      "Primary Color Combination",
		question:   "Which two colors can be combined to create green?",
		dificulity: 1, // EASY
		status:     1, // OK
		answer:     "Blue and Yellow",
		choice:     []string{"Red and Blue", "Yellow and Red", "Blue and Yellow"},
	},
	{
		title:      "Largest Ocean on Earth",
		question:   "What is the largest ocean on Earth?",
		dificulity: 1, // EASY
		status:     1, // OK
		answer:     "Pacific Ocean",
		choice:     []string{"Atlantic Ocean", "Indian Ocean", "Pacific Ocean"},
	},
	{
		title:      "Number of Players in a Soccer Team",
		question:   "How many players are there in a standard soccer team?",
		dificulity: 1, // EASY
		status:     1, // OK
		answer:     "11",
		choice:     []string{"11", "13", "15"},
	},
	{
		title:      "Current Currency of Germany",
		question:   "What is the current currency of Germany?",
		dificulity: 1, // EASY
		status:     1, // OK
		answer:     "Euro",
		choice:     []string{"Deutsche Mark", "Euro", "Swiss Franc"},
	},
	{
		title:      "Smallest Planet in our Solar System",
		question:   "Which is the smallest planet in our solar system?",
		dificulity: 1, // EASY
		status:     1, // OK
		answer:     "Mercury",
		choice:     []string{"Venus", "Mars", "Mercury"},
	},
	{
		title:      "Country with the Longest Coastline",
		question:   "Which country has the longest coastline in the world?",
		dificulity: 1, // EASY
		status:     1, // OK
		answer:     "Canada",
		choice:     []string{"United States", "China", "Canada"},
	},
	{
		title:      "Human Body's Strongest Muscle",
		question:   "Which muscle in the human body is considered the strongest?",
		dificulity: 1, // EASY
		status:     1, // OK
		answer:     "Masseter",
		choice:     []string{"Quadriceps", "Masseter", "Gluteus Maximus"},
	},
	{
		title:      "Inventor of the Light Bulb",
		question:   "Who is credited with inventing the light bulb?",
		dificulity: 1, // EASY
		status:     1, // OK
		answer:     "Thomas Edison",
		choice: []string{
			"Nikola Tesla",
			"Thomas Edison",
			"Galileo Galilei",
		},
	},
	{
		title:      "Element with Atomic Number 1",
		question:   "Which element has the atomic number 1?",
		dificulity: 1, // EASY
		status:     1, // OK
		answer:     "Hydrogen",
		choice:     []string{"Helium", "Lithium", "Hydrogen"},
	},
}

func SeedTrivia(ctx context.Context) error {

	for _, v := range trivia {
		// create user first
		password := random_name()
		hash, err := util.CreateHashPass(password)
		if err != nil {
			return err
		}
		id, err := repository.Query.CreateUser(ctx, repository.CreateUserParams{
			Name:     random_name(),
			UserName: random_name(),
			RoleID:   repository.USER,
			Email:    random_email(),
			Password: hash,
		})
		if err != nil {
			return err
		}
		// user detail
		err = repository.Query.CreateUserDetail(ctx, id)
		if err != nil {
			return err
		}

		// create trivia
		trivia_id, err := repository.Query.CreateTrivia(ctx, repository.CreateTriviaParams{
			Title:        v.title,
			Question:     v.question,
			Answer:       v.answer,
			DificulityID: v.dificulity,
		})
		if err != nil {
			return err
		}
		// trivia choice
		for _, v := range v.choice {
			err := repository.Query.CreateTriviaChoice(ctx, repository.CreateTriviaChoiceParams{
				TriviaID: trivia_id,
				Choice:   v,
			})
			if err != nil {
				return err
			}
		}
		// trivia detail
		err = repository.Query.CreateTriviaDetail(ctx, repository.CreateTriviaDetailParams{
			UserID:         id,
			TriviaID:       trivia_id,
			TriviaStatusID: v.status,
		})
		if err != nil {
			return err
		}
		slog.Infof("success migrate user_id %d and trivia %s", id, v.title)
	}
	return nil
}
