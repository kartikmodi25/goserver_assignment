This applications servers 5 kind of operations

Base URL - "http://localhost:8080"

1. Register User (POST): "/users"
Data Format (email, name, dob)
2. ListUsers (GET): "/users"

3. AddMovie (POST): "/movies"
Data Format (user_id, title)

4. DeleteMovie (DELETE): "/movies/:id"
 
5. ListMoviesForUser (GET): "/movies/:user_id"