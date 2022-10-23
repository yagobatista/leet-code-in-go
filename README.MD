# Leet code in Golang

## cli

```
echo -e $'alias lc_new_case="cd cmd/make/ && go run main.go"' >> ~/.zshrc && source ~/.zshrc
echo -e $'alias lc_cover="go test -v -coverprofile cover.out ./... && go tool cover -html cover.out -o cover.html"' >> ~/.zshrc && source ~/.zshrc
```
