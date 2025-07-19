import Form from "../framework/form.js";
import Box from "../framework/box.js";

class EditChatView extends Form {
    constructor(chat) {
        super();
        this.setTitle("Edit Chat");

        this.chat = chat;

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
}

export default EditChatView;