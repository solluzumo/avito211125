
Этот проект представляет собой тестовое задание, сервис разработан на **Golang** с использованием **PostgreSQL** в качестве основной базы данных. Сервис упакован в Docker-контейнеры и может быть запущен через Docker Compose.

Реализованы:

* Нагрузочное тестирование двух ключевых методов (папка `tests`).
* Эндпоинт для **массовой деактивации пользователей**, с последующим **переназначением открытых PR** деактивированных пользователей.
* Вопрос, который не был уточнен в ТЗ: при создании команды пользователь автоматически **отвязывается от предыдущей команды и всех открытых PR**, так как больше не принадлежит старой команде.

В `init` можно раскомментировать скрипт для заполнения базы данных тестовыми данными, предназначенными для нагрузочного тестирования.

---

## Стэк

* **Golang**
* **PostgreSQL**
* **Docker / Docker Compose**

---

## Docker

Для сборки и запуска проекта в контейнерах выполните:

```bash
docker compose -f testing.docker-compose.yml build --no-cache
docker compose -f testing.docker-compose.yml up -d
```

---


## Результаты нагрузочного теста


### **CreatePullRequest** (RPS = 5)

* Total requests: 1000
* Successful requests: 1000
* Success rate: 100.000%
* Average latency: 21.79 ms
* p50 latency: 20.98 ms
* p95 latency: 35.03 ms
* p99 latency: 36.87 ms

### **TeamDeactivate** (RPS = 5)

* Total requests: 200
* Successful requests: 200
* Success rate: 100.000%
* Average latency: 12.90 ms
* p50 latency: 14.47 ms
* p95 latency: 19.84 ms
* p99 latency: 24.22 ms

---

