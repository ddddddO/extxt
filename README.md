# Description
Implemented Google Cloud Visoin API command with go.
Extracts text from the specified image file.

# Install

```
go get github.com/ddddddO/extxt/cmd/extxt
```

# Usage
## As a command
1. Install `extxt` command.
1. Create a service account for cloud vison.
1. `export GOOGLE_APPLICATION_CREDENTIALS=/path/to/service_account_key_json`
1. Please execute `extxt -i /path/to/local/src -o /path/to/dest`. Or `-i http://path/to/src` or `-i gs://path/to/src`.
