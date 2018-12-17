# API-exercise

This exercise is to assess your technical proficiency with Software Engineering, DevOps and Infrastructure tasks.
There is no need to do all the exercises, but try to get as much done as you can, so we can get a good feel of your skillset.  Don't be afraid to step outside your comfort-zone and try something new.

If you have any questions, feel free to reach out to us.

## Exercise

This exercise is split in several subtasks. We are curious to see where you feel most comfortable and where you struggle.

### 0. Fork this repository
All your changes should be made in a **private** fork of this repository. When you're done please, please:
* Share your fork with the **container-solutions-test**** user (Settings -> Members -> Share with Member)
* Make sure that you grant this user the Reporter role, so that our reviewers can check out the code using Git.
* Reply to the email that asked you to do this API exercise, with a link to the repository that the **container-solutions-test** user now should have access to.

### 1. Setup & fill database
In the root of this project you'll find a csv-file with passenger data from the Titanic. Create a database and fill it with the given data. SQL or NoSQL is your choice.

### 2. Create an API
Create an HTTP-API (e.g. REST) that allows reading & writing (maybe even updating & deleting) data from your database.
Tech stack and language are your choice. The API we would like you to implement is described in [API.md](./API.md)

### 3. Dockerize
Automate setup of your database & API with Docker, so it can be run everywhere comfortably with one or two commands.
The following build tools are acceptable:
 * docker
 * docker-compose
 * groovy
 * minikube (see 4.)

No elaborate makefiles please.

#### Hints

- [Docker Install](https://www.docker.com/get-started)

### 4. Deploy to Kubernetes
Enable your Docker containers to be deployed on a Kubernetes cluster.

#### Hints

- Don't have a Kubernetes cluster to test against?
  - [MiniKube](https://kubernetes.io/docs/setup/minikube/) (free, local)
  - [GKE](https://cloud.google.com/kubernetes-engine/) (more realistic deployment, may cost something)

### 5. Whatever you can think of
Do you have more ideas to optimize your workflow or automate deployment? Feel free to go wild and dazzle us with your solutions.
