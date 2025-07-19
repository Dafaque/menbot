import View from "./view.js";

class Table extends View {
    constructor(data = []) {
        super();
        this.data = data;
        this.selectedRow = 0;
        this.onSelected = null;
    }

    setData(data) {
        this.data = data;
    }

    setOnSelected(callback) {
        this.onSelected = callback;
    }

    renderContent() {
        const container = document.createElement('div');
        container.className = 'table';
        
        if (this.data.length === 0) {
            const emptyMessage = document.createElement('div');
            emptyMessage.className = 'empty-message';
            emptyMessage.textContent = 'No data available';
            container.appendChild(emptyMessage);
            return container;
        }
        
        // Создаем таблицу
        const table = document.createElement('table');
        table.classList.add('terminal-table');
        
        // Заголовки
        const thead = document.createElement('thead');
        const headerRow = document.createElement('tr');
        
        const headers = Object.keys(this.data[0]);
        headers.forEach(header => {
            const th = document.createElement('th');
            th.textContent = header;
            headerRow.appendChild(th);
        });
        
        thead.appendChild(headerRow);
        table.appendChild(thead);
        
        // Данные
        const tbody = document.createElement('tbody');
        
        this.data.forEach((row, rowIndex) => {
            const tr = document.createElement('tr');
            
            if (rowIndex === this.selectedRow) {
                tr.classList.add('selected');
            }
            
            headers.forEach(header => {
                const td = document.createElement('td');
                td.textContent = row[header];
                tr.appendChild(td);
            });
            
            tbody.appendChild(tr);
        });
        
        table.appendChild(tbody);
        container.appendChild(table);
        
        return container;
    }

    onKeyDown(e) {
        switch (e.key) {
            case 'ArrowUp':
                e.preventDefault();
                this.selectedRow = Math.max(0, this.selectedRow - 1);
                this.updateSelection();
                break;
                
            case 'ArrowDown':
                e.preventDefault();
                this.selectedRow = Math.min(this.data.length - 1, this.selectedRow + 1);
                this.updateSelection();
                break;
                
            case 'Enter':
                e.preventDefault();
                this.selectCurrentRow();
                break;
        }
    }

    updateSelection() {
        const rows = document.querySelectorAll('tbody tr');
        rows.forEach((row, index) => {
            if (index === this.selectedRow) {
                row.classList.add('selected');
            } else {
                row.classList.remove('selected');
            }
        });
    }

    selectCurrentRow() {
        if (this.data.length > 0 && this.onSelected) {
            const selectedData = this.data[this.selectedRow];
            this.onSelected(selectedData);
        }
    }
}

export default Table; 