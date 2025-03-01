# ApiBdGoExample

**Пример контейнеризации приложения на go с базой данных и настроенный CI/CD пайплайн.**

_Проект разработан в рамках НИС Введение в облачные технологии._

## Как запустить локально?

1. **Клонируйте репозиторий:**

   ```bash
   git https://github.com/Vladisluv228/ApiBdGoExample.git
   cd ApiBdGoExample
   cd app
2. **Соберите и запустите Docker-образ:**

   ```bash
   docker compose -f 'app\docker-compose.yml' up -d --build
