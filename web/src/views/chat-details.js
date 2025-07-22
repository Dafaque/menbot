import { Menu } from "kibodo";

class ChatDetailsView extends Menu {
    constructor(data) {
        super();
        this.addItem("Edit Chat", () => {
            window.app.router.navigate("/chats/details/edit", this.chat);
        });
        this.addItem("Roles", () => {
            window.app.router.navigate("/chats/details/roles", this.chat);
        });
        this.addItem("Users", () => {
            window.app.router.navigate("/chats/details/users", this.chat);
        });
    }

    async fetchChat() {
        fetch(`/api/chats/${this.chat.id}`).then(response => {
            if (response.ok) {
                return response.json();
            }
            throw new Error("Failed to fetch chat");
        }).then(data => {
            this.chat = data;
            this.render();
        });
    }

    appear = (data) => {
        if (data) {
            this.chat = data;
        }
        this.setTitle(`Chat Details of ${this.chat.tg_chat_name}`);
        this.fetchChat();
    }
}

export default ChatDetailsView;