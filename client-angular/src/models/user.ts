export class User {
    id: number = 0;
    name: string = "";
    username: string = "";
    isOnline: boolean = true;
    status: string = "";
    birthDate: string = "";
    city: string = "";
    avatarURL: string | null = null;

    constructor(id=0, name="", isOnline=true, status="", birthDate="", city="", avatarURL=null) {
        this.id = id;
        this.name = name;
        this.isOnline = isOnline;
        this.status = status;
        this.birthDate = birthDate;
        this.city = city;
        this.avatarURL = avatarURL;
    }
}