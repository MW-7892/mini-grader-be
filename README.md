# ⚙️ Mini Grader Backend
This one is for practicing Go. The structure will mainly be golang, GQL, and MySQL

## ✅ Dependencies
* **make** for migration purpose

## 📂 Structure
* `internal`: Main infrastructure of the backend
    * `internal/controllers`: The part that communicates with the API and sends the processed input to the services.
    * `internal/services`: This part takes data from controllers to process them. Models can only be called from this layer.
    * `internal/models`: The closest part to the database itself, shouldn't be called in the outer layer (we'll see how it goes)
* `graph`: Related to GraphQL, i.e. the API layer.
* `database`: Database setup and migrations.
* `utils`: Other utility functions that is common to the entire backend.

## ⏩ Migrations
* When you first use this project, or when you want to migrate after adding a new migration, make sure to run this command first
```
make migrator
```
It will create a binary in `<project_root>/migrator`. Use that to call for database migration

* (Suppose you are in the project root dir) In the topmost level of this repo, you can run the following command to migrate the database.
```
./migrator up
```
* And this is for reroll-ing the migration by one step.
```
./migrator down
```
* To create migration, use this command. Note the the `-dir` should points to the migrations folder (or things can get messy).
```
./migrator create <migration_name>
```
