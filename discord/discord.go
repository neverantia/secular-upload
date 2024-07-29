package secular

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

// Bot parameters
var (
	GuildID        = "1264214645166575676"
	BotToken       = os.Getenv("secular_token")
	RemoveCommands = false
)

var s *discordgo.Session

func init() {
	flag.Parse()

	var err error
	s, err = discordgo.New("Bot " + BotToken)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}

	// Register command handlers
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
}

var (
	commands = []*discordgo.ApplicationCommand{
		{
			Name:        "upload",
			Description: "Upload Files To Server!",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "upload_file",
					Description: "File To Upload To Server",
					Type:        discordgo.ApplicationCommandOptionAttachment,
				},
			},
		},
	}

	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"upload": CommandUpload,
	}
)

func Run() {
	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})

	err := s.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	log.Println("Adding commands...")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, v := range commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, GuildID, v)
		if err != nil {
			log.Printf("Cannot create '%v' command: %v", v.Name, err)
			continue
		}
		registeredCommands[i] = cmd
	}

	defer s.Close()

	log.Println("Press Ctrl+C to exit")

	fmt.Scanln()
}
