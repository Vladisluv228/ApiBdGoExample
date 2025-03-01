# ApiBdGoExample

**Пример контейнеризации приложения на go с базой данных и настроенный build and push пайплайн.**

_Проект разработан в рамках НИС Введение в облачные технологии._

## Как запустить?

### Локально

1. **Клонируйте репозиторий:**

   ```bash
   git clone https://github.com/Vladisluv228/ApiBdGoExample.git
   cd ApiBdGoExample
   cd app
2. **Соберите и запустите Docker-образ:**

   ```bash
   docker compose -f 'app\docker-compose.yml' up -d --build

### Через Docker Hub

1. **Скачайте образ из Docker Hub:**

   ```bash
   docker pull vladisluv228/api_bd_go_example:latest

2. **Запустите контейнер:**

   ```bash
   docker run -it vladisluv228/api_bd_go_example:latest

3. **Проверьте список запущенных контейнеров:**

   ```bash
   docker ps
