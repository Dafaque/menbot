import { Form, FieldType } from "kibodo";
import { Chat } from "../models";

export default class EditChatView extends Form {
    chat: Chat;

    constructor(chat: Chat) {
        super();
        this.title = `Edit Chat ${chat.tg_chat_name}`;
        this.chat = chat;
        this.fetchChat();
    }

    onSave = (values: Record<string, string | number | boolean>) => {
        fetch(`/api/chats/${this.chat.id}?authorized=${values.authorized}`, {
            method: "PUT",
        }).then(response => {
            if (!response.ok) {
                throw new Error("Failed to update chat");
            }
            window.app.pop();
        }).catch(error => {
            this.error = error;
            this.render();
        });
    }

    async fetchChat() {
        fetch(`/api/chats/${this.chat.id}`).then(response => {
            if (!response.ok) {
                throw new Error("Failed to fetch chat");
            }
            return response.json();
        }).then(chat => {
            this.fields = [
                {
                    name: "id",
                    label: "ID",
                    type: FieldType.TEXT,
                    value: chat.id,
                    readonly: true,
                },
                {
                    name: "tg_chat_id",
                    label: "TG Chat ID",
                    type: FieldType.TEXT,
                    value: chat.tg_chat_id,
                    readonly: true,
                },
                {
                    name: "tg_chat_name",
                    label: "TG Chat Name",
                    type: FieldType.TEXT,
                    value: chat.tg_chat_name,
                    readonly: true,
                },
                {
                    name: "authorized",
                    label: "Authorized",
                    type: FieldType.CHECKBOX,
                    value: chat.authorized,
                },
            ]
        }).catch(error => {
            this.error = error;
        }).finally(() => {
            this.render();
        });
    }

}
