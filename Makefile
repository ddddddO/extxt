run:
	GOOGLE_APPLICATION_CREDENTIALS=/mnt/c/Users/lbfde/Downloads/tag-mng-b8e1b87744fc.json \
	go run cmd/extxt/main.go -i ./testdata/image.JPG -o testdata/output/x.json

runsrv:
	GOOGLE_APPLICATION_CREDENTIALS=/mnt/c/Users/lbfde/Downloads/tag-mng-b8e1b87744fc.json \
	go run cmd/extxt/main.go server

# release:
# 	git commit -m 'msg' && \
# 	git tag -a v1.0.X -m 'msg' && \
# 	git push origin v1.0.X
#	git push