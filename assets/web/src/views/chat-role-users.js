import Form from "../framework/form.js";

class ChatRoleUsersView extends Form {
    constructor(data) {
        super();
    }

    async fetchUsers() {
        this.roleUsers = {};
        this.users = [];
        fetch(`/api/chats/${this.role.chat_id}/users`).then(response => {
            if (response.ok) {
                return response.json();
            }
        }).then(data => {
            this.users = data.UserListResponseJSONResponse;
        }).then(async () => {
            return fetch(`/api/chats/${this.role.chat_id}/roles/${this.role.id}`).then(response => {
                if (response.ok) {
                    return response.json();
                }
            });
        }).then(data => {
            if (data.UserListResponseJSONResponse) {
                data.UserListResponseJSONResponse.forEach(user => {
                    this.roleUsers[user.tg_user_id] = true;
                    console.log(this.roleUsers);
                });
            }
        }).then(() => {
            for (const user of this.users) {
                this.addField(
                    `user_${user.tg_user_id}`,
                    user.tg_user_name,
                    "checkbox",
                    {
                        value: this.roleUsers[user.tg_user_id] || false,
                    }
                );
            }
            this.render();
        });
    }
    onSave = (values) => {

        var payload = [];
        for (const user of this.users) {
            payload.push({
                user_id: user.id,
                assign: values[`user_${user.tg_user_id}`],
            });
        }
        fetch(`/api/chats/${this.role.chat_id}/roles/${this.role.id}`, {
            method: "POST",
            body: JSON.stringify(payload),
        }).then(response => {
            if (response.ok) {
                this.goBack();
            }
        });
    }

    appear = (data) => {
        if (data) {
            this.role = data;
        }
        this.setTitle(`Users of ${data.name}`);
        this.fetchUsers();
    }
}

export default ChatRoleUsersView;