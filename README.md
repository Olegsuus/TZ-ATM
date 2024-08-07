# API для имитации работы банкомата

Этот проект реализует REST API на языке программирования Golang в рамках тестового задания с использованием фреймворка Echo и базы данных SQLite. Он имитирует основные функции банкомата, включая создание аккаунтов, пополнение средств, снятие средств и проверку баланса аккаунтов.

## Функциональность

- **Создание аккаунтов**: Позволяет создавать новые банковские аккаунты.
- **Пополнение средств**: Поддерживает пополнение средств на существующие аккаунты.
- **Снятие средств**: Дает возможность снимать средства со счетов.
- **Проверка баланса**: Позволяет проверять текущий баланс аккаунта.

## Технологии

- **Golang**: Язык программирования для серверной части.
- **Echo**: Веб-фреймворк для Golang, используемый для обработки HTTP запросов.
- **SQLite**: Локальная база данных для хранения информации об аккаунтах и транзакциях.

## Тестирование

### Использование Postman для тестирования API

Для тестирования функционала API вы можете использовать Postman. Это популярное приложение для тестирования API позволяет легко отправлять запросы к вашему серверу и анализировать ответы.

#### Шаги для тестирования с помощью Postman:

1. **Импорт коллекции в Postman**
    - Скачайте файл с коллекцией запросов для Postman (`postman_collection.json`) из репозитория проекта.
    - Откройте Postman.
    - Нажмите на кнопку "Import" в верхней левой части интерфейса.
    - Выберите скачанный файл и подтвердите импорт.

2. **Тестирование операций**
    - В импортированной коллекции в Postman найдите запросы, которые соответствуют различным операциям API:
        - Создание аккаунта
        - Пополнение средств
        - Снятие средств
        - Проверка баланса
    - Заполните необходимые параметры (например, `id` аккаунта и `amount` для суммы транзакции).
    - Отправьте запрос и анализируйте ответ сервера.




## Примеры запросов для тестирования API

### Создание аккаунта
- **Метод:** POST
- **URL:** `/accounts`
- **Тело запроса:** Не требуется
- **Пример в Postman:**
  ```plaintext
  POST http://localhost:8080/accounts


### Пополнение средств
- **Метод:** POST
- **URL:** `/accounts/{id}/deposit`
- **Параметры в URL:**
    - **id**: Идентификатор аккаунта, которому нужно пополнить баланс.
- **Параметры в теле запроса:**
    - **amount**: Сумма, которую нужно зачислить на счет.
- **Пример в Postman:**
  ```plaintext
  POST http://localhost:8080/accounts/{id}/deposit?amount=500


### Снятие средств
- **Метод:** POST
- **URL:** `/accounts/{id}/withdraw`
- **Параметры в URL:**
    - **id**: Идентификатор аккаунта, с которого будут сниматься средства.
- **Параметры в теле запроса:**
    - **amount**: Сумма, которую нужно снять с аккаунта.
- **Пример в Postman:**
  ```plaintext
  POST http://localhost:8080/accounts/{id}/withdraw?amount=250



### Проверка баланса
- **Метод:** GET
- **URL:** `/accounts/{id}/balance`
- **Параметры в URL:**
    - **id**: Идентификатор аккаунта, баланс которого требуется проверить.
- **Пример в Postman:**
  ```plaintext
  GET http://localhost:8080/accounts/{id}/balance