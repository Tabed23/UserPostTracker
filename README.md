Golang and Gin framework, using MongoDB as the database. It allows users to create, like, and unlike posts.

## Prerequisites

- Docker installed
- Docker Compose installed

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/Tabed23/UserPostTracker.git
   cd UserPostTracker
   ```

2. .env
    ```bash
    MONGODB_URL="mongodb://mongo:27017"
    ```

3. Run the API
 ```bash
  docker-compose up --build -d
```


4. Access the application:
```bash
The application is now running and accessible at http://localhost:8080
```

5. API Endpoints

**Create a user:** POST ```http://localhost:8080/v1/api/users```

**Create a post:** POST ```http://localhost:8080/v2/api/posts```

**Comment on a post:** POST ```http://localhost:8080/v2/api/posts/{post_id}/comments```

**Like a post:** PUT ```http://localhost:8080/v2/api/posts/like```

**Unlike a post:** PUT ```http://localhost:8080/v2/api/posts/unlike```

**Get all posts:** GET ```http://localhost:8080/v2/api/posts```

**Get a specific post:** ```GET http://localhost:8080/v2/api/posts/{post_id}```

**Update a post:** PUT ```http://localhost:8080/v2/api/post```

6.Postman Collection
```
A Postman collection is available in the repository for testing the API endpoints. Import the Postman collection into your Postman application to access the pre-configured requests.
```

