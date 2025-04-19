# gator
boot.dev guided project.

## Requirements:
This project requires the use of PostgreSQL and Go installed to run.

## Installation:
gator project should be installed using the `go install github.com/BlochLior/gator`
before trying to run the program, you will need to add at your home_directory a file named .gatorconfig.json, entering your postgresql connection string, the username will be automatically added when you register and login to gator.
```
{"db_url":"<postgresqlConnectionString>","current_user_name":"username_will_go_here"}
```

and with this behind you, you have successfuly installed gator!

## Commands supported:
The following commands are currently supported by the gator tool:
### register:
register a new user to the database. lastly registered user is set to the logged in user.
this register feature does not hold a password, it's just a convenience, you can set these users as folders for your rss feeds to contain. 

`gator register userName`

### login:
log in to a user held within the database. a user first needs to be registered for it to work.
this command is used can be used to follow or unfollow specific feeds in a specific user's feed.
`gator login userName`

### users:
retrieve a list of all users contained in the database.
`gator users`

### agg (WARNING):
use this command with care, so as not to DOS the servers of your followed blogs and feeds. 

### addfeed:

### feeds:

### follow:

### following:

### unfollow:

### browse:

### reset (WARNING):

## possible further extentions down the line:
Add sorting and filtering options to the browse command
Add pagination to the browse command
Add concurrency to the agg command so that it can fetch more frequently
Add a search command that allows for fuzzy searching of posts
Add bookmarking or liking posts
Add a TUI that allows you to select a post in the terminal and view it in a more readable format (either in the terminal or open in a browser)
Add an HTTP API (and authentication/authorization) that allows other users to interact with the service remotely
Write a service manager that keeps the agg command running in the background and restarts it if it crashes