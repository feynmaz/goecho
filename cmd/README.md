# cmd

* The <em>cmd</em> directory can be thought of as the "entrypoint" for the application. By separating the <em>cmd/service</em> package from the root package of the project, we are adding structure to allow for multiple applications in one project.

* For example, if you wanted to implement the scheduler portion of the application as a standalone daemon that needs to use modules defined within the same repository, you are able to do so by creating a new <em>cmd/scheduler</em> package to handle this case

* Eventually, after a product is developed, it is often found that creating microservices is useful. With this <em>cmd</em> package, we would be able to easily create different packages which will become new executable services, and still use the existing codebase structure to support this task.
