import { AfterViewInit, ChangeDetectorRef, Component, ElementRef, OnInit, Renderer2, ViewChild } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Chat } from 'src/models/chat';
import { Message } from 'src/models/message';
import { User } from 'src/models/user';

@Component({
  selector: 'app-chat',
  templateUrl: './chat.component.html',
  styleUrls: ['./chat.component.scss']
})
export class ChatComponent implements AfterViewInit {
  @ViewChild('block') block: ElementRef<HTMLDivElement> = {} as ElementRef;
  contentWidth = 0;
 
  chat: Chat = new Chat(0, "Kirill Klimonov",
     {
        1: new User(1, "Saveliy Karpukhin", true),
        2: new User(2, "Kirill Klimonov", false),
    },
    [
        new Message("I was really cool party, thanks!", new Date(new Date().getSeconds() - 24 * 60 * 60), 1),
        new Message("I glad you've enjoyed :)", new Date(new Date().getSeconds() - 24 * 60 * 60 + 14), 2),
        new Message("Hi! How are you doing?", new Date(), 2),
    ]
  )

  constructor(private route: ActivatedRoute, private router: Router, private renderer: Renderer2, private cdref: ChangeDetectorRef) {
    let id = this.route.snapshot.paramMap.get("id");
    if (id == null)
      this.router.navigateByUrl('404', {skipLocationChange: true})

    // this.chatId = +(id as string);
  }

  onResize(): void {
    this.contentWidth = this.block.nativeElement.scrollWidth;
    this.cdref.detectChanges();
  }

  ngAfterViewInit(): void {
    this.contentWidth = this.block.nativeElement.scrollWidth;
    this.cdref.detectChanges();
    
  }

}
