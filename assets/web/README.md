# Terminal UI Framework

Минимальный фреймворк для создания терминальных UI в стиле Debian installer.

## Архитектура

### Основные компоненты:

1. **App** - главный класс приложения
2. **Router** - навигация между views
3. **View** - базовый класс для экранов
4. **Menu** - компонент вертикального меню
5. **Form** - компонент для работы с формами

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
    }

    render() {
        this.container.innerHTML = "";
        
        const title = document.createElement("h1");
        title.textContent = "Мой экран";
        title.classList.add("title");
        
        this.container.appendChild(title);
        
        return this.container;
    }

    handleKey(key) {
        if (key === "Escape") {
            this.goBack();
        }
    }
}
```

### 3. Создание меню

```javascript
import Menu from "./framework/menu.js";

const menu = new Menu([
    { text: "Опция 1", value: "option1" },
    { text: "Опция 2", value: "option2" }
], (item) => {
    console.log("Выбрано:", item.value);
});
```

### 4. Регистрация маршрутов

```javascript
app.registerRoute("/", MenuView);
app.registerRoute("/settings", SettingsView);
app.registerRoute("/info", InfoView);
```

### 5. Запуск приложения

```javascript
app.start("/");
```

## Навигация

### Меню
- **Стрелки вверх/вниз** - навигация по пунктам меню
- **Enter** - выбор пункта меню

### Формы
- **Стрелки вверх/вниз** - навигация по полям формы
- **Enter** - войти в режим редактирования поля
- **Escape** - выйти из режима редактирования
- **Enter** (в режиме редактирования) - сохранить и перейти к следующему полю
- **Enter** (на кнопке) - отправить форму

### Общее
- **Escape** - возврат назад

## Создание форм

### Простой способ

```javascript
render() {
    this.container.innerHTML = "";
    
    const form = document.createElement("div");
    form.classList.add("form");
    
    // Создаем поля формы
    const input = document.createElement("input");
    input.type = "text";
    input.placeholder = "Введите текст";
    
    const button = document.createElement("button");
    button.textContent = "Сохранить";
    button.classList.add("btn", "btn-primary");
    button.addEventListener("click", () => {
        // Обработка сохранения
    });
    
    form.appendChild(input);
    form.appendChild(button);
    
    this.container.appendChild(form);
    return this.container;
}
```

### Использование Form компонента

```javascript
import Form from "../framework/form.js";

class MyView extends View {
    constructor() {
        super();
        this.form = new Form([], (values) => {
            console.log("Форма отправлена:", values);
        });
        
        // Добавляем поля
        this.form.addField("name", "Имя", "text", {
            placeholder: "Введите имя"
        });
        
        this.form.addField("type", "Тип", "select", {
            choices: [
                { value: "user", text: "Пользователь" },
                { value: "admin", text: "Администратор" }
            ]
        });
        
        this.form.addField("active", "Активен", "checkbox");
    }

    render() {
        this.container.innerHTML = "";
        this.container.appendChild(this.form.render());
        return this.container;
    }

    handleKey(key) {
        this.form.handleKey(key);
    }
}
```

## API

### App

- `registerRoute(path, viewClass)` - регистрация маршрута
- `start(path, data)` - запуск приложения
- `navigate(path, data)` - навигация
- `back()` - возврат назад

### View

- `render()` - рендер view (переопределяется)
- `show()` - показать view
- `handleKey(key)` - обработка клавиш (переопределяется)
- `goBack()` - возврат назад

### Menu

- `addItem(text, value)` - добавить пункт меню
- `select(index)` - выбрать пункт по индексу
- `selectNext()` - следующий пункт
- `selectPrev()` - предыдущий пункт
- `getSelected()` - получить выбранный пункт
- `handleKey(key)` - обработка клавиш

### Form

- `addField(name, label, type, options)` - добавить поле формы
- `getValue(name)` - получить значение поля
- `setValue(name, value)` - установить значение поля
- `getAllValues()` - получить все значения формы
- `focusNext()` - следующий фокус
- `focusPrev()` - предыдущий фокус
- `handleKey(key)` - обработка клавиш
- `render()` - рендер формы

### Router

- `register(path, viewClass)` - регистрация маршрута
- `navigate(path, data)` - навигация
- `back()` - возврат назад
- `getCurrentView()` - получить текущий view

## Стили

Фреймворк использует терминальные стили с зеленым текстом на черном фоне. Все компоненты имеют соответствующие CSS классы для стилизации.

## Примеры

Смотрите файлы в папке `views/` для примеров использования:

- `menu-view.js` - главное меню
- `settings-view.js` - экран настроек с формой
- `info-view.js` - информационный экран
- `api-view.js` - пример работы с API 