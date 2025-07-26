import { Menu } from "kibodo";

class ChatRoleOptionsView extends Menu {
    constructor(data) {
        super();
        this.addItem("Edit users", () => {
            window.app.router.navigate("/chats/details/roles/list/options/manage", this.role);
        });
        this.addItem("Remove role", () => {
            window.app.router.navigate("/chats/details/roles/list/options/remove", this.role);
        });
    }

    appear = (data) => {
        if (data) {
            this.role = data;
        }
        this.setTitle(`Role ${this.role.name} options`);
        this.render();
    }
}

export default ChatRoleOptionsView;