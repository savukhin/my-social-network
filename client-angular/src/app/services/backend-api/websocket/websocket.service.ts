import { Injectable } from '@angular/core';
import { BehaviorSubject, Observable, Observer, Subject, SubscriptionLike } from 'rxjs';
import { Message } from 'src/models/message';

@Injectable({
  providedIn: 'root',
})
export class WebsocketService {
  private socket?: WebSocket;
  public message = new Subject<Message>();

  private onOpen(ev: Event) {
    console.log("Connection opened ", ev);
  }

  private onClose(ev: CloseEvent) {
    console.log("Connection closed ", ev);
  }

  private onMessage(ev: MessageEvent<any>) {
    const response = JSON.parse(ev.data)
    this.message.next(response as Message)
    
  }

  private create (url: string) {
    console.log("create", this.message);
    let socket = new WebSocket(url)

    socket.onclose = (ev) => { this.onClose(ev) }
    socket.onmessage = (ev) => { this.onMessage(ev) }
    socket.onopen = (ev) => { this.onOpen(ev) }

    return socket
  }

  public conect (url: string) {
    if (!this.socket)
      this.socket = this.create(url)

    return this.socket
  }

  public send (message: Message) {
    if (!this.socket) {
      throw new Error("You should init socket first")
    }

    this.socket.send(JSON.stringify(message))
  }
}
