class View {
    constructor() {
        this.data = null;
        this.container = null;
        this.title = null; // Автоматический title
    }

    setTitle(title) {
        this.title = title;
    }

    render() {
        // Создаем контейнер
        this.container = document.createElement('div');
        this.container.className = 'view';
        
        // Устанавливаем title из data если он есть
        if (this.data?.field?.label) {
            this.title = this.data.field.label;
        }
        
        // Добавляем заголовок если есть
        if (this.title) {
            const titleElement = document.createElement('h1');
            titleElement.textContent = this.title;
            titleElement.classList.add('title');
            this.container.appendChild(titleElement);
        }
        
        // Рендерим содержимое
        const child = this.renderContent();
        if (child) {
            this.container.appendChild(child);
        }
        
        // Заменяем содержимое app
        const appElement = document.getElementById('app');
        if (appElement) {
            appElement.innerHTML = '';
            appElement.appendChild(this.container);
        }
        
        return this.container;
    }

    renderContent() {
        // Переопределяется в наследниках
    }

    goBack() {
        if (this.app && this.app.router) {
            const currentPath = window.app.router.getCurrentPath();
            const parentPath = this.getParentPath(currentPath);
            if (parentPath) {
                this.app.router.navigate(parentPath);
            }
        }
    }

    getParentPath(path) {
        const parts = path.split('/').filter(part => part);
        if (parts.length === 0) {
            return null;
        }
        parts.pop();
        if (parts.length === 0) {
            return "/";
        }
        return "/" + parts.join("/");
    }

    // Методы для обработки событий (переопределяются в наследниках)
    onKeyDown(e) {
        // Переопределяется в наследниках
    }

    onKeyUp(e) {
        // Переопределяется в наследниках
    }

    onSubmit(e) {
        // Переопределяется в наследниках
    }

    // Не кэшировать вью
    cacheable() {
        return true;
    }

    // Вью показана
    appear(args) {}
}

export default View; 