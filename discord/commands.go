package secular

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"path"

	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
)

func extractFileExtension(urlStr string) string {
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		log.Println("Error parsing URL:", err)
		return ""
	}

	ext := path.Ext(parsedURL.Path)
	if ext != "" {
		return ext
	}
	return ""
}

func CommandUpload(s *discordgo.Session, i *discordgo.InteractionCreate) {
	attachmentID := i.ApplicationCommandData().Options[0].Value.(string)
	attachmentUrl := i.ApplicationCommandData().Resolved.Attachments[attachmentID].URL

	req, err := http.NewRequest("GET", attachmentUrl, nil)
	if err != nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{
					{
						Description: "Error creating request: " + err.Error(),
						Color:       0xFFD700,
					},
				},
			},
		})
		return
	}

	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br, zstd")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{
					{
						Description: "Error handling attachment: " + err.Error(),
						Color:       0xFFD700,
					},
				},
			},
		})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{
					{
						Description: "Error reading response body: " + err.Error(),
						Color:       0xFFD700,
					},
				},
			},
		})
		return
	}

	fmt.Print(len(body))

	fileExtension := extractFileExtension(attachmentUrl)
	if fileExtension == "" {
		fileExtension = ".bin" // default file extension if none found
	}

	uuid := uuid.New()
	link, err := Upload(body, uuid.String()+fileExtension)
	if err != nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{
					{
						Description: "Error uploading file: " + err.Error(),
						Color:       0xFFD700,
					},
				},
			},
		})
		return
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Description: link,
					Color:       0xFFD700,
				},
			},
		},
	})
}
