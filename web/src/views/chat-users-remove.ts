import { Menu } from "kibodo";
import { User } from "../models";

export default class ChatUsersRemoveView extends Menu {
    user: User;
    constructor(data: User) {
        super([
            {
                label: "Yes, im sure i want to remove this user from the chat",
                action: () => this.removeUser()
            },
            {
                label: "Back",
                action: () => window.app.pop()
            }
        ]);
        this.title = `Remove user ${data.tg_user_name} from chat ${data.tg_chat_name}`;
        this.user = data;
    }


    async removeUser() {
        fetch(`/api/chats/${this.user.chat_id}/users/${this.user.id}`, {
            method: "DELETE",
        }).then(response => {
            if (!response.ok) {
                throw new Error("Failed to remove user: " + response.statusText);
            }
            window.app.pop(true);
        }).catch(error => {
            this.error = error.message;
            this.render();
        });
    }
    
}
