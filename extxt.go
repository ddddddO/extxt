package extxt

import (
	"context"
	"fmt"
	"io"
	"os"

	vision "cloud.google.com/go/vision/apiv1"
	"google.golang.org/api/option"
)

// Client is ...
type client struct {
	*vision.ImageAnnotatorClient
}

func newClient(ctx context.Context, opt option.ClientOption) (*client, error) {
	c, err := vision.NewImageAnnotatorClient(ctx, opt)
	if err != nil {
		return nil, err
	}

	return &client{
		c,
	}, nil
}

// Run is ...
func Run(w io.Writer, targetFile, serviceAccountFile string) error {
	ctx := context.Background()
	opt := option.WithCredentialsFile(serviceAccountFile)
	cli, err := newClient(ctx, opt)
	if err != nil {
		return err
	}

	if err := cli.detectText(ctx, w, targetFile); err != nil {
		return err
	}

	return nil
}

// detectText gets text from the Vision API for an image at the given file path.
func (c *client) detectText(ctx context.Context, w io.Writer, targetFile string) error {
	f, err := os.Open(targetFile)
	if err != nil {
		return err
	}
	defer f.Close()

	image, err := vision.NewImageFromReader(f)
	if err != nil {
		return err
	}
	annotations, err := c.DetectTexts(ctx, image, nil, 10)
	if err != nil {
		return err
	}

	if len(annotations) == 0 {
		fmt.Fprintln(w, "No text found.")
	} else {
		fmt.Fprintln(w, "Text:")
		for _, annotation := range annotations {
			fmt.Fprintf(w, "%q\n", annotation.Description)
		}
	}

	return nil
}
