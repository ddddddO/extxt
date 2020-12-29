package extxt

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	vision "cloud.google.com/go/vision/apiv1"
)

type client struct {
	*vision.ImageAnnotatorClient
}

func newClient(ctx context.Context) (*client, error) {
	c, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		return nil, err
	}

	return &client{
		c,
	}, nil
}

// Run is ...
func Run(w io.Writer, targetFile string) error {
	ctx := context.Background()
	cli, err := newClient(ctx)
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
		fmt.Fprint(w, `{"text":`)
		for i, annotation := range annotations {
			s := annotation.Description
			if i == 0 {
				s = strings.ReplaceAll(s, "\n", "")
				fmt.Fprintf(w, `%q,"words":[`, s)
				continue
			}
			if i == len(annotations)-1 {
				fmt.Fprintf(w, "%q]", s)
				continue
			}

			fmt.Fprintf(w, "%q,", s)
		}
		fmt.Fprintln(w, "}")
	}

	return nil
}
