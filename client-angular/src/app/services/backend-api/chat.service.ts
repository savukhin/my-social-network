import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { map } from 'rxjs/operators';
import { environment } from 'src/environments/environment';
import { Chat, ChatDTO } from 'src/models/chat';
import { Message } from 'src/models/message';
import { Post } from 'src/models/post';
import { AuthService } from './auth.service';

@Injectable()
export class ChatService {

  constructor(private auth: AuthService, private http: HttpClient) { }

  private processChat(chat: Chat | ChatDTO) {
    if (!chat.is_personal || chat.participants.length != 2 || !this.auth.user) 
      return chat

    let other_user_index = 0
    if (chat.participants[0].id == this.auth.user.id) {
      other_user_index = 1
    }

    chat.photo_url = chat.participants[other_user_index].avatar_url
    chat.title = chat.participants[other_user_index].name
    return chat
  }

  getChats () {
    if (this.auth.user == undefined)
      return
      
    const token = this.auth.getTokenHeader()
    if (token == false)
      return

    let observer = this.http.get<Chat[]>(
      `${environment.serverUrl}/api/chats`, 
      {headers: token, observe: 'response'}
    )

    return observer.pipe(
      map((response) => {
        if (response.status == 200 && response.body) {
          let chats = response.body as Chat[]

          for (let chat of chats) {
            this.processChat(chat)
          }
          console.log(chats );

          return chats
        }
        return false;
      })
    )
  }

  getChat (chat_id: number) {
    if (this.auth.user == undefined)
      return
      
    const token = this.auth.getTokenHeader()
    if (token == false)
      return

    let observer = this.http.post<ChatDTO>(
      `${environment.serverUrl}/api/chat/by_id/${chat_id}`, 
      {},
      {headers: token, observe: 'response'}
    )

    return observer.pipe(
      map((response) => {
        if (response.status == 200 && response.body) {
            return this.processChat(response.body as ChatDTO)
        }
        return false;
      })
    )
  }

  getPersonalChat (user_id: number) {
    if (this.auth.user == undefined)
      return
      
    const token = this.auth.getTokenHeader()
    if (token == false)
      return

    let observer = this.http.post<ChatDTO>(
      `${environment.serverUrl}/api/chat/by_user/${user_id}`, 
      {},
      {headers: token, observe: 'response'}
    )

    return observer.pipe(
      map((response) => {
        if (response.status == 200 && response.body) {
          return this.processChat(response.body as ChatDTO)
        }
        return false;
      })
    )
  }

  getPersonalChatMessages (offset: number, count: number, chat_id: number) {
    if (this.auth.user == undefined)
      return
      
    const token = this.auth.getTokenHeader()
    if (token == false)
      return

    let observer = this.http.post<{messages: Message[]}>(
      `${environment.serverUrl}/api/chat/getMessages`, 
      {offset, count, chat_id},
      {headers: token, observe: 'response'}
    )

    return observer.pipe(
      map((response) => {
        if (response.status == 200 && response.body) {
          return response.body as {messages: Message[]}
        }
        return false;
      })
    )
  }

  sendMessage (chat_id: number, text: string) {
    if (this.auth.user == undefined)
      return
      
    const token = this.auth.getTokenHeader()
    if (token == false)
      return

    let observer = this.http.post<{message: string}>(
      `${environment.serverUrl}/api/chat/sendMessage`, 
      {chat_id, text},
      {headers: token, observe: 'response'}
    )

    return observer.pipe(
      map((response) => {
        if (response.status == 200 && response.body) {
          return response.body as {message: string}
        }
        return false;
      })
    )
  }
}
