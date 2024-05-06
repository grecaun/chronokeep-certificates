package types

import (
	"chronokeep/certificates/util"
	"context"
	"fmt"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
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

func (c Certificate) GenerateCertificate(config *util.Config) ([]byte, error) {
	log.WithFields(log.Fields{
		"name":  c.Name,
		"event": c.Event,
		"time":  c.Time,
		"date":  c.Date,
	}).Debug("Creating certificate.")
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	var buf []byte
	if err := chromedp.Run(
		ctx,
		chromedp.Tasks{
			chromedp.Navigate("about:blank"),
			chromedp.ActionFunc(
				func(ctx context.Context) error {
					frameTree, err := page.GetFrameTree().Do(ctx)
					if err != nil {
						return err
					}
					return page.SetDocumentContent(frameTree.Frame.ID, GetCertificateHTML(c.Name, c.Event, c.Time, c.Date, config.CertificateImage)).Do(ctx)
				},
			),
			chromedp.FullScreenshot(&buf, 90),
		}); err != nil {
		return nil, err
	}
	return buf, nil
}

func GetCertificateHTML(name string, event string, time string, date string, certImg string) string {
	return fmt.Sprintf(
		"<html>"+
			"<head></head>"+
			"<body style='width:800;height:565;padding:0px;background-image:url(\"data:image/png;base64,%s\");background-size:cover;'>"+
			"<div style='margin:0px;width:800px;height:565px;position:relative;'>"+
			"<div style='width:100%%;margin:0;position:absolute;top:50%%;-ms-transform:translateY(-50%%);transform:translateY(-50%%);'>"+
			"<div style='font-size:60px;text-align:center;font-weight:bold;'>%s</div>"+
			"<div style='font-size:30px;text-align:center;margin-left:100px;width:600px;'>finished the %s with a time of</div>"+
			"<div style='font-size:60px;text-align:center;font-weight:bold;'>%s</div>"+
			"<div style='font-size:30px;text-align:center;'>on this day of %s</div>"+
			"</div>"+
			"</div>"+
			"</body>"+
			"</html>",
		certImg,
		name,
		event,
		time,
		date,
	)
}
