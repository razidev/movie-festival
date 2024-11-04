### Preparation
### Database
1. In the root project, please run the file `database.sql` to your database application
2. If successful, there are 1 Schema which is `movie-festival` and four tables namely `movie, genres, user_votes, user`

### Postman
1. In the root project, please import json file `Movie Collection` to your postman
2. If Successful, there is collection named `LionParcel - Movie Collection` and the request in the collection

---

### API Explanation
### For Admin
In this APIs, the admin is responsible for Insert Movie, Update Movie, Get Highest Movie Vote and Get Highest Movie View.

1. POST /movie
```javascript
This API is responsible for inserting or uploading movies that will be view or voted on by users. The payload :

{
    title: string, // Venom
    description: string, // Optional, your movie description
    duration: int, // 10800, in seconds
    artist_name: json, // ["daniel", "alex"]
    genre_ids: json, // [1, 3, 5] based on id's genre table
    watch_url: string // url to upload the movie
}
```

2. PUT /movie/{unique_id}
```javascript
This API is responsible for updating the existing movie data. The payload is same like `POST /movie` in number 1
```

3. GET /movie/highest-vote
```javascript
This API is responsible for get the highest vote movie based on user's vote count.
```

4. GET /movie/highest-view
```javascript
This API is responsible for get the highest view movie based on user's view count
```

### For User
In this APIs, the user have authenticated and non authenticated API requests. The user can see the list of movies, watch movies, list user vote movies, Register, Login, Vote movies, and Unvote movies

### No Authentication
1. GET /user/movies?page=1&limit=5&search
```javascript
This API is responsible to show the movies list, you can search for movies and there is pagination
```

2. PUT /user/movies/{unique_id}
```javascript
This API is responsible to watch the movies, it will increase view by movie and by genre
```

3. GET /user/votes
```javascript
This API is responsible to get user who already voted on the movie
```

### Authentication
1. POST /user/register
```javascript
This API is responsible to register a user. you cannot use the same email to register. the payload :
{
    email: string, // email format string
    password: string, // password with minumum length is 6 and max length is 20
}
```

2. POST /user/login
```javascript
This API is responsible to login a user. After login you can vote movie and unvote movie. the payload is same like register
```

3. PUT /user/movies/votes/{unique_id}
```javascript
This API is responsible to vote a movie
```

4. PUT /user/movies/unvotes/{unique_id}
```javascript
This API is responsible to unvote a movie
```