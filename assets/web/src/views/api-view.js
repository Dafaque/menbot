import View from "../framework/view.js";
import Form from "../framework/form.js";

class ApiView extends View {
    constructor() {
        super();
        this.setTitle("API Тестер");
        this.form = new Form([], (values) => {
            console.log(values);
            this.makeApiCall(values);
        });
        
        // Добавляем поля
        this.form.addField("method", "Метод", "select", {
            choices: [
                { text: "GET", value: "GET" },
                { text: "POST", value: "POST" },
                { text: "PUT", value: "PUT" },
                { text: "DELETE", value: "DELETE" }
            ],
            defaultValue: "GET"
        });
        
        this.form.addField("url", "URL", "text", {
            defaultValue: "https://jsonplaceholder.typicode.com/posts/1"
        });
        
        this.form.addField("data", "Данные", "text", {
            defaultValue: ""
        });
        
        this.result = null;
    }

    renderContent() {
        const formContainer = this.form.render();
        this.container.appendChild(formContainer);
        
        // Добавляем блок с результатом
        if (this.result) {
            const resultContainer = document.createElement("div");
            resultContainer.classList.add("result-container");
            
            const resultTitle = document.createElement("h3");
            resultTitle.textContent = "Результат:";
            
            const resultContent = document.createElement("pre");
            resultContent.textContent = JSON.stringify(this.result, null, 2);
            
            resultContainer.appendChild(resultTitle);
            resultContainer.appendChild(resultContent);
            this.container.appendChild(resultContainer);
        }
    }

    async makeApiCall(values) {
        try {
            const { method, url, data } = values;
            
            const options = {
                method: method,
                headers: {
                    'Content-Type': 'application/json',
                }
            };
            
            if (method !== 'GET' && data) {
                try {
                    options.body = data;
                } catch (e) {
                    alert("Ошибка в JSON данных");
                    return;
                }
            }
            
            const response = await fetch(url, options);
            const responseData = await response.json();
            
            this.result = {
                status: response.status,
                statusText: response.statusText,
                headers: Object.fromEntries(response.headers.entries()),
                data: responseData
            };
            
            this.render();
            
        } catch (error) {
            this.result = {
                error: error.message
            };
            this.render();
        }
    }
}

export default ApiView; 