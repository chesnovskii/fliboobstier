package writer

import (
	"fmt"
	"os"

	"github.com/chesnovsky/fliboobstier/logger"
	"gopkg.in/telegram-bot-api.v4"
)

// WritePhoto writing PhotoID to file
func WritePhoto(photoID *[]tgbotapi.PhotoSize) {
	file, err := os.Create("boobs.txt")
	if err != nil {
		logger.Logger.Fatalf("Cannot create file")
	}
	defer file.Close()

	fmt.Fprintln(file, photoID)

}
