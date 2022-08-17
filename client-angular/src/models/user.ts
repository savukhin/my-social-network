export class User {
    id: number = 0;
    name: string = "";
    username: string = "";
    is_online: boolean = true;
    status: string = "";
    birthDate: string = "";
    city: string = "";
    avatar_url: string | null = null;
    id_token: string | null = null;
    expires_at: string | null = null;

    constructor(id=0, name="", isOnline=true, status="", birthDate="", city="", avatarURL=null) {
        this.id = id;
        this.name = name;
        this.is_online = isOnline;
        this.status = status;
        this.birthDate = birthDate;
        this.city = city;
        this.avatar_url = avatarURL;
    }
}