import { Menu } from "kibodo";

class ChatRolesView extends Menu {
    constructor(data) {
        super();
        this.setTitle("Chat Roles");
        this.addItem("Add Role", () => {
            window.app.router.navigate("/chats/details/roles/add", data);
        });
        this.addItem("List Roles", () => {
            window.app.router.navigate("/chats/details/roles/list", data);
        });
    }
}

export default ChatRolesView;