# ⚙️ Mini Grader Backend
This one is for practicing Go. The structure will mainly be golang, GQL, and MySQL

## ✅ Dependencies
* [**goose**](https://github.com/pressly/goose/tree/master) for migration. For arch users, you can install `goose` from AUR.
* **make** for some migration commands. No need to install if you want to type the DB string manually 🙃

## 📂 Structure
* `internal`: Main infrastructure of the backend
    * `internal/controllers`: The part that communicates with the API and sends the processed input to the services.
    * `internal/services`: This part takes data from controllers to process them. Models can only be called from this layer.
    * `internal/models`: The closest part to the database itself, shouldn't be called in the outer layer (we'll see how it goes)
* `graph`: Related to GraphQL, i.e. the API layer.
* `database`: Database setup and migrations.
* `utils`: Other utility functions that is common to the entire backend.

## ⏩ Migrations
* In the topmost level of this repo, you can run the following command to migrate the database.
```
make migrate_up
```
* And this is for reroll-ing the migration by one step. Other than that you probably need to write the command manually (managed by goose).
```
make migrate_down
```
* To create migration, use this command. Note the the `-dir` should points to the migrations folder (or things can get messy).
```
goose -dir database/migrations create <migration_name> sql
```
