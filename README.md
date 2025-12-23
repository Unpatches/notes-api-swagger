## Имя: Дорджиев Виктор
## Группа: ЭФМО-02-25
# Проект notes-api-swagger

Цели
1. Освоить основы спецификации OpenAPI (Swagger) для REST API.
2. Подключить автогенерацию документации к проекту из ПЗ 11(notes-api).
3. Научиться публиковать интерактивную документацию (SwaggerUI / ReDoc) на эндпоинте GET /docs.
4. Синхронизировать код и спецификацию (комментарии-аннотации → генерация) и/или «schema-first» (генерация кода из openapi.yaml).
5. Подготовить процесс обновления документации (Makefile/скрипт).

Code-First (аннотационный)

Программист пишет код API (хэндлеры, маршруты) и добавляет над
функциями специальные комментарии-аннотации, из которых инструменты
вроде swaggo/swag автоматически генерируют OpenAPI-файл.

Преимущества:
- быстрое подключение к уже готовому проекту;
- минимальная ручная работа с YAML;
- простая поддержка актуальности при изменениях кода.

Недостатки:
- ограниченная гибкость описания сложных сценариев;
- возможны расхождения, если аннотации забыли обновить.

---

## Установка и запуск

(Необходимы предустановленные Go версии 1.22 и выше и Git)

Клонировать репозиторий:

```
git clone <URL_РЕПОЗИТОРИЯ>
cd notes-api-swagger
```

Команда запуска сервера:

```
make swagger
```
```
make run
```
------

## Структура проекта

```plaintext
notes-api/
├── api/
│   └── openapi.yaml                
├── cmd/
│   └── api/
│       └── main.go                 
├── docs/
│   ├── docs.go                      
│   ├── swagger.json                 
│   └── swagger.yaml                 
├── internal/
│   ├── core/
│   │   └── note.go                 
│   ├── http/
│   │   ├── router.go               
│   │   └── handlers/
│   │       └── notes.go            
│   └── repo/
│       └── note_mem.go             
├── Makefile                         
├── go.mod                           
├── go.sum                          
└── README.md      
```

## Отчёт о проделанной работе
### Снимки аннотаций

<img width="643" height="313" alt="image" src="https://github.com/user-attachments/assets/c1a11f0e-0cac-4595-bc3c-55905cea320c" />

<img width="631" height="310" alt="image" src="https://github.com/user-attachments/assets/d45a8d0b-d4f8-4583-84d7-e6f6f50bd777" />

<img width="606" height="404" alt="image" src="https://github.com/user-attachments/assets/a7dd4113-8af8-4d04-aed1-615280abb2f0" />

### Скриншот работающей страницы Swagger UI

<img width="1871" height="1005" alt="image" src="https://github.com/user-attachments/assets/fb511855-830b-48ad-b2b2-3ffd5e804165" />

### Команда генерации документации

<img width="677" height="195" alt="image" src="https://github.com/user-attachments/assets/61ff6620-0e83-4c4f-82ce-aa4cbc163517" />

### Работа сервера

<img width="688" height="609" alt="image" src="https://github.com/user-attachments/assets/fb79d248-3f4f-4033-becd-810e82f009bb" />

