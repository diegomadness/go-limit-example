# Rate-limited API Challenge
## Goal
Create a Dockerized REST API that exposes a single rate-limited endpoint
/limit. No single user should be able to hit the endpoint more than N requests per second.
You can identify a user by their API key in the X-API-TOKEN in the request header.
If there is no API key, then you should immediately reject the request.

## Running the Docker image
`docker run -p 3001:3001 diegomadness/limit &`

## Discuss any design decisions, considerations, trade-offs, etc you made
Original task requests limit was 10 per second, but I've set it to 2 so you can have
easier time hitting the limit if the server is deployed far away from you.

This exact task was already completed by me in PHP+Nginx+Redis stack. I've
decided to spend less than 6 hours on this solution and make code to be as simple as
possible. What I am trying to show with this solution:
1. I can complete the task using Golang
2. Deploy process is very simple
3. I use Dockerfile for app deployment

## Describe how you would make the server production-ready:
### Deployment: How would you deploy this server?
1. `git clone https://github.com/diegomadness/limit`
2. `cd limit`
3. `docker build -t diegomadness/limit .`
4. `docker run -p 3001:3001 diegomadness/limit &`

You can repeat step 4 as much times as you have ports.
To stop this nonsense see `docker ps` for running container ID and kill it
typing `docker stop %container-id%`

### Scalability: How would you scale this server to many concurrent users?
Assuming this service is going to be a part of Kubernetes deployment, Kubernetes
load balancer can take care of this challenge. To improve the app performance,
`fasthttp` package can be used instead of `net/http` to provide the server
functionality.

I would also consider replacing `go-cache` with standalone Redis or Memcached
deployment to increase stability and performance.

### Monitoring: How would you monitor the server (For example, request rate per API key)?
Any requests over the limit can easily be tracked by service like Sentry. I can also have 
Prometheus connecting to the application every now and then to get a snapshot of cached 
requests statistics. 

### Test results:
![Test result](https://i.ibb.co/5YZXgQW/example.png)


