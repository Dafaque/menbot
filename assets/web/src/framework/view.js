class View {
    constructor() {
        this.container = document.createElement("div");
        this.container.classList.add("view");
        this.app = null;
        this.data = null;
        this.path = null;
        this.title = null; // Автоматический title
    }

    setTitle(title) {
        this.title = title;
    }

    render() {
        this.container.innerHTML = "";
        
        // Устанавливаем title из data если он есть
        if (this.data?.field?.label) {
            this.title = this.data.field.label;
        }
        
        // Автоматически добавляем title если он есть
        if (this.title) {
            const titleElement = document.createElement("h1");
            titleElement.textContent = this.title;
            titleElement.classList.add("title");
            this.container.appendChild(titleElement);
        }
        
        // Автоматически передаем app в компоненты
        if (this.app) {
            if (this.form && this.form.setApp) {
                this.form.setApp(this.app);
            }
            if (this.menu && this.menu.setApp) {
                this.menu.setApp(this.app);
            }
        }
        
        // Вызываем renderContent для наследников
        this.renderContent();
        
        return this.container;
    }

    renderContent() {
        // Переопределяется в наследниках
        // Здесь должна быть основная логика рендера
    }

    show() {
        // Просто вызываем render - router сам заменит содержимое
        this.render();
    }

    handleKey(key) {
        // Автоматически хэндлим ESC во всех views
        if (key === "Escape") {
            this.goBack();
        }
    }

    goBack() {
        // Возвращаемся по иерархии URL
        if (this.app && this.app.router) {
            const currentPath = this.app.router.getCurrentPath();
            const parentPath = this.app.router.getParentPath(currentPath);
            
            // Переходим к родительскому пути
            this.navigate(parentPath);
        }
    }

    navigate(path, data = null) {
        if (this.app) {
            this.app.navigate(path, data);
        }
    }
}

export default View; 