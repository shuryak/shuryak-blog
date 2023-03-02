# shuryak-blog

## Сборка образов Docker для Kubernetes

[How to Run Locally Built Docker Images in Kubernetes](https://medium.com/swlh/how-to-run-locally-built-docker-images-in-kubernetes-b28fbc32cc1d).

- Сервис `user`:

  ```bash
  docker build -f internal/user/Dockerfile -t user-server:latest .
  ```

- Сервис `articles`:

  ```bash
  docker build -f internal/articles/Dockerfile -t articles-server:latest .
  ```
