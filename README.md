### Create Common Package Steps
- Create Go Project from CMD
```shell
go mod init github.com/dhanushka-gayashan/common-sturcts
```

- Add Initial files and commit to local
```shell
git init 
git add --all 
git commit -m 'init'
```

- Create empty GIT Repo with same name (Ex: common-sturcts)

- Push your changers into Repo
```shell
git remote add origin <REPO_URL>
git branch -M main
git push -u origin main
```

- Create Tag
```shell
git tag "v1.0.0"
git push --tags
```

- Add module into other projects
```shell
go get -u <GITHUB_URL>
```