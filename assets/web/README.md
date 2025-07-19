# Terminal UI Framework

Минимальный фреймворк для создания терминальных UI в стиле Debian installer с иерархической навигацией.

## Архитектура

### Основные компоненты:

1. **App** - главный класс приложения, централизованная обработка клавиатуры
2. **Router** - иерархическая навигация с автоматической регистрацией маршрутов
3. **View** - базовый класс для экранов с автоматическим title
4. **Menu** - компонент вертикального меню с клавиатурной навигацией
5. **Form** - компонент для работы с формами и отдельными экранами редактирования
6. **TextView** - экран редактирования текста
7. **SelectView** - экран выбора опций

## Использование

### 1. Создание приложения

```javascript
import App from "./framework/app.js";

const app = new App();
```

### 2. Создание View

```javascript
import View from "./framework/view.js";

class MyView extends View {
    constructor() {
        super();
        this.setTitle("Мой экран"); // Автоматический title
    }

    renderContent() {
        // Основная логика рендера
        const content = document.createElement("div");
        content.textContent = "Содержимое экрана";
        this.container.appendChild(content);
    }

    handleKey(key) {
        if (key === "Escape") {
            this.goBack(); // Автоматический возврат по иерархии
        }
    }
}
```

### 3. Создание меню

```javascript
import Menu from "./framework/menu.js";

const menu = new Menu();
menu.addItem("Опция 1", () => {
    console.log("Выбрано: Опция 1");
});
menu.addItem("Опция 2", () => {
    console.log("Выбрано: Опция 2");
});
```

### 4. Регистрация маршрутов

```javascript
app.registerRoute("/", MenuView);
app.registerRoute("/menu/settings", SettingsView);
app.registerRoute("/menu/info", InfoView);
```

**Автоматически создаются маршруты для редактирования:**
- `/menu/settings/text-edit` - редактирование текста
- `/menu/settings/select-edit` - выбор опций

### 5. Запуск приложения

```javascript
app.start("/");
```

## Навигация

### Иерархическая навигация

Фреймворк использует иерархическую навигацию по URL:

- `/menu` → `/menu/settings` → `/menu/settings/text-edit`
- ESC возвращает к родительскому пути автоматически

### Управление

#### Меню
- **Стрелки вверх/вниз** - навигация по пунктам меню
- **Enter** - выбор пункта меню
- **Escape** - возврат назад

#### Формы
- **Стрелки вверх/вниз** - навигация по полям формы
- **Enter** - войти в режим редактирования поля
- **Escape** - выйти из режима редактирования
- **Enter** (в режиме редактирования) - сохранить и вернуться
- **Enter** (на кнопке Save) - отправить форму

#### Редактирование текста
- **Ввод текста** - обычный ввод
- **Enter** - сохранить и вернуться
- **Escape** - отменить и вернуться

#### Выбор опций
- **Стрелки вверх/вниз** - навигация по опциям
- **Enter** - выбрать опцию и вернуться
- **Escape** - отменить и вернуться

## Создание форм

### Использование Form компонента

```javascript
import Form from "../framework/form.js";

class SettingsView extends View {
    constructor() {
        super();
        this.setTitle("Настройки");
        
        this.form = new Form([], (values) => {
            console.log("Форма отправлена:", values);
        });
        
        // Добавляем поля
        this.form.addField("name", "Имя", "text", {
            defaultValue: "Пользователь",
            placeholder: "Введите имя"
        });
        
        this.form.addField("type", "Тип", "select", {
            choices: [
                { value: "user", text: "Пользователь" },
                { value: "admin", text: "Администратор" }
            ],
            defaultValue: "user"
        });
        
        this.form.addField("active", "Активен", "checkbox", {
            defaultValue: true
        });
    }

    renderContent() {
        const formContainer = this.form.render();
        this.container.appendChild(formContainer);
    }
}
```

### Типы полей

#### Text
```javascript
this.form.addField("name", "Имя", "text", {
    defaultValue: "Значение по умолчанию",
    placeholder: "Подсказка"
});
```

#### Select
```javascript
this.form.addField("type", "Тип", "select", {
    choices: [
        { value: "option1", text: "Опция 1" },
        { value: "option2", text: "Опция 2" }
    ],
    defaultValue: "option1"
});
```

#### Checkbox
```javascript
this.form.addField("active", "Активен", "checkbox", {
    defaultValue: true
});
```

## API

### App

- `registerRoute(path, viewClass)` - регистрация маршрута
- `start(path)` - запуск приложения
- `navigate(path, data)` - навигация

### View

- `setTitle(title)` - установить заголовок
- `renderContent()` - рендер содержимого (переопределяется)
- `handleKey(key)` - обработка клавиш (переопределяется)
- `goBack()` - возврат по иерархии URL
- `navigate(path, data)` - навигация

### Menu

- `addItem(text, action)` - добавить пункт меню
- `handleKey(key)` - обработка клавиш
- `render()` - рендер меню

### Form

- `addField(name, label, type, options)` - добавить поле формы
- `getValue(name)` - получить значение поля
- `setValue(name, value)` - установить значение поля
- `getAllValues()` - получить все значения формы
- `handleKey(key)` - обработка клавиш
- `render()` - рендер формы

### Router

- `register(path, viewClass)` - регистрация маршрута
- `navigate(path, app, data)` - навигация
- `getCurrentPath()` - получить текущий путь
- `getParentPath(path)` - получить родительский путь
- `getChildPath(parentPath, childSegment)` - получить дочерний путь

## Особенности

### Автоматическая регистрация маршрутов

При регистрации маршрута `/settings` автоматически создаются:
- `/settings/text-edit` - для редактирования текста
- `/settings/select-edit` - для выбора опций

### Централизованная обработка клавиатуры

App обрабатывает все клавиатурные события и передает их в соответствующие компоненты.

### Иерархическая навигация

Навигация работает строго по иерархии URL, ESC всегда возвращает к родительскому пути.

### Автоматический title

Базовый View автоматически отображает title если он установлен через `setTitle()`.

## Стили

Фреймворк использует терминальные стили с зеленым текстом на черном фоне. Курсор скрыт глобально для терминального эффекта.

## Примеры

Смотрите файлы в папке `views/` для примеров использования:

- `menu-view.js` - главное меню
- `settings-view.js` - экран настроек с формой
- `info-view.js` - информационный экран
- `api-view.js` - пример работы с API 