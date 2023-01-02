# alternativeco-challenge

## Dependencies

- `docker`/`docker-compose`. May also work with `podman` and `podman-compose`, but not tested.

## Setup

### development

Run `docker-compose up -d` to start the containers. The app will
run directly from the source code, so you can make changes and see
them immediately.

### production

Run `docker-compose -f compose.yml -f compose.prod.yml up -d`
build the binaries and the container images and run the
containers.

## Features

- [-] email service
    - [x] API endpoint
    - [x] Database connection
    - [x] RabbitMQ connection
    - [ ] Communication to ban service
- [-] ban service
    - [ ] Data structure to store the ban list (trie/rbtree or golang's map)
    - [x] RabbitMQ connection
    - [ ] Consume email events
    - [ ] Publish ban event

## TODO/Improvements

- Use versioned migrations
- Only store the email after consulting the ban list
- Use transactions on the database
- Unit tests
- Add a factory to thunder to create structs that implements `HTTPHandler`
- Add a database abstraction layer to thunder
- Document the necessary environment variables (`PORT`, `RABBITMQ_URL`)
- Manage environment vars with a config file
- Naming (`EmailService`)
- Project structure compliance
