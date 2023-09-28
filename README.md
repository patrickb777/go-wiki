# go-wiki

A simple wiki application developed to create additional programming experience on the topics covered in the go.dev tutorials for building web applications and for accessing databases in Golang.

In the spirit of cross learning the application utilises a microservices approach and can be deployed using a container orchestration service such as Docker Compose / Tilt / Kubernetes etc.

### DB Container instructions

Build the image
```docker build ./database -f Dockerfile --tag=wiki-db```

Run the container
```docker run -d -p 3306:3306 --name=wiki-db wiki-db```

Note, assumes the DB runs on 172.17.0.2.  If other containers the IP Address is the `DBCXN()`` function needs to be updated.

### Useful reosurces

* https://connelblaze.medium.com/displaying-database-table-data-on-html-table-tag-92761c07e01f
