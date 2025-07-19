import View from "./view.js";
import Menu from "./menu.js";

class SelectView extends View {
    constructor() {
        super();
        this.menu = new Menu();
        
        // Переопределяем menu.handleKey чтобы перехватывать ESC
        const originalHandleKey = this.menu.handleKey.bind(this.menu);
        this.menu.handleKey = (key) => {
            if (key === "Escape") {
                this.navigate(this.data.returnPath);
            } else {
                originalHandleKey(key);
            }
        };
    }

    renderContent() {
        // Очищаем и создаем меню заново
        this.menu = new Menu();
        
        // Снова переопределяем handleKey для нового меню
        const originalHandleKey = this.menu.handleKey.bind(this.menu);
        this.menu.handleKey = (key) => {
            if (key === "Escape") {
                this.navigate(this.data.returnPath);
            } else {
                originalHandleKey(key);
            }
        };
        
        if (this.data?.field?.choices) {
            this.data.field.choices.forEach(choice => {
                this.menu.addItem(choice.text, () => {
                    if (this.data.onSave) {
                        this.data.onSave(choice.value);
                    }
                    this.navigate(this.data.returnPath);
                });
            });
        }
        
        const menuContainer = this.menu.render();
        this.container.appendChild(menuContainer);
    }
}

export default SelectView; 