# Terminal UI Framework

Минимальный JavaScript фреймворк для создания терминальных UI с иерархической навигацией и клавиатурным управлением.

## Архитектура

### Глобальная архитектура
- **App** - главный класс, эмитит события клавиатуры
- **Router** - маршрутизация с иерархическими URL
- **View** - базовый класс для всех экранов
- **Компоненты** - Menu, Form, Table, TextView, SelectView

### Система событий
Все компоненты подписываются на глобальные события через `window.app`:
- Подписка происходит в `render()` методе
- Отписка происходит при переходе в Router
- Каждый компонент сохраняет ссылку на bound функцию

## Использование

### Создание App
```javascript
import App from "./framework/app.js";

const app = new App();
app.registerRoute("/", IndexView);
app.registerRoute("/chats", ChatsView);
app.start("/");
```

### Создание View
```javascript
import View from "../framework/view.js";
import Menu from "../framework/menu.js";

class IndexView extends View {
    constructor() {
        super();
        this.setTitle("Index");
        this.menu = new Menu();
        
        this.menu.addItem("Chats", () => {
            this.navigate("/chats");
        });
    }

    renderContent() {
        return this.menu;
    }
}
```

### Компоненты

#### Menu
```javascript
const menu = new Menu();
menu.addItem("Option 1", () => console.log("Selected 1"));
menu.addItem("Option 2", () => console.log("Selected 2"));
```

#### Table
```javascript
const table = new Table();
table.setData(
    ["ID", "Name", "Status"],
    [
        { ID: "1", Name: "Item 1", Status: "Active" },
        { ID: "2", Name: "Item 2", Status: "Inactive" }
    ]
);
table.setOnRowSelect((row, index) => console.log(row));
```

#### Form
```javascript
const form = new Form();
form.addField("name", "Name", "text", { placeholder: "Enter name" });
form.addField("status", "Status", "select", {
    choices: [
        { value: "active", text: "Active" },
        { value: "inactive", text: "Inactive" }
    ]
});
form.onSubmit = (values) => console.log(values);
```

## Навигация

### Иерархические URL
- `/` - главная страница
- `/chats` - список чатов
- `/chats/text-edit` - редактирование текста
- `/chats/select-edit` - выбор опции

### Автоматическая регистрация
Фреймворк автоматически регистрирует маршруты для редактирования:
- `{path}/text-edit` - TextView
- `{path}/select-edit` - SelectView

### Клавиатурное управление
- **Arrow Up/Down** - навигация по элементам
- **Enter** - выбор/активация
- **Escape** - возврат назад

## Стили

### Terminal CSS
```css
body {
    background: #000;
    color: #0f0;
    font-family: monospace;
    cursor: none;
}

.selected {
    background: #0f0;
    color: #000;
}
```

### Компоненты
- `.menu` - меню
- `.table` - таблица
- `.form` - форма
- `.text-input` - текстовый ввод с мигающим курсором

## События

### Глобальная система
Все компоненты подписываются на события через `window.app`:
```javascript
// Подписка в render()
if (window.app && !this._subscribed) {
    this._boundHandleKey = this.handleKey.bind(this);
    window.app.on("keydown", this._boundHandleKey);
    this._subscribed = true;
}
```

### Отписка при переходе
Router автоматически отписывает старые компоненты:
```javascript
// Отписка в router.navigate()
if (this.currentView._boundHandleKey) {
    window.app.off("keydown", this.currentView._boundHandleKey);
    this.currentView._subscribed = false;
}
```

## Особенности

### Автоматический title
View автоматически отображает title из data:
```javascript
// data.field.label автоматически становится title
this.data = { field: { label: "Edit Name" } };
```

### Мигающий курсор
Text input имеет терминальный мигающий курсор:
```css
.text-input {
    caret-color: #0f0;
    animation: blink 1s infinite;
}
```

### Row-based навигация
Table поддерживает навигацию по строкам с передачей всей строки данных при выборе.

### Централизованная обработка
Все события клавиатуры обрабатываются централизованно в App и эмитятся для подписчиков. 