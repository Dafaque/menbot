import View from "./view.js";
import Menu from "./menu.js";

class SelectView extends View {
    constructor() {
        super();
        this.setTitle("Выбор опции");
        this._menu = null; // Скрываем от App
    }

    renderContent() {
        const container = document.createElement("div");
        container.classList.add("select-edit-container");
        
        // Создаем меню из опций
        const choices = this.data?.field?.choices || [];
        const currentValue = this.data?.currentValue;
        
        this._menu = new Menu();
        
        // Добавляем опции в меню
        choices.forEach(choice => {
            this._menu.addItem(choice.text, () => {
                if (this.data?.onSave) {
                    this.data.onSave(choice.value);
                }
                this.goBack();
            });
        });
        
        this._menu.setApp(this.app);
        
        container.appendChild(this._menu.render());
        this.container.appendChild(container);
    }

    handleKey(key) {
        if (key === "Escape") {
            this.goBack();
        } else if (this._menu) {
            this._menu.handleKey(key);
        }
    }
}

export default SelectView; 