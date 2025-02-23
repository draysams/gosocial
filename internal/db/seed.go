package db

import (
	"context"
	"log"
	"strconv"
	"strings"

	"github.com/draysams/gosocial/internal/store"
)

var usernames = []string{
	"GhanaGuy123",
	"AccraAce",
	"KumasiKid",
	"AkuaAfrica",
	"BlackStarBoy",
	"GoldCoastGirl",
	"Djanii_D",
	"GhanaianGal",
	"AshantiAngel",
	"TamaleTough",
	"Kofi_Kal",
	"AdomChampion",
	"KofiKing",
	"AsanteYaa",
	"TemaTiger",
	"BuiBoss",
	"GhanaGirl20",
	"CapeCoastCraze",
	"NanaNtiamoah",
	"SankofaStar",
}

var titles = []string{
	"Sunset at the Beach",
	"New Project Announcement",
	"Cooking Adventure",
	"Gratitude for Support",
	"Hiking in the Mountains",
	"Incredible Book Recommendation",
	"Special Day Celebration",
	"Fitness Goals Update",
	"Quality Time with Pets",
	"Inspiring Conference Experience",
	"Exploring a New City",
	"Spa Day Relaxation",
	"Art Purchase from Local Artist",
	"Productive Day at Work",
	"Stunning Sunrise Experience",
	"New Chapter in Life",
	"Cozy Movie Night",
	"Challenging Puzzle Completed",
	"Spontaneous Road Trip",
	"Gratitude for the Little Things",
}

var contents = []string{
	"Just enjoyed a beautiful sunset at the beach!",
	"Excited to announce my new project! Stay tuned for more updates.",
	"Cooked up a storm in the kitchen tonight. Check out my latest dish!",
	"Feeling grateful for all the love and support. Thank you all! ❤️",
	"Had an amazing day hiking in the mountains. Nature is so beautiful!",
	"Just finished reading an incredible book. Highly recommend it!",
	"Celebrating a special day with family and friends. Couldn't be happier!",
	"Working hard on my fitness goals. No pain, no gain!",
	"Enjoying some quality time with my furry friend.",
	"Attended an inspiring conference today. Learned so much!",
	"Exploring a new city and loving every minute of it!",
	"Spent the day relaxing at the spa. Feeling rejuvenated!",
	"Just bought some amazing art from a local artist. So talented!",
	"Had a productive day at work. Feeling accomplished!",
	"Caught the most stunning sunrise this morning. Worth waking up early for!",
	"Excited to start a new chapter in my life. Wish me luck!",
	"Enjoying a cozy night in with a good movie. Perfect evening!",
	"Just finished a challenging puzzle. Feeling proud!",
	"Took a spontaneous road trip and discovered some hidden gems.",
	"Feeling grateful for the little things in life. Every day is a blessing!",
}

var tags = [][]string{
	{"#BeachLife", "#Nature", "#Sunset"},
	{"#NewProject", "#Announcement", "#Work"},
	{"#Foodie", "#Cooking", "#Dinner"},
	{"#Grateful", "#Support", "#Love"},
	{"#Hiking", "#Mountains", "#Nature"},
	{"#BookLovers", "#Reading", "#Recommendation"},
	{"#Celebration", "#SpecialDay", "#Family"},
	{"#FitnessJourney", "#Workout", "#Goals"},
	{"#PetLove", "#QualityTime", "#FurryFriend"},
	{"#Inspiration", "#Conference", "#Learning"},
	{"#Travel", "#NewCity", "#Exploration"},
	{"#SelfCare", "#Spa", "#Relaxation"},
	{"#ArtLovers", "#LocalArtist", "#Talent"},
	{"#WorkGoals", "#ProductiveDay", "#Accomplishment"},
	{"#Sunrise", "#Morning", "#Nature"},
	{"#NewBeginnings", "#Chapter", "#Life"},
	{"#MovieNight", "#Cozy", "#Relaxation"},
	{"#PuzzleLovers", "#Challenge", "#Proud"},
	{"#RoadTrip", "#Spontaneous", "#Gems"},
	{"#Gratitude", "#LittleThings", "#Blessing"},
}

var commentSlice = []string{
	"Wow, that sounds amazing! 😍",
	"Can't wait to hear more about it!",
	"Looks delicious! 😋",
	"Congratulations! 🎉",
	"Nature is so beautiful, isn't it?",
	"Thanks for the recommendation! 📚",
	"Enjoy every moment! ❤️",
	"You got this! 💪",
	"Adorable! ",
	"Learning is the best!",
	"Such an inspiring experience!",
	"Sounds like a great adventure!",
	"Relaxation is key! 🧖‍♀️",
	"Supporting local artists is so important!",
	"Keep up the good work! 👏",
	"Sunrises are the best! 🌅",
	"Good luck on your new journey!",
	"Cozy nights are the best! 🌙",
}

func Seed(store store.Storage) error {
	ctx := context.Background()

	users := generateUsers(100)
	for _, user := range users {
		if err := store.Users.Create(ctx, &user); err != nil {
			log.Println("Error creating user", err)
			return err
		}
	}

	posts := generatePosts(100, users)
	for _, post := range posts {
		if err := store.Posts.Create(ctx, &post); err != nil {
			log.Println("Error creating post", err)
			return err
		}
	}

	comments := generateComments(300, posts, users)
	for _, comment := range comments {
		if err := store.Comments.Create(ctx, &comment); err != nil {
			log.Println("Error creating comment", err)
			return err
		}
	}

	log.Println("Successfully seeded database")
	return nil

}

func generateUsers(n int) []store.User {
	users := make([]store.User, n)

	for i := 0; i < n; i++ {
		users[i] = store.User{
			Username: usernames[i%len(usernames)] + strconv.Itoa(i),
			Email:    strings.ToLower(usernames[i%len(usernames)]) + strconv.Itoa(i) + "@gmail.com",
			Password: "***p4ssw0rd***",
			RoleID:   1,
		}
	}

	return users
}

func generatePosts(n int, users []store.User) []store.Post {
	posts := make([]store.Post, n)

	for i := 0; i < n; i++ {
		// userId starts from 2 and is incremented by 1

		posts[i] = store.Post{
			Title:   titles[i%len(titles)],
			Content: contents[i%len(contents)],
			Tags:    tags[i%len(tags)],
			UserID:  int64(i + 1),
		}
	}

	return posts
}

func generateComments(n int, posts []store.Post, users []store.User) []store.Comment {
	comments := make([]store.Comment, n)

	for i := 0; i < n; i++ {
		comments[i] = store.Comment{
			Content: commentSlice[i%len(commentSlice)],
			PostID:  posts[i%len(posts)].ID,
			UserID:  users[i%len(users)].ID,
		}
	}

	return comments
}
