import View from "./view.js";

class TextView extends View {
    constructor() {
        super();
        this.setTitle("Редактирование текста");
        this.input = null;
        this.caretPosition = 0;
    }

    renderContent() {
        const container = document.createElement("div");
        container.classList.add("text-edit-container");
        
        // Создаем видимый input для ввода
        this.input = document.createElement("input");
        this.input.type = "text";
        this.input.classList.add("text-input");
        this.input.value = this.data?.currentValue || "";
        
        container.appendChild(this.input);
        this.container.appendChild(container);
        
        // Устанавливаем фокус после добавления в DOM
        setTimeout(() => {
            this.input.focus();
        }, 100);
    }

    save() {
        if (this.data?.onSave) {
            this.data.onSave(this.input.value);
        }
        this.goBack();
    }

    handleKey(key) {
        if (key === "Enter") {
            this.save();
        } else if (key === "Escape") {
            this.goBack();
        }
        // Не вызываем super.handleKey() чтобы избежать двойной обработки ESC
    }
}

export default TextView; 