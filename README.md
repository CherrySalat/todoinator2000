На таблице 1 представлен список доступных эндпоинтов

Таблица 1 - Список эндпоинтов
| Эндпоинт  | Метод  | Что делает                            |
|-----------|--------|---------------------------------------|
| /todos    | GET    | Возращает список всех имеющихся задач |
| /todos?id | GET    | Возвращает одну задачу(id)            |
| /todos    | POST   | Добавляет задачу                      |
| /todos?id | PUT    | Изменяет задачу(id)                   |
| /todos?id | DELETE | Удаляет задачу(id)                    |


На таблице 2 представлен список кодов ответа на запросы и их расшифровка

Таблица 2 - Список кодов ответов
| Код запроса | Значение                  |
|-------------|---------------------------|
| 200         | Принято                   |
| 201         | Создано                   |
| 403         | Запрещено                 |
| 404         | Не найдено                |
| 500         | Внутренняя ошибка сервера |
