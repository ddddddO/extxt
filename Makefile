run:
	GOOGLE_APPLICATION_CREDENTIALS=/mnt/c/Users/lbfde/Downloads/tag-mng-b8e1b87744fc.json \
	go run cmd/extxt/main.go -i ./testdata/image.JPG -o testdata/output/x.json

runsrv:
	GOOGLE_APPLICATION_CREDENTIALS=/mnt/c/Users/lbfde/Downloads/tag-mng-b8e1b87744fc.json \
	BASIC_AUTH_NAMES=user1,user2 \
	BASIC_AUTH_PASSWORDS=pass1,pass2 \
	go run cmd/extxt/main.go server

# 自動化しないこと
test:
	GOOGLE_APPLICATION_CREDENTIALS=/mnt/c/Users/lbfde/Downloads/tag-mng-b8e1b87744fc.json \
	BASIC_AUTH_NAMES="aaa,xxx" BASIC_AUTH_PASSWORDS="pass,xxxpass" \
	go test -coverprofile=./testdata/output/cover.out ./... && \
	go tool cover -html=./testdata/output/cover.out -o ./testdata/output/cover.html

# release:
# 	git commit -m 'msg' && \
# 	git tag -a v1.0.X -m 'msg' && \
# 	git push origin v1.0.X
#	git push

## NOTE: extxtに変更があった場合は、make buildappでイメージを更新&GCRへpushする。で、cloud runをdestroy -> applyする
buildapp:
	docker build -t gcr.io/extxt-300211/extxt --no-cache=true -f deployment/dockerfile/Dockerfile . && \
	docker push gcr.io/extxt-300211/extxt