# Go answer generator

It's a hacky small project for me to try Go.
I've had the idea to make a bible quiz app which I started a few years ago and then left behind.
To try go I had the idea to create a service that can give me answer options to any question regarding the bible which was a pain in the app.
I use GPT for this, so if you want to build the generator you have to use your own GPT API key. Have fun! If you have any questions feel free to hit me up :)

## Build from source (only option right now)

1. clone the repo
2. add .env in your root folder

```
GPT_API_TOKEN=""
GPT_PROJECT_ID=""
GPT_ORG_ID=""
```

3. run `go build`
4. execute binary `/answer_generator`
5. ask any question
6. choose an option

This program is far from perfect. It's just a fun hobby project.
