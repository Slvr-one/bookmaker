# bookmaker_app:

This app was designed as a simple web server for some minimal bookmaking, aka gambling.
It allows investors to participate in various events.
The app is connected to a MongoDB database for persistancy between events.
Built with Go, utilining MongoDB, Nginx, Docker & Compose.


(comparison of popular REST API frameworks for Go development:

Gin - A lightweight framework with fast performance and easy-to-use routing capabilities. It also has a built-in middleware system for handling common tasks like logging and error handling.

Echo - A high-performance framework with a simple and intuitive API. It has features like middleware, routing, and error handling built-in, making it easy to build robust APIs quickly.

Chi - A lightweight and fast framework that offers a lot of flexibility in terms of routing and middleware. It also has a built-in logger and support for HTTP/2.

Beego - A full-featured framework with built-in support for ORM, caching, and session management. It also has a built-in admin panel for managing your application.

Gorilla - A toolkit for building web applications, including a router, middleware, and a set of handlers for common tasks like authentication and CSRF protection. It's not a full-fledged framework, but it's highly flexible and can be used to build custom solutions.

Ultimately, the choice of framework comes down to personal preference and the specific needs of your project. Each of these frameworks has its own strengths and weaknesses, so it's important to evaluate them based on your project requirements and development style.)

## Breakdown:
This app is part of my Portfolio project, which includes
developing and implementing CI/CD workflows with Jenkins, github actions, argocd.
Observability, Logging & Monitoring with EFK, Prometheus & grafana.
Deployed & configured with Kubernetes, Helm, Ansible, Terraform.
While Implementing:
    Microservices Architecture
    Gitops
    Github Flow
    Modularity
    Automation
    

## Architecture:
![image](image.png)

## REST API REF:

| Path | Method | Description |
| :-------- | :------- | :------- | 
| `/home` | `GET` | `home page`
| `/assets` | `GET` | `images & html` |
| `/health` | `GET` | `for health check` |
| `/metrics` | `GET` | `for monitoring purposes` |
| `/LH` | `GET` | `list all available horses` |
| `/GH/{name}` | `GET` | `list a specific horse detail` |
| `/UH/{name}` | `PUT` | `update a specific horse by name` |
| `/invest/{horse}/{amount}` | `UPDATE` | `invest - gamble on a horse` |






## Tech Stack

**Client:** HTML, CSS, Bootstrap, JS, Nginx

**Server:** GoLang, MongoDB

**CI/CD:** Jenkins / Github Actions

**Cloud:** AWS / GCP

**IAC:** Terraform, Ansible


## 🔗 Links

[![linkedin](https://img.shields.io/badge/linkedin-0A66C2?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/dvir-gross-929252224/)

