import View from "../framework/view.js";
import Menu from "../framework/menu.js";

class MenuView extends View {
    constructor() {
        super();
        this.setTitle("Главное меню");
        this.menu = new Menu();
        
        this.menu.addItem("Настройки", () => {
            this.navigate("/menu/settings", { returnPath: "/menu" });
        });
        
        this.menu.addItem("Информация", () => {
            this.navigate("/menu/info", { returnPath: "/menu" });
        });
        
        this.menu.addItem("API Тестер", () => {
            this.navigate("/menu/api", { returnPath: "/menu" });
        });
    }

    renderContent() {
        const menuContainer = this.menu.render();
        this.container.appendChild(menuContainer);
    }
}

export default MenuView; 