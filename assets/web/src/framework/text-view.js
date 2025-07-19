import View from "./view.js";

class TextView extends View {
    constructor(data) {
        super();
        this.setTitle(`Edit ${data.field.label}`);
        this.value = data.currentValue;
        this.placeholder = data.field.placeholder;
        this.onSave = data.onSave;
        this.field = data.field;
    }

    renderContent() {
        const container = document.createElement('div');
        container.className = 'text-edit-container';
        
        const input = document.createElement('input');
        input.type = 'text';
        input.className = 'text-input';
        input.value = this.value || '';
        input.placeholder = this.placeholder || 'Enter text...';
        
        // Устанавливаем курсор в конец
        input.addEventListener('focus', () => {
            input.setSelectionRange(input.value.length, input.value.length);
        });
        
        container.appendChild(input);
        
        // Фокусируем поле
        setTimeout(() => {
            input.focus();
        }, 100);
        
        return container;
    }

    onKeyDown(e) {
        const input = document.querySelector('.text-input');
        
        switch (e.key) {
            case 'Enter':
                e.preventDefault();
                this.save();
                break;
                
            case 'Escape':
                e.preventDefault();
                this.goBack();
                break;
        }
    }

    save() {
        const input = document.querySelector('.text-input');
        const value = input.value;
        
        if (this.onSave) {
            this.onSave(value);
        }
        
        this.goBack();
    }

    cacheable = () => {
        return false;
    }
}

export default TextView; 