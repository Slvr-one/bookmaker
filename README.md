# bookmaker_app:

This app was designed as a simple web server for some minimal bookmaking, aka gambling.
It allows investors to participate in various events.
The app is connected to a MongoDB database for persistancy between events.
Built with Go, utilining MongoDB, Nginx, Docker & Compose.


## Breakdown:
This app is part of my Portfolio project, which includes:
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

####  /home
####  /assets

####  /health
####  /metrics


#### /LH         (GET all available horses)
#### /GH/{name}  (GET a specific horse)
#### /UH/{name}  (Update a specific horse)
#### /invest/{horse}/{amount}


| Path | Method | Description |
| :-------- | :------- | 
| `/home` | `GET` | 
| `/assets` | `GET` | 
| `/health` | `GET` | 
| `/metrics` | `GET` | 






## Tech Stack

**Client:** HTML, CSS, Bootstrap, JS, Nginx

**Server:** GoLang, MongoDB

**CI/CD:** Jenkins / Github Actions

**Cloud:** AWS / GCP

**IAC:** Terraform, Ansible


## 🔗 Links

[![linkedin](https://img.shields.io/badge/linkedin-0A66C2?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/dvir-gross-929252224/)

