import View from "./view.js";

class Form extends View {
    constructor(fields = []) {
        super();
        this.fields = fields;
        this.currentFieldIndex = 0;
        this.values = {};
        
        // Инициализируем значения
        this.fields.forEach(field => {
            this.values[field.name] = field.value || '';
        });
    }

    addField(name, label, type, options = {}) {
        const field = {
            name,
            label,
            type,
            placeholder: options.placeholder,
            readonly: options.readonly || false,
            value: options.value,
            defaultValue: options.defaultValue,
            ...options
        };
        
        this.fields.push(field);
        
        // Устанавливаем defaultValue сразу при добавлении поля
        if (field.defaultValue !== undefined) {
            this.values[field.name] = field.defaultValue;
        } else if (field.value !== undefined) {
            this.values[field.name] = field.value;
        }
    }

    renderContent() {
        const form = document.createElement('form');
        form.className = 'form';
        let firstEditableFieldFound = false;
        this.fields.forEach((field, index) => {
            if (!firstEditableFieldFound && !field.readonly) {
                firstEditableFieldFound = true;
                this.currentFieldIndex = index;
            }
            
            const fieldGroup = document.createElement('div');
            fieldGroup.className = 'form-group';
            fieldGroup.dataset.fieldIndex = index;
            
            if (index === this.currentFieldIndex) {
                fieldGroup.classList.add('selected');
                fieldGroup.focus();
            }
            
            if (field.readonly) {
                fieldGroup.classList.add('readonly');
            }
            
            const label = document.createElement('label');
            label.textContent = field.label || field.name;
            fieldGroup.appendChild(label);
            
            if (field.readonly) {
                const value = document.createElement('div');
                value.className = 'field-value readonly';
                value.textContent = this.values[field.name] || field.value || '';
                fieldGroup.appendChild(value);
            } else if (field.type === 'checkbox') {
                const checkbox = document.createElement('div');
                checkbox.className = 'field-value checkbox';
                checkbox.textContent = this.values[field.name] || field.value ? '[X]' : '[ ]';
                checkbox.dataset.fieldName = field.name;
                fieldGroup.appendChild(checkbox);
            } else {
                const value = document.createElement('div');
                value.className = 'field-value';
                value.textContent = this.values[field.name] || field.value || field.placeholder || 'Enter...';
                fieldGroup.appendChild(value);
            }
            
            form.appendChild(fieldGroup);
        });
        
        // Добавляем кнопку Save
        const saveButton = document.createElement('button');
        saveButton.className = 'btn save-button';
        saveButton.textContent = 'Save';
        saveButton.type = 'submit';
        saveButton.dataset.fieldIndex = this.fields.length;
        
        if (this.currentFieldIndex === this.fields.length) {
            saveButton.classList.add('selected');
        }
        
        form.appendChild(saveButton);
        
        return form;
    }

    onKeyDown(e) {
        switch (e.key) {
            case 'ArrowUp':
                e.preventDefault();
                this.navigateField(-1);
                break;
                
            case 'ArrowDown':
                e.preventDefault();
                this.navigateField(1);
                break;
                
            case 'Tab':
                e.preventDefault();
                this.navigateField(1);
                break;
                
            case 'Enter':
                e.preventDefault();
                this.handleEnter();
                break;
        }
    }

    navigateField(direction) {
        // Убираем выделение с текущего элемента
        const currentSelected = document.querySelector('.form-group.selected, .save-button.selected');
        if (currentSelected) {
            currentSelected.classList.remove('selected');
        }
        
        // Вычисляем новый индекс
        let newIndex = this.currentFieldIndex + direction;
        const maxIndex = this.fields.length; // включая кнопку Save
        
        if (newIndex < 0) {
            newIndex = maxIndex;
        } else if (newIndex > maxIndex) {
            newIndex = 0;
        }
        
        // Пропускаем readonly поля при навигации
        if (newIndex < this.fields.length) {
            const field = this.fields[newIndex];
            if (field && field.readonly) {
                // Пропускаем readonly поле
                if (direction > 0) {
                    newIndex++;
                } else {
                    newIndex--;
                }
                // Рекурсивно вызываем для следующего поля
                this.currentFieldIndex = newIndex;
                this.navigateField(direction);
                return;
            }
        }
        
        this.currentFieldIndex = newIndex;
        
        // Выделяем новый элемент
        if (newIndex === this.fields.length) {
            // Кнопка Save
            const saveButton = document.querySelector('.save-button');
            if (saveButton) {
                saveButton.classList.add('selected');
            }
        } else {
            // Поле формы
            const fieldGroup = document.querySelector(`[data-field-index="${newIndex}"]`);
            if (fieldGroup) {
                fieldGroup.classList.add('selected');
            }
        }
    }

    handleEnter() {
        if (this.currentFieldIndex === this.fields.length) {
            // Кнопка Save
            this.save();
        } else {
            // Поле формы
            const field = this.fields[this.currentFieldIndex];
            if (field && !field.readonly) {
                this.editField(field);
            }
        }
    }

    editField(field) {
        if (field.type === 'checkbox') {
            // Переключаем checkbox
            this.values[field.name] = !this.values[field.name];
            const checkbox = document.querySelector(`[data-field-name="${field.name}"]`);
            if (checkbox) {
                checkbox.textContent = this.values[field.name] ? '[X]' : '[ ]';
            }
        } else if (field.type === 'select') {
            // Переходим к выбору опции
            const currentPath = window.app.router.getCurrentPath();
            const selectPath = this.getChildPath(currentPath, "select-edit");
            
            window.app.router.navigate(selectPath, {
                field: field,
                currentValue: this.values[field.name],
                onSave: (value) => {
                    this.values[field.name] = value;
                    this.render();
                }
            });
        } else {
            // Переходим к вводу текста
            const currentPath = window.app.router.getCurrentPath();
            const textPath = this.getChildPath(currentPath, "text-edit");
            
            window.app.router.navigate(textPath, {
                field: field,
                currentValue: this.values[field.name],
                onSave: (value) => {
                    this.values[field.name] = value;
                    this.render();
                }
            });
        }
    }

    getChildPath(parentPath, childSegment) {
        if (parentPath === "/") {
            return "/" + childSegment;
        }
        return parentPath + "/" + childSegment;
    }

    save() {
        // Обрабатываем стилизованные checkbox
        const checkboxes = document.querySelectorAll('.checkbox');
        checkboxes.forEach(checkbox => {
            const fieldName = checkbox.dataset.fieldName;
            if (fieldName) {
                this.values[fieldName] = checkbox.textContent === '[X]';
            }
        });
        
        // Вызываем callback если есть
        if (this.onSave) {
            this.onSave(this.values);
        }
    }

    onSubmit(e) {
        e.preventDefault();
        this.save();
    }
}

export default Form; 