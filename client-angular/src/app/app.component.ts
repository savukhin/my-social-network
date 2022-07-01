import { Component } from '@angular/core';
import { HelloWorldService } from './hello-world.service';
import { Title } from '@angular/platform-browser';

interface Answer {
  title: string
}

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'client-angular';

  constructor(private hw: HelloWorldService, private titleService: Title) {}

  ngOnInit() {
    this.hw.getTitle()
      .subscribe(data => {
        this.title = (data as Answer).title;
        console.log(this.title);
        this.titleService.setTitle(this.title);
      });

    console.log(this.title);
  }
}
