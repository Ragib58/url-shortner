# How It Works
1. The API is built with Go and Gin.
2. Redis is used as the backend database for storing short URLs.
3. The API exposes endpoints to:
   * Shorten URLs
   * Retrieve original URLs
   * Edit or delete URLs
   * Add tags to URLs


# Running the App
1. Ensure Docker and Docker Compose are installed.
2. Fix the volumes path in `docker-compose.yml` to `./data:/data` if not already set.
3. Run:
   ```
   docker-compose up --build
   ```
4. The API will be available at [http://localhost:8000](http://localhost:8000), and Redis at `localhost:6379`.

# Notes
- All data is stored in Redis, with persistence enabled via the `data` volume.
- The app is modular, with clear separation between routes, database logic, and models.