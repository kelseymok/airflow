# Airflow
This is an example repository for Airflow.

From: https://airflow.apache.org/docs/apache-airflow/stable/start/docker.html

## Prerequisites
* Docker Compose (2.1.1): `brew install docker-compose` or `brew upgrade docker-compose` 
  * **Hint:** the Airflow `docker-compose.yaml` requires some new features from the new `docker-compose` spec. If you have errors running the setup steps, check the version of docker-compose to verify you have version 2.1.1 (`docker-compose --version`)
* Docker

## Setup
Download the latest Airflow `docker-compose.yaml`
```bash
curl -LfO 'https://airflow.apache.org/docs/apache-airflow/2.2.2/docker-compose.yaml'
```

Run it:
```bash
docker-compose up
```

View it in the browser:
```bash
open http://localhost:8080
```
Username/password = `airflow/airflow`

## Tutorial
* Run through the [offical Tutorial](https://airflow.apache.org/docs/apache-airflow/stable/tutorial.html)
* HINT: add new DAGs in the `./dag` directory (it is mounted as a volume)

## Notes
* This takes up a fair bit of energy. Make sure you designate 4GB Memory for Docker (Docker > Preferences > Resources)
