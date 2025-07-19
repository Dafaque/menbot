class Router {
    constructor() {
        this.routes = new Map();
        this.views = new Map();
        this.currentView = null;
        this.currentPath = null;
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

    // Получить дочерний путь
    getChildPath(parentPath, childSegment) {
        if (parentPath === "/") {
            return "/" + childSegment;
        }
        return parentPath + "/" + childSegment;
    }

    navigate(path, data = null) {
        console.log(`Navigating to: ${path}`);
        
        // Проверяем, есть ли уже созданный view для этого пути
        if (!this.views.has(path)) {
            // Создаем новый view
            const ViewClass = this.routes.get(path);
            if (!ViewClass) {
                console.error(`Route not found: ${path}`);
                return;
            }
            
            this.currentView = new ViewClass(data);
            if (!this.currentView.service()) {
                this.views.set(path, this.currentView);
            }
            
        } else {
            // Используем существующий view из кэша
            this.currentView = this.views.get(path);
        }

        // Устанавливаем view в app
        app.setCurrentView(this.currentView);
        
        // Устанавливаем данные
        this.currentView.app = window.app;
        
        // Рендерим view
        this.currentView.render();
        
        // Сохраняем текущий путь
        this.currentPath = path;
    }

    getCurrentView() {
        return this.currentView;
    }

    getCurrentPath() {
        return this.currentPath;
    }
}

export default Router; 