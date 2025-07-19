import View from "../framework/view.js";

class InfoView extends View {
    render() {
        this.container.innerHTML = "";
        
        const title = document.createElement("h1");
        title.textContent = "Информация";
        title.classList.add("title");
        
        const content = document.createElement("div");
        content.classList.add("content");
        content.innerHTML = `
            <h2>О приложении</h2>
            <p>Это простое терминальное приложение в стиле BIOS.</p>
            <p>Используйте стрелки для навигации, Enter для выбора, Escape для возврата.</p>
            <br>
            <h3>Управление:</h3>
            <ul>
                <li>↑↓ - навигация</li>
                <li>Enter - выбрать/редактировать</li>
                <li>Escape - назад</li>
            </ul>
            <br>
            <h3>Возможности:</h3>
            <ul>
                <li>Настройки с сохранением данных</li>
                <li>API тестер с отображением результатов</li>
                <li>BIOS-style редактирование форм</li>
                <li>Полная навигация с клавиатуры</li>
            </ul>
        `;
        
        this.container.appendChild(title);
        this.container.appendChild(content);
        
        return this.container;
    }
}

export default InfoView; 