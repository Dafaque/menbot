import { Menu } from "kibodo";

class ChatRoleRemoveView extends Menu {
    constructor(data) {
        super();
        this.addItem("Yes, im sure i want to remove this role from the chat", this.removeRole.bind(this));
        this.addItem("Back", this.goBack.bind(this));
    }

    appear = (data) => {
        if (data) {
            this.role = data;
        }
        this.setTitle(`Remove role ${this.role.name} from chat ${this.role.tg_chat_name}`);
        this.render();
    }

    async removeRole() {
        fetch(`/api/chats/${this.role.chat_id}/roles/${this.role.id}`, {
            method: "DELETE",
        }).then(response => {
            if (response.ok) {
                window.app.router.navigate("/chats/details/roles/list");
            }
        });
    }
}

export default ChatRoleRemoveView;