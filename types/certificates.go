package types

import (
	"chronokeep/certificates/util"
	"image"
	"image/draw"
	"os"

	"github.com/golang/freetype"
	log "github.com/sirupsen/logrus"
)

type (
	Certificate struct {
		Name  string
		Event string
		Time  string
		Date  string
	}

	Label struct {
		Text     string
		FontPath string
		FontType string
		Size     float64
		YPos     int
	}
)

func (c Certificate) GenerateCertificate(config *util.Config) (*image.RGBA, error) {
	certImg := config.CertificateImage
	bgImg := image.NewRGBA(image.Rect(0, 0, certImg.Bounds().Dx(), certImg.Bounds().Dy()))
	draw.Draw(bgImg, certImg.Bounds(), certImg, image.Point{}, draw.Src)
	// add labels
	bgImg, err := addLabels(bgImg, []Label{
		{
			Text:     c.Name,
			FontPath: "",
			FontType: "luxirb.ttf",
			Size:     48,
			YPos:     300,
		},
		{
			Text:     "finished the",
			FontPath: "",
			FontType: "luxirr.ttf",
			Size:     24,
			YPos:     360,
		},
		{
			Text:     c.Event,
			FontPath: "",
			FontType: "luxirr.ttf",
			Size:     36,
			YPos:     390,
		},
		{
			Text:     "with a time of",
			FontPath: "",
			FontType: "luxirr.ttf",
			Size:     24,
			YPos:     440,
		},
		{
			Text:     c.Time,
			FontPath: "",
			FontType: "luxirb.ttf",
			Size:     48,
			YPos:     480,
		},
		{
			Text:     "on this day of",
			FontPath: "",
			FontType: "luxirr.ttf",
			Size:     24,
			YPos:     530,
		},
		{
			Text:     c.Date,
			FontPath: "",
			FontType: "luxirr.ttf",
			Size:     36,
			YPos:     560,
		},
	})
	if err != nil {
		return nil, err
	}
	return bgImg, nil
}

func addLabels(bgImg *image.RGBA, labels []Label) (*image.RGBA, error) {
	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetClip(bgImg.Bounds())
	c.SetDst(bgImg)
	c.SetSrc(image.Black)

	for _, label := range labels {
		fontBytes, err := os.ReadFile(label.FontPath + label.FontType)
		if err != nil {
			return nil, err
		}
		f, err := freetype.ParseFont(fontBytes)
		if err != nil {
			return nil, err
		}
		c.SetFont(f)
		c.SetFontSize(label.Size)

		// positioning
		pt := freetype.Pt(10, label.YPos+int(c.PointToFixed(label.Size)>>6))

		// draw label
		_, err = c.DrawString(label.Text, pt)
		if err != nil {
			log.Println(err)
			return bgImg, nil
		}
		pt.Y += c.PointToFixed(label.Size * 1.5)
	}
	return bgImg, nil
}
