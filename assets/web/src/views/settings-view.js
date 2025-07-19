import View from "../framework/view.js";
import Form from "../framework/form.js";

class SettingsView extends View {
    constructor() {
        super();
        this.setTitle("Настройки");
        this.form = new Form();
        
        this.form.addField("theme", "Тема", "select", {
            choices: [
                { value: "dark", text: "Темная" },
                { value: "light", text: "Светлая" },
                { value: "auto", text: "Авто" }
            ],
            defaultValue: "dark",
            placeholder: "Выберите тему..."
        });
        
        this.form.addField("language", "Язык", "select", {
            choices: [
                { value: "ru", text: "Русский" },
                { value: "en", text: "English" }
            ],
            defaultValue: "ru",
            placeholder: "Выберите язык..."
        });
        
        this.form.addField("notifications", "Уведомления", "checkbox", {
            defaultValue: true
        });
        
        this.form.addField("username", "Имя пользователя", "text", {
            defaultValue: "admin",
            placeholder: "Введите имя..."
        });
        
        this.form.addField("email", "Email", "text", {
            defaultValue: "",
            placeholder: "Введите email..."
        });
        
        this.form.onSubmit = (values) => {
            console.log("Settings saved:", values);
            this.navigate("/menu");
        };
    }

    renderContent() {
        const formContainer = this.form.render();
        this.container.appendChild(formContainer);
    }
}

export default SettingsView; 