import Menu from "../framework/menu.js";

class ChatDetailsView extends Menu {
    constructor(data) {
        super();
        this.setTitle("Chat Details");
        this.addItem("Edit Chat", () => {
            window.app.router.navigate("/chats/details/edit", data);
        });
        this.addItem("Roles", () => {
            window.app.router.navigate("/chats/details/roles", data);
        });
        this.addItem("Users", () => {
            window.app.router.navigate("/chats/details/users", data);
        });
    }
}

export default ChatDetailsView;