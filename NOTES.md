# Notes

## Assumptions
- The order of updates is not important
- The user of the tool are operators that have a toolset to accomplish their job
- The application outputs results but we assume that the webservice also has monitoring and will be the one raising alerts
- It is a "manual operation" but in its form, it could be turned into a CRON job

## Design choices
- It is a CLI because I feel it is easier and lighter for its user
- The language used is Go because I feel it is a good and efficient languages for web services and CLI apps.
- The update process relies on a worker pool or thread pool because:
    - It can be pretty long if we update players one by one
    - We can take advantage of our multiple CPUs
    - It is a good middle point, in my opinion, between doing things in batches and one by one.
- The Go packages structure follows a thought process that I really like described here: https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1

## What can be done better
- There are unit tests but I tend to favor integration and end-to-end tests. I believe this can be done with docker-compose where a fake service would have been started mimicking the update API and the cli would be in another container and call that API and from that we could tail and analyze the logs.
- The way it is build, the cli could have many subcommands to accomplish other operations.
- The build instructions assume that the machine running it has go and make installed. We could reduce that to only needing docker and build inside of a controller docker container and just map the output to the outside.
- How it is distributed: I enjoy it when this type of CLI are self-aware of their version and can update themselves.
