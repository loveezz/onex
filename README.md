# Products API

- Gin/Echo (роутинг)
- PostgreSQL (хранение данных)
- Docker (контейнеризация)

## Запуск

```bash
docker-compose up --build

# Остановка
docker-compose down

## Endpoints

GET /api/products  - получить весь товар

GET /api/products/{id} - получить товар по ID

POST /api/products - создать товар  Body: {"name" : "alabaster", "price": 999999.99}

PUT /api/products/{id} - обновить товар {"price": 10000.10}

DELETE /api/products/{id} - удалить товар 


