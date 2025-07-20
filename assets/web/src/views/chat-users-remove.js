import Menu from "../framework/menu.js";

class ChatUsersRemoveView extends Menu {
    constructor(data) {
        super();
        console.log(this);
        this.addItem("Yes, im sure i want to remove this user from the chat", this.removeUser.bind(this));
        this.addItem("Back", this.goBack.bind(this));
    }

    appear = (data) => {
        if (data) {
            this.user = data;
        }
        this.setTitle(`Remove user ${this.user.tg_user_name} from chat ${this.user.tg_chat_name}`);
        this.render();
    }

    async removeUser() {
        fetch(`/api/chats/${this.user.chat_id}/users/${this.user.id}`, {
            method: "DELETE",
        }).then(response => {
            if (response.ok) {
                this.goBack();
            }
        });
    }
    
}

export default ChatUsersRemoveView;