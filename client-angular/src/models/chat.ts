import { Message } from "./message";
import { User } from "./user";

export class Chat {
    id: number = 0;
    title: string = "";
    is_persional: boolean = false
    participants: { [id:number] : User; } = {}
    messages: Message[] = [];

    constructor(id=0, title="", participants={}, messages:Message[]=[]) {
        this.id = id;
        this.title = title;
        this.participants=participants;
        this.messages=messages;
    }
}