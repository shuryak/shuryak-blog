# shuryak-blog

**shuryak-blog** — блог, написанный с применением микросервисной архитектуры, построенной на фреймворке
[go-micro](https://github.com/go-micro/go-micro).

## Цели и задачи

Основная задача разработки блога — изучение мной микросервисной архитектуры.

## Структура проекта

Для разработки проекта была выбрана [Monorepo](https://earthly.dev/blog/golang-monorepo/)-структура, поскольку:

- Я единолично занимаюсь проектом;
- Для всех backend-сервисов используется один язык программирования — **Go**;
- Учитывая вышеперечисленные моменты, становится проще разделять кодовую базу между сервисами.

Монорепозиторий в проекте реализуется следующим образом:

```go
shuryak-blog/
├── cmd/
│   └── название_сервиса/
│       ├── main.go <- точка входа в сервис
│       └── plugins.go <- плагины go-micro
├── internal/
│   └── название_сервиса/
│       ├── app/
│       │   └── app.go <- основная функция для запуска сервиса
│       ├── config/
│       │   ├── config.go <- Go-структура конфигурации сервиса
│       │   └── config.yml <- YAML-конфигурация сервиса
│       ├── ... <- специфичные для сервиса пакеты
│       └── Dockerfile
├── pkg <- Межсервисные пакеты
├── proto/
│   └── название_сервиса/
│       └── *.proto
└── k8s <- YAML-манифесты Kubernetes
```

> Основной код сервисов хранится в каталоге `internal` вместо `services` для большей совместимости с
> [Standard Go Project Layout](https://github.com/golang-standards/project-layout).

## Полезные ссылки

- [So you need to wait for some Kubernetes resources?](https://vadosware.io/post/so-you-need-to-wait-for-some-kubernetes-resources/)
  — о том, как дождаться готовности ресурсов Kubernetes;
- [Про токены, JSON Web Tokens (JWT), аутентификацию и авторизацию. Token-Based Authentication](https://gist.github.com/zmts/802dc9c3510d79fd40f9dc38a12bccfc)
  — всё, что нужно знать про распределённую аутентификацию и авторизацию с помощью JWT-токенов;
- [Contract between Skaffold and Custom Build Script](https://skaffold.dev/docs/builders/builder-types/custom/#contract-between-skaffold-and-custom-build-script)
  — переменные окружения для взаимодействия Skaffold с пользовательским build-скриптом.
- ...
