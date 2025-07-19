import View from "./view.js";

class TextView extends View {
    constructor() {
        super();
        this.input = null;
    }

    renderContent() {
        // Показываем текущее значение
        const valueDisplay = document.createElement("div");
        valueDisplay.classList.add("value-display");
        valueDisplay.textContent = this.data?.currentValue || "";
        
        // Создаем input для ввода (скрытый)
        this.input = document.createElement("input");
        this.input.type = "text";
        this.input.classList.add("hidden-input");
        this.input.value = this.data?.currentValue || "";
        
        // Обработчики для input
        this.input.addEventListener("input", (e) => {
            valueDisplay.textContent = e.target.value || "";
        });
        
        this.input.addEventListener("keydown", (e) => {
            if (e.key === "Enter") {
                e.preventDefault();
                e.stopPropagation();
                if (this.data?.onSave) {
                    this.data.onSave(this.input.value);
                }
                this.navigate(this.data.returnPath);
            } else if (e.key === "Escape") {
                e.preventDefault();
                e.stopPropagation();
                this.navigate(this.data.returnPath);
            }
        });
        
        this.container.appendChild(valueDisplay);
        this.container.appendChild(this.input);
        
        // Фокусируемся на input
        setTimeout(() => {
            this.input.focus();
        }, 100);
    }

    handleKey(key) {
        // Если input в фокусе, не обрабатываем клавиши
        if (this.input && this.input === document.activeElement) {
            return;
        }
        
        // Если input не в фокусе, фокусируемся на нем
        if (this.input) {
            this.input.focus();
        }
    }
}

export default TextView; 