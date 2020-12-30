package extxt

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	vision "cloud.google.com/go/vision/apiv1"
	pb "google.golang.org/genproto/googleapis/cloud/vision/v1"
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

	annotations, err := cli.detectText(ctx, targetFile)
	if err != nil {
		return err
	}

	if len(annotations) == 0 {
		fmt.Fprintln(w, "No text found.")
		return nil
	}

	r, err := genJSONReader(annotations)
	if err != nil {
		return err
	}

	io.Copy(w, r)
	return nil
}

// detectText gets text from the Vision API for an image at the given file path.
func (c *client) detectText(ctx context.Context, targetFile string) ([]*pb.EntityAnnotation, error) {
	f, err := os.Open(targetFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	image, err := vision.NewImageFromReader(f)
	if err != nil {
		return nil, err
	}
	annotations, err := c.DetectTexts(ctx, image, nil, 10)
	if err != nil {
		return nil, err
	}

	return annotations, nil
}

func genJSONReader(annotations []*pb.EntityAnnotation) (io.Reader, error) {
	jsonWriter := &strings.Builder{}
	for i, annotation := range annotations {
		s := annotation.Description
		if i == 0 {
			s = strings.ReplaceAll(s, "\n", "")
			tmp := fmt.Sprintf(`{"text":%q,"words":[`, s)
			if _, err := jsonWriter.WriteString(tmp); err != nil {
				return nil, err
			}
			continue
		}
		if i == len(annotations)-1 {
			tmp := fmt.Sprintf("%q]", s)
			if _, err := jsonWriter.WriteString(tmp); err != nil {
				return nil, err
			}
			continue
		}

		tmp := fmt.Sprintf("%q,", s)
		if _, err := jsonWriter.WriteString(tmp); err != nil {
			return nil, err
		}
	}
	if _, err := jsonWriter.WriteString("}"); err != nil {
		return nil, err
	}

	return strings.NewReader(jsonWriter.String()), nil
}
