# i18n-rest-example
## Run the Code 
```zsh
git clone https://github.com/Xyedo/i18n-rest-example.git
cd i18n-rest-example
go run main.go
```

## Call the API
```zsh
curl --header "Accept-Language:id;q=0.9, en-US;q=0.8, " http://127.0.0.1:3000
```

## Add New Languange Support
expample : add espanyol lang

```zsh
cd pkg/locale
touch translate.es.toml
goi18n merge active.en.toml translate.es.toml
```
inside translate.es.toml, change it from english to espanyol and add it to a pkg/locale/bundle.go
