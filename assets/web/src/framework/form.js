class Form {
    constructor(fields = [], onSubmit = null) {
        this.fields = fields;
        this.onSubmit = onSubmit;
        this.values = {};
        this.selectedField = 0;
        this.container = document.createElement("div");
        this.container.classList.add("form");
        
        // Устанавливаем значения по умолчанию
        this.fields.forEach(field => {
            if (field.defaultValue !== undefined) {
                this.values[field.name] = field.defaultValue;
            }
        });
    }

    addField(name, label, type, options = {}) {
        this.fields.push({
            name,
            label,
            type,
            placeholder: options.placeholder,
            ...options
        });
    }

    getValue(name) {
        return this.values[name];
    }

    setValue(name, value) {
        this.values[name] = value;
    }

    getAllValues() {
        return this.values;
    }

    handleKey(key) {
        if (key === "Escape") {
            // Возвращаемся назад
            if (this.app) {
                this.app.router.getCurrentView().goBack();
            }
        } else if (key === "ArrowUp") {
            this.selectedField = this.selectedField > 0 ? this.selectedField - 1 : this.fields.length;
            this.render();
        } else if (key === "ArrowDown") {
            this.selectedField = this.selectedField < this.fields.length ? this.selectedField + 1 : 0;
            this.render();
        } else if (key === "Enter") {
            this.handleEnter();
        }
    }

    handleEnter() {
        if (this.selectedField === this.fields.length) {
            // Кнопка отправки
            if (this.onSubmit) {
                this.onSubmit(this.getAllValues());
            }
        } else {
            // Редактировать поле
            const field = this.fields[this.selectedField];
            this.editField(field);
        }
    }

    editField(field) {
        if (field.type === "select") {
            // Переходим к выбору опции
            this.app.navigate("/select", {
                field: field,
                currentValue: this.values[field.name],
                returnPath: this.app.router.getCurrentPath(),
                onSave: (value) => {
                    this.values[field.name] = value;
                }
            });
        } else if (field.type === "text") {
            // Переходим к вводу текста
            this.app.navigate("/text", {
                field: field,
                currentValue: this.values[field.name],
                returnPath: this.app.router.getCurrentPath(),
                onSave: (value) => {
                    this.values[field.name] = value;
                }
            });
        } else if (field.type === "checkbox") {
            // Переключаем checkbox
            this.values[field.name] = !this.values[field.name];
            this.render();
        }
    }

    setApp(app) {
        this.app = app;
    }

    render() {
        this.container.innerHTML = "";
        
        this.fields.forEach((field, index) => {
            const group = document.createElement("div");
            group.classList.add("form-group");
            
            if (index === this.selectedField) {
                group.classList.add("selected");
            }
            
            const label = document.createElement("label");
            label.textContent = field.label;
            
            const value = document.createElement("div");
            value.classList.add("field-value");
            
            if (field.type === "select") {
                const currentValue = this.values[field.name];
                const choice = field.choices.find(c => c.value === currentValue);
                value.textContent = choice ? choice.text : (field.placeholder || "Select...");
            } else if (field.type === "text") {
                value.textContent = this.values[field.name] || (field.placeholder || "Enter...");
            } else if (field.type === "checkbox") {
                value.textContent = this.values[field.name] ? "✓" : "✗";
            }
            
            group.appendChild(label);
            group.appendChild(value);
            this.container.appendChild(group);
        });
        
        // Кнопка отправки
        const submitButton = document.createElement("button");
        submitButton.textContent = "Save";
        submitButton.classList.add("btn");
        if (this.selectedField === this.fields.length) {
            submitButton.classList.add("selected");
        }
        this.container.appendChild(submitButton);
        
        return this.container;
    }
}

export default Form; 