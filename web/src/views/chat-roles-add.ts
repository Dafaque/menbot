import { FieldType, Form } from "kibodo";
import { Chat } from "../models";

export default class ChatRolesAddView extends Form {
    chat: Chat;
    constructor(data: Chat) {
        super([
            {
                name: "name",
                label: "Name",
                type: FieldType.TEXT,
            },
        ]);
        this.chat = data;
        this.title = `Add Role to Chat ${this.chat.tg_chat_name}`;
        
    }

    onSave = (values: Record<string, string>) => {
        if (!values.name) {
            this.error = "Name is required";
            this.render();
            setTimeout(() => {
                this.error = "";
                this.render();
            }, 1000);
            return;
        }
        fetch(`/api/chats/${this.chat.id}/roles`, {
            method: "POST",
            body: JSON.stringify({
                name: values.name,
            }),
        }).then(response => {
            if (!response.ok) {
                throw new Error("Failed to add role: " + response.statusText);
            }
            window.app.pop();
        }).catch(error => {
            this.error = error;
            this.render();
        });
    }
}
