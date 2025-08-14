import { Form, FieldType } from "kibodo";
import { Role, User } from "../models";

export default class ChatRoleUsersView extends Form {
    role: Role;
    users: User[];
    roleUsers: { [key: string]: boolean };
    constructor(data: Role) {
        super();
        this.role = data;
        this.title = `Users of ${data.name}`;
        this.fetchUsers();
    }

    async fetchUsers() {
        this.roleUsers = {};
        this.users = [];
        fetch(`/api/chats/${this.role.chat_id}/users`).then(response => {
            if (!response.ok) {
                throw new Error("Failed to fetch users: " + response.statusText);
            }
            return response.json();
        }).then(data => {
            this.users = data.UserListResponseJSONResponse;
        }).then(async () => {
            return fetch(`/api/chats/${this.role.chat_id}/roles/${this.role.id}`).then(response => {
                if (!response.ok) { throw new Error("Failed to fetch role users: " + response.statusText); }
                return response.json();
            });
        }).then(data => {
            if (data.UserListResponseJSONResponse) {
                data.UserListResponseJSONResponse.forEach((user: User) => {
                    this.roleUsers[user.tg_user_id] = true;
                });
            }
        }).then(() => {
            this.users.forEach(user => {
                this.fields.push({
                    name: user.tg_user_id,
                    label: user.tg_user_name,
                    type: FieldType.CHECKBOX,
                    value: this.roleUsers[user.tg_user_id] || false,
                });
            });
        }).catch(error => {
            this.error = error;
        }).finally(() => {
            this.render();
        });
    }
    
    onSave = (values: Record<string, boolean>) => {
        var payload = [];
        for (const user of this.users) {
            payload.push({
                user_id: user.id,
                assign: values[user.tg_user_id],
            });
        }
        fetch(`/api/chats/${this.role.chat_id}/roles/${this.role.id}`, {
            method: "POST",
            body: JSON.stringify(payload),
        }).then(response => {
            if (!response.ok) {
                throw new Error("Failed to save role users: " + response.statusText);
            }
        }).then(() => {
            window.app.pop();
        }).catch(error => {
            this.error = error;
            this.render();
        });
    }
}
