import { Message } from "./message";
import { User } from "./user";

export class ChatDTO {
    id: number = 0;
    title: string = "";
    is_persional: boolean = false
    participants: User[] = []
    messages: Message[] = [];
    
}

export class Chat {
    id: number = 0;
    title: string = "";
    is_persional: boolean = false
    participants: { [id:number] : User; } = {}
    messages: Message[] = [];

    constructor(id=0, title="", participants={}, messages:Message[]=[]) {
        this.id = id;
        this.title = title;
        this.participants = participants;
        this.messages = messages;
    }

    static fromDTO(dto: ChatDTO) {
        let chat = new Chat()
        
        chat.id = dto.id
        chat.title = dto.title;
        dto.participants.forEach(val => {
            chat.participants[val.id] = val
        })
        chat.messages = dto.messages;

        return chat
    }
}