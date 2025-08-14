import { Menu } from "kibodo";
import { Role } from "../models";

export default class ChatRoleRemoveView extends Menu {
    role: Role;
    constructor(data: Role) {
        super(
            [
                {
                    label: "Yes, im sure i want to remove this role from the chat",
                    action: () => this.removeRole()
                },
                {
                    label: "Back",
                    action: () => window.app.pop(false)
                }
            ]
        );
        this.title = `Remove role ${data.name} from chat ${data.tg_chat_name}`;
        this.role = data;
    }

    async removeRole() {
        fetch(`/api/chats/${this.role.chat_id}/roles/${this.role.id}`, {
            method: "DELETE",
        }).then(response => {
            if (!response.ok) {
                throw new Error("Failed to remove role: "+ response.statusText);
            }
            window.app.pop(true);
        }).catch(error => {
            this.error = error.message;
            this.render();
        });
    }
}
