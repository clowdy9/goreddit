# Reddit Go Project

This library uses Reddit’s JSON API to fetch media posts from a specified subreddit or user.

## Installation

   ```bash
   require https://github.com/clowdy9/goreddit.git
   ```


## Usage

   ```bash
   import https://github.com/clowdy9/goreddit
   ```

Fetch media data for a subreddit or user:
   ```go
   package main

   import (
       "fmt"
       "log"
       "https://github.com/clowdy9/goreddit"
   )

   func main() {

   	reddit := goreddit.New(goreddit.RedditOpts{
		    Proxy: "",
	    })

       posts, err := reddit.Subreddit("youtube", "")
       if err != nil {
           log.Fatalf("Error fetching post: %v", err)
       }
   }
   ```

Search reddit:
   ```go
   package main

   import (
       "fmt"
       "log"
       "https://github.com/clowdy9/goreddit"
   )

   func main() {

   	reddit := goreddit.New(goreddit.RedditOpts{
		    Proxy: "",
	    })

       posts, err := reddit.Search("youtube", "")
       if err != nil {
           log.Fatalf("Error fetching post: %v", err)
       }
   }
   ```

Fetch info for a specific post:
   ```go
   package main

   import (
       "fmt"
       "log"
       "https://github.com/clowdy9/goreddit"
   )

   func main() {

   	reddit := goreddit.New(goreddit.RedditOpts{
		    Proxy: "",
	    })

       posts, err := reddit.GetPost("youtube", postID, false)
       if err != nil {
           log.Fatalf("Error fetching post: %v", err)
       }
   }
   ```

## License

This project is licensed under the MIT License. See the LICENSE file for details.


