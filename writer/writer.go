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


func LogPhotoId(photoID *[]tgbotapi.PhotoSize) {
	bestIndex := -1
	logger.Logger.Debugf("Finding the best photo from struct:\n%v", *photoID)
	for k,v := range *photoID {
		logger.Logger.Debugf("Photo index is <%d>", k)
		logger.Logger.Debugf("Photo size is <%d>", v.FileSize)
		logger.Logger.Debugf("Photo ID is <%s>", v.FileID)
		if bestIndex == -1 {
			logger.Logger.Debug("No bestIndex found, assigning")
			bestIndex = k
		} else {
			if v.FileSize > (*photoID)[bestIndex].FileSize {
				logger.Logger.Debugf("This photo is bigger than <%d> , assigning", (*photoID)[bestIndex].FileSize)
				bestIndex = k
			}
		}
	}
	if bestIndex != -1 {
		logger.Logger.Debugf("Best photo ID is <%s>", (*photoID)[bestIndex].FileID)
		logger.Logger.Debugf("Last photo ID is <%s>", (*photoID)[len(*photoID) - 1].FileID)
	} else {
		logger.Logger.Debug("No bestIndex found. Empty photo list?")
	}
}
