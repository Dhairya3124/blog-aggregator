# Blog Aggregator
This is a simple blog aggregator that fetches the latest posts from a list of blogs and displays them in the cli. The blogs/posts are fetched using their RSS feeds. The aggregator is written in Golang. 
Since it is created for only local use so it does not contains any server.
The first step to create a file named as ~/.gatorconfig.json on the home directory of the system. The file contains the following structure:
```json
{
  "db_url": "connection_string_goes_here",
  "current_user_name": "username_goes_here"
}
```
In the db_url field, the connection string of the database is provided so you can add the postgres database url. The current_user_name field contains the name of the user who is currently using the aggregator and no need to make any changes there as it will be defined by the application.

After creating the config file, you can install the cli application by running the following command:
```bash
go install github.com/Dhairya3124/blog-aggregator@latest
```
If you want to run the application without installing it, you can run the following command by cloning the repository and run the following command in the root directory: 
```bash
go run . <command-name>
```
List of commands that can be used with the aggregator:
- login
- register
- reset
- users
- agg
- addfeed
- feeds
- follow
- following
- unfollow
- browse

