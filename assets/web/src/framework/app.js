import Router from "./router.js";
import SelectView from "./select-view.js";
import TextView from "./text-view.js";

class App {
    constructor() {
        this.router = new Router();
        this.setupEventListeners();
        this.registerFrameworkRoutes();
    }

    registerRoute(path, ViewClass) {
        this.router.register(path, ViewClass);
    }

    registerFrameworkRoutes() {
        // Автоматически регистрируем framework routes
        this.router.register("/select", SelectView);
        this.router.register("/text", TextView);
    }

    start(path) {
        this.router.navigate(path, this);
    }

    navigate(path, data = null) {
        this.router.navigate(path, this, data);
    }

    setupEventListeners() {
        document.addEventListener("keydown", (e) => {
            const currentView = this.router.getCurrentView();
            if (currentView) {
                // Сначала пробуем передать клавишу в компоненты view
                if (currentView.form && currentView.form.handleKey) {
                    currentView.form.handleKey(e.key);
                } else if (currentView.menu && currentView.menu.handleKey) {
                    currentView.menu.handleKey(e.key);
                } else if (currentView.handleKey) {
                    // Если нет компонентов, передаем в сам view
                    currentView.handleKey(e.key);
                }
            }
        });
    }
}

export default App; 