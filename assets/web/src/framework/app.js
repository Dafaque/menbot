import Router from "./router.js";
import SelectView from "./select-view.js";
import TextView from "./text-view.js";

export default class App {
    constructor() {
        this.currentView = null;
        this.router = null;
        
        // Глобальный обработчик ESC
        document.addEventListener('keydown', (e) => {
            if (e.key === 'Escape') {
                if (this.currentView && this.currentView.goBack) {
                    this.currentView.goBack();
                }
            }
        });
        
        // Проксирование событий в currentView
        document.addEventListener('keydown', (e) => {
            if (this.currentView && this.currentView.onKeyDown) {
                this.currentView.onKeyDown(e);
            }
        });
        
        document.addEventListener('keyup', (e) => {
            if (this.currentView && this.currentView.onKeyUp) {
                this.currentView.onKeyUp(e);
            }
        });
        
        document.addEventListener('submit', (e) => {
            if (this.currentView && this.currentView.onSubmit) {
                this.currentView.onSubmit(e);
            }
        });
        window.app = this;
    }
    
    setCurrentView(view) {
        this.currentView = view;
    }
    
    setRouter(router) {
        this.router = router;
    }
}