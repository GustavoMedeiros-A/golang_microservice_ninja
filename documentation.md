# Download go

# Setup ENV

    - "$env:Path = [System.Environment]::GetEnvironmentVariable("Path","Machine")"
    - Make this to the operating system recognized the GO in you machine

# Write the start_line of code

# Go-chi

    - Download go-chi (is a lightweight router for build GO Http services)
    - go get -u github.com/go-chi/chi/v5
    - go.mod is add the repor in the file

# go mod tiny

    - tell go to search for any packages within our project

# Can test the first line of code with curl, point to the port


After create and config the first line of code, create a application folder and add the app and the router configurations.

And change the main to get the app to run the project


# Now, we'll install redis to use as database

    - go get github.com/redis/go-redis/v9 (to install redis)
    - docker run -p 6379:6379 redis:latest (run in docker)


# Install uuid in the project to increment the values in model
    - go get github.com/google/uuid