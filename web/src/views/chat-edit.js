import { Form } from "kibodo";

class EditChatView extends Form {
    constructor(chat) {
        super();
    }

    onSave = (values) => {
        fetch(`/api/chats/${this.chat.id}?authorized=${values.authorized}`, {
            method: "PUT",
        }).then(response => {
            if (response.ok) {
                this.goBack();
            } else {
                console.error("Failed to update chat");
            }
        });
    }

    appear = (data) => {
        this.chat = data;
        this.setTitle(`Edit Chat ${this.chat.tg_chat_name}`);

        this.addField(
            "id", "ID", "text", {
                defaultValue: this.chat.id,
                readonly: true,
            }
        );
        this.addField(
            "tg_chat_id", "TG Chat ID", "text", {
                value: this.chat.tg_chat_id,
                readonly: true,
            }
        );
        this.addField("tg_chat_name", "TG Chat Name", "text", {
            value: this.chat.tg_chat_name,
            readonly: true,
        });
        this.addField("authorized", "Authorized", "checkbox", {
            value: this.chat.authorized,
        });
        this.render();
    }
}

export default EditChatView;