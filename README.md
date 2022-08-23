SIRKA Recruitment Technical Test
-
## Deployed application
https://sirka-psql-test.herokuapp.com/

- `GET`: https://sirka-psql-test.herokuapp.com/MyApp/DisplayAllUsers
- `POST` (See problem question for request body): https://sirka-psql-test.herokuapp.com/MyApp/DisplayUser

## Questions
1. **What is your Name:**
   Fahdii Ajmalal Fikrie
2. **What is the url of the Golang version API on the internet?**
   Go 1.16 is used on this project.
3. **What is the git repo url of the Golang version?**
   https://github.com/golang/go/releases/tag/go1.16
4. **What is the gdrive url containing the screenshot of the API request/response?**
   I'm using Notion to give the result of the screenshot.
   https://bbbadi.notion.site/Sirka-Backend-Technical-Test-Writeup-607a37505ae844a1847ebbe4496135a2
5. **What do you know about clean architecture?**
   I don't know much about clean architecture beside it is an approach of writing and organizing code by respecting each layer of program.
6. **Explain what authentication method used on API Call?**
   Method used are `GET` and `POST`. `GET` is generally used for fetching data, meanwhile `POST` is mainly used for creating new data.


## Running the Program
**Using Makefile**
Fill the `.env` file on `pkg/common/env` with the following data:

```bash
PORT=:8000
DB_URL=<postgres url goes here>
```

**Docker**
Run the command below
```
docker-compose up -d
```

## Lessons Learned
1. I'm having trouble with making sure the server container run after db/postgres is ready. I already tried [using depends on & wait_for_postgres.sh](https://docs.docker.com/compose/startup-order/) but ended up having no luck with it.
2. Tried to seed the data using [this tutorial](https://onexlab-io.medium.com/docker-compose-postgres-database-seed-108297cac09a), the data was successfully injected, but the `userid` & `name` fields somehow are shown as blank string.
3. Setting up go for heroku deployment is surprisingly easy (basically drag and drop for deploy, no Procfile/database adjustment needed)
