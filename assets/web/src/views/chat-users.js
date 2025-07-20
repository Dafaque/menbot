import Table from "../framework/table.js";

class ChatUsersView extends Table {
    constructor(data) {
        super();
        this.setTitle("Chat Users");
        this.chat = data;
    }


    async fetchUsers() {
        this.data = [];
        fetch(`/api/chats/${this.chat.id}/users`).then(response => {
            if (response.ok) {
                return response.json();
            }
            throw new Error("Failed to fetch users");
        }).then(data => {
            if (data.UserListResponseJSONResponse) {
                this.data = data.UserListResponseJSONResponse;
                this.render();
            }
        });
    }

    onSelected = (row) => {
        window.app.router.navigate("/chats/details/users/remove", row);
    }

    appear = (data) => {
        if (data) {
            this.chat = data;
        }
        this.fetchUsers();
    }
}

export default ChatUsersView;