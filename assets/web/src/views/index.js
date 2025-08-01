import Menu from "../framework/menu.js";

class IndexView extends Menu {
    constructor() {
        super();
        this.setTitle("Index");
        this.addItem("Chats", () => {
            window.app.router.navigate("/chats");
        });
    }
}

export default IndexView; 