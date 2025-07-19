class Router {
    constructor() {
        this.routes = new Map();
        this.currentView = null;
        this.currentPath = null;
        this.views = new Map(); // Кэш view'ов
        this.history = []; // История навигации
    }

    register(path, ViewClass) {
        this.routes.set(path, ViewClass);
        
        // Автоматически регистрируем маршруты для редактирования
        this.registerEditRoutes(path, ViewClass);
    }

    // Автоматически регистрируем маршруты для редактирования
    registerEditRoutes(path, ViewClass) {
        // Импортируем view'ы для редактирования
        import("./text-view.js").then(module => {
            const TextView = module.default;
            const textPath = this.getChildPath(path, "text-edit");
            this.routes.set(textPath, TextView);
        });
        
        import("./select-view.js").then(module => {
            const SelectView = module.default;
            const selectPath = this.getChildPath(path, "select-edit");
            this.routes.set(selectPath, SelectView);
        });
    }

    // Получить родительский путь
    getParentPath(path) {
        const parts = path.split('/').filter(part => part);
        if (parts.length <= 1) {
            return "/"; // Корневой путь
        }
        parts.pop(); // Убираем последний сегмент
        return "/" + parts.join("/");
    }

    // Получить дочерний путь
    getChildPath(parentPath, childSegment) {
        if (parentPath === "/") {
            return "/" + childSegment;
        }
        return parentPath + "/" + childSegment;
    }

    navigate(path, app, data = null) {
        const ViewClass = this.routes.get(path);
        if (!ViewClass) {
            throw new Error(`Route ${path} not found`);
        }

        // Проверяем, есть ли уже созданный view для этого пути
        if (!this.views.has(path)) {
            // Создаем новый view только если его еще нет
            this.views.set(path, new ViewClass());
        }

        // Сохраняем текущий view в историю
        if (this.currentView) {
            this.history.push(this.currentPath);
        }

        // Получаем существующий или созданный view
        this.currentView = this.views.get(path);
        this.currentView.app = app;
        this.currentView.data = data;
        this.currentView.path = path;
        this.currentPath = path;

        // Заменяем содержимое app
        const appElement = document.getElementById("app");
        if (!appElement) {
            throw new Error("App container not found");
        }
        
        appElement.innerHTML = "";
        appElement.appendChild(this.currentView.render());
    }

    back() {
        if (this.history.length > 0) {
            const previousPath = this.history.pop();
            
            // Получаем предыдущий view
            const previousView = this.views.get(previousPath);
            if (previousView) {
                this.currentView = previousView;
                
                // Заменяем содержимое app
                const appElement = document.getElementById("app");
                if (appElement) {
                    appElement.innerHTML = "";
                    appElement.appendChild(this.currentView.render());
                }
                return true;
            }
        }
        return false;
    }

    getCurrentView() {
        return this.currentView;
    }

    getCurrentPath() {
        return this.currentPath;
    }
}

export default Router; 